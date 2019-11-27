package model

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"../lib"
	"github.com/jinzhu/gorm"
)

type Member struct {
	Id, Active                                     int
	Email, Password                                string
	Fullname, Gender, Description, CreatedDatetime string
}

type MemberAllColumn struct {
	Email, Password, CreatedDatetime, Fullname   string
	Birthplace, Birthdate, Phone, Gender, Hobby  string
	Address, BankName, BankAccount, BankRekening string
	ImagesPath, ImagesFile                       string
	ProvinceId, CityId                           int
}

type Ada struct {
	Id int
}

func (Member) TableName() string {
	return "t_members"
}

func (MemberAllColumn) TableName() string {
	return "t_members"
}

/**********/

func MemberGetActiveSession(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	sesi, err := lib.GetSession("Memberdata", w, r)

	if err != nil {
		return make(map[string]interface{}), errors.New("Member Session Not Found")
	}

	sesi_arr := sesi.(map[string]interface{})
	return sesi_arr, nil
} // end func

func MemberGetActiveId(w http.ResponseWriter, r *http.Request) (int, error) {
	sesi_arr, err := MemberGetActiveSession(w, r)

	if err != nil {
		return 0, errors.New("Member Session Not Found")
	}

	return sesi_arr["ID"].(int), nil
} // end func

func MemberGetActiveName(w http.ResponseWriter, r *http.Request) (string, error) {
	sesi_arr, err := MemberGetActiveSession(w, r)

	if err != nil {
		return "", errors.New("Member Session Not Found")
	}

	return sesi_arr["Membername"].(string), nil
} // end func

/******/

func ModMemberAuth(DB *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	now := time.Now().Format("2006-01-02 15:04:05") // only date

	// cari Member by email
	table := []Member{}
	proses := DB.Where("email = ?", r.PostFormValue("email")).First(&table)

	// not found
	if proses.RowsAffected == 0 || proses.Error != nil {
		var msg error

		if proses.Error != nil {
			msg = proses.Error
		} else {
			msg = errors.New("Member Not Found")
		}

		return r.PostFormValue, msg
	}

	// type assertion biar index bisa dipanggil
	members := proses.Value.(*[]Member)
	Member_found := (*members)[0]

	// match password

	// update last login

	// done
	generatedToken, _ := lib.GenerateJWT(strconv.Itoa(Member_found.Id))

	result := map[string]interface{}{
		"Token":    generatedToken,
		"Datetime": now,
		"Member":   Member_found,
		// "Membername": registeredMember.Membername,
		// "url":      "/profile?ref=" + token,
	}

	return result, nil
} // end func

func ModMemberRegister(DB *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	tgl_date := time.Now().Format("2006-01-02") // only date
	tgl_time := time.Now().Format("15:04:05")   // only time

	// assign post to struct
	row := MemberAllColumn{
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		Fullname:        r.PostFormValue("fullname"),
		Birthplace:      r.PostFormValue("birthplace"),
		Birthdate:       r.PostFormValue("birthdate"),
		Phone:           r.PostFormValue("phone"),
		Gender:          r.PostFormValue("gender"),
		Address:         r.PostFormValue("address"),
		CreatedDatetime: tgl_date + " " + tgl_time,
	}

	// cek email here
	// if register, reject registerer

	// simpan ke table log
	save := DB.Create(&row)
	log.Printf("%+v", save.Error)

	// jika error
	if save.Error != nil {
		return r.PostFormValue, save.Error
	}

	// done
	return save.Value.(*MemberAllColumn), nil
} // end func

/*** PROFILE ***/

func ModMemberView(kode string, DB *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// now := time.Now().Format("2006-01-02 15:04:05")

	var intKode int

	if kode == "me" {
		tmpIntKode, err := MemberGetActiveId(w, r)

		if err != nil {
			return kode, err
		}

		intKode = tmpIntKode
	} else {
		intKode, _ = strconv.Atoi(kode)
	}

	log.Println(intKode)

	// cari Member by email
	table := []Member{}
	proses := DB.Where("id = ?", intKode).First(&table)

	// not found
	if proses.RowsAffected == 0 || proses.Error != nil {
		var msg error

		if proses.Error != nil {
			msg = proses.Error
		} else {
			msg = errors.New("Member Not Found")
		}

		return r.PostFormValue, msg
	}

	// type assertion biar index bisa dipanggil
	members := proses.Value.(*[]Member)
	Member_found := (*members)[0]

	return Member_found, nil
} // end func

func ModMemberUpdate(DB *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// panggil row yg mau diedit
	id := r.PostFormValue("id")
	row := MemberAllColumn{}
	DB.Where("id = ?", id).First(&row)

	// set nilai baru
	row.Fullname = r.PostFormValue("fullname")
	row.Birthplace = r.PostFormValue("birthplace")
	row.Birthdate = r.PostFormValue("birthdate")
	row.Phone = r.PostFormValue("phone")
	row.Gender = r.PostFormValue("gender")
	row.Address = r.PostFormValue("address")

	// simpan
	proses := DB.Save(&row)

	if proses.Error != nil {
		return row, proses.Error
	}

	return row, nil
}
