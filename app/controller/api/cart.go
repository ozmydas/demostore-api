package api

import (
	"log"
	"net/http"
	"net/url"
	"strconv"

	"../../../app"
	"../../lib"
	"../../model"
	"github.com/gorilla/mux"
	"github.com/thedevsaddam/govalidator"
)

/******/

/******/

func CartView(e *app.Env, w http.ResponseWriter, r *http.Request) (int, error) {
	params := mux.Vars(r)

	if params["id"] == "" {
		lib.JsonRender(w, false, "Empty Identifier", data_prototype, 101)
		return 200, nil
	}

	kode := params["id"]

	// lookup ke table
	result, total, err := model.ModCartView(kode, e.DB, w, r)

	// err found
	if err != nil {
		lib.JsonRender(w, false, err.Error(), data_prototype, 102)
		return 200, nil
	}

	jumlah := strconv.FormatInt(total, 10)

	lib.JsonRender(w, true, jumlah+" Item(s) Found", result, 200)
	return 200, nil
}

func CartAdd(e *app.Env, w http.ResponseWriter, r *http.Request) (int, error) {
	// validasi form
	isFormValid := CartValidateAdd(r)

	if len(isFormValid) != 0 {
		log.Printf("%+v", isFormValid)
		lib.JsonRender(w, false, "FAIL", isFormValid, 101)
		return 200, nil
	}

	// lookup ke table
	result, product, err := model.ModCartAdd(e.DB, w, r)

	// err found
	if err != nil {
		lib.JsonRender(w, false, err.Error(), data_prototype, 102)
		return 200, nil
	}

	lib.JsonRender(w, true, product+" Successfully Added to Cart ", result, 200)
	return 200, nil
} // end func

/******/

func CartValidateAdd(r *http.Request) url.Values {
	rules := govalidator.MapData{
		"member_id":  []string{"required"},
		"product_id": []string{"required"},
		"qty":        []string{"required"},
	}

	return lib.ExecuteValidator(rules, r)
} // end func
