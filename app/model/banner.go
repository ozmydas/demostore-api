package model

import (
	"errors"
	_ "log"
	"net/http"
	_ "strconv"
	_ "time"

	_ "../lib"
	"github.com/jinzhu/gorm"
)

/******/

type Banners struct {
	Id                        int
	Images, Link, Description string
}

func (Banners) TableName() string {
	return "t_banner"
} // end func

/******/

func ModBannerList(DB *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, int64, error) {
	// select from table
	table := []Banners{}
	proses := DB.Where("published = ?", 1).Find(&table)

	if proses.RowsAffected == 0 || proses.Error != nil {
		var msg error

		if proses.Error != nil {
			msg = proses.Error
		} else {
			msg = errors.New("Data Not Found")
		}

		return r.PostFormValue, 0, msg
	}

	result := proses.Value

	return result, proses.RowsAffected, nil
} // end func
