package model

import (
	"errors"
	_ "log"
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

type Ada struct {
	Id int
}

func (Member) TableName() string {
	return "t_member"
}

/**********/

func MemberGetActiveSession(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	sesi, err := lib.GetSession("Memberdata", w, r)

	if err != nil {
		return make(map[string]interface{})
	}

	sesi_arr := sesi.(map[string]interface{})
	return sesi_arr
} // end func

func MemberGetActiveId(w http.ResponseWriter, r *http.Request) int {
	sesi_arr := MemberGetActiveSession(w, r)
	return sesi_arr["ID"].(int)
} // end func

func MemberGetActiveName(w http.ResponseWriter, r *http.Request) string {
	sesi_arr := MemberGetActiveSession(w, r)
	return sesi_arr["Membername"].(string)
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
