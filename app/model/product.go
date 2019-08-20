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

type Products struct {
	Id, Published, Price                                  int
	Name, Thumbnail, Images, Description, CreatedDatetime string
}

func (Products) TableName() string {
	return "t_product"
}

/******/

func ModProductList(DB *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, int64, error) {
	// select from table
	table := []Products{}
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

func ModProductDetail(kode string, DB *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// select from table
	table := []Products{}
	proses := DB.Where("published = ? AND id = ?", 1, kode).First(&table)

	if proses.RowsAffected == 0 || proses.Error != nil {
		var msg error

		if proses.Error != nil {
			msg = proses.Error
		} else {
			msg = errors.New("Data Not Found")
		}

		return r.PostFormValue, msg
	}

	// type assertion biar index bisa dipanggil
	result := proses.Value.(*[]Products)
	result_found := (*result)[0]

	return result_found, nil
} // end func
