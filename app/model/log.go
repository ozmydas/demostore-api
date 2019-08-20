package model

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type LogApiCall struct {
	Id, UserId                                int
	Action, Params, Response, CreatedDatetime string
}

type LogAppInit struct {
	DeviceId, OsName, OsVersion, Manufacture, Model string
	Ip, Provider, NetworkType                       string
	DeviceCpu                                       string
	DeviceRam                                       int
	CreatedDate, CreatedTime                        string
}

func (LogAppInit) TableName() string {
	return "log_app_raw"
}

/******/

func ModLogDeviceSave(DB *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	tgl_date := time.Now().Format("2006-01-02")             // only date
	tgl_time := time.Now().Format("15:04:05")               // only time
	myRam, _ := strconv.Atoi(r.PostFormValue("device_ram")) // convert value to int

	// assign post to struct
	row := LogAppInit{
		DeviceId:    r.PostFormValue("device_id"),
		OsName:      r.PostFormValue("os_name"),
		OsVersion:   r.PostFormValue("os_version"),
		Manufacture: r.PostFormValue("manufacture"),
		Model:       r.PostFormValue("model"),
		Ip:          r.PostFormValue("ip_address"),
		Provider:    r.PostFormValue("provider"),
		NetworkType: r.PostFormValue("network_type"),
		DeviceCpu:   r.PostFormValue("device_cpu"),
		DeviceRam:   myRam,
		CreatedDate: tgl_date,
		CreatedTime: tgl_time,
	}

	// simpan ke table log
	save := DB.Create(&row)
	log.Printf("%+v", save.Error)

	// jika error
	if save.Error != nil {
		return r.PostFormValue, save.Error
	}

	// done
	return save.Value.(*LogAppInit), nil
} // end func
