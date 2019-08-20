package api

import (
	_ "encoding/json"
	"log"
	"net/http"
	"net/url"

	"../../../app"
	"../../lib"
	"../../model"
	_ "github.com/gorilla/mux"
	"github.com/thedevsaddam/govalidator"
)

/******/

// struct untuk menampung input post form

type MemberLoginInput struct {
	Email, Password string
}

type MemberRegisterInput struct {
	Email, Password, VerifyPassword                                                              string
	Fullname, Birthplace, Birthdate, Phone, Gender, Hobby, Lulusan                               string
	Address, ProvinceId, CityId, BankName, BankAccount, BankRekening                             string
	ImagesPath, ImagesFile, FilesIjazah, FilesSertifikat_1, FilesSertifikat_2, FilesSertifikat_3 string
}

/******/

func MemberLogin(e *app.Env, w http.ResponseWriter, r *http.Request) (int, error) {
	// validasi form
	isFormValid := MemberValidateLogin(r)

	if len(isFormValid) != 0 {
		log.Printf("%+v", isFormValid)
		lib.JsonRender(w, false, "FAIL", isFormValid, 101)
		return 200, nil
	}

	// lookup ke table
	result, err := model.ModMemberAuth(e.DB, w, r)

	// err found
	if err != nil {
		lib.JsonRender(w, false, "Login Failed Error", err.Error(), 102)
		return 200, nil
	}

	lib.JsonRender(w, true, "Login Success", result, 200)
	return 200, nil
}

/*** VALIDATOR ***/

func MemberValidateLogin(r *http.Request) url.Values {
	rules := govalidator.MapData{
		"email":    []string{"required"},
		"password": []string{"required"},
	}

	return lib.ExecuteValidator(rules, r)
} // end func

func MemberValidateReg(r *http.Request) url.Values {
	rules := govalidator.MapData{
		"email":           []string{"required"},
		"password":        []string{"required"},
		"verify_password": []string{"required"},
		"fullname":        []string{"required"},
	}

	return lib.ExecuteValidator(rules, r)
} // end func
