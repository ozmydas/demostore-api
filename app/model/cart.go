package model

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	_ "time"

	_ "../lib"
	"github.com/jinzhu/gorm"
)

type Cart struct {
	Id, ProductId, MemberId, Price int
	Name, Cat, Img                 string
}

func (Cart) TableName() string {
	return "t_cart"
}

/*******/

func ModCartAdd(DB *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, string, error) {
	// check product first base on product id
	product, err := ModProductDetail(r.PostFormValue("product_id"), DB, w, r)

	if err != nil {
		return nil, "", err
	}

	// also check stok if needed

	// check if product id already in cart if needed

	// end get product here
	productId, _ := strconv.Atoi(r.PostFormValue("product_id"))
	memberId, _ := strconv.Atoi(r.PostFormValue("member_id"))
	prod := product.(Products)

	row := Cart{
		ProductId: productId,
		MemberId:  memberId,
		Price:     prod.Price,
		Name:      prod.Name,
		Cat:       prod.Cat,
		Img:       prod.Img,
	}

	// simpan
	save := DB.Create(&row)
	log.Printf("%+v", save.Error)

	// jika error
	if save.Error != nil {
		return r.PostFormValue, "", save.Error
	}

	// done
	return save.Value.(*Cart), prod.Name, nil
} // end func

func ModCartView(kode string, DB *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, int64, error) {
	// select from table
	table := []Cart{}
	proses := DB.Where("member_id = ?", kode).Find(&table)

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
