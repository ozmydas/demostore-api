package api

import (
	_ "encoding/json"
	_ "log"
	"net/http"
	_ "net/url"
	"strconv"

	"../../../app"
	"../../lib"
	"../../model"
	"github.com/gorilla/mux"
	_ "github.com/thedevsaddam/govalidator"
)

/******/

/******/

func ProductList(e *app.Env, w http.ResponseWriter, r *http.Request) (int, error) {
	// lookup ke table
	result, total, err := model.ModProductList(e.DB, w, r)

	// err found
	if err != nil {
		lib.JsonRender(w, false, err.Error(), data_prototype, 102)
		return 200, nil
	}

	jumlah := strconv.FormatInt(total, 10)

	lib.JsonRender(w, true, jumlah+" Item(s) Found", result, 200)
	return 200, nil
}

func ProductDetail(e *app.Env, w http.ResponseWriter, r *http.Request) (int, error) {
	// cek variable terlampir
	params := mux.Vars(r)

	if params["id"] == "" {
		lib.JsonRender(w, false, "Empty Identifier", data_prototype, 101)
		return 200, nil
	}

	kode := params["id"]

	// lookup ke table
	result, err := model.ModProductDetail(kode, e.DB, w, r)

	// err found
	if err != nil {
		lib.JsonRender(w, false, err.Error(), data_prototype, 102)
		return 200, nil
	}

	lib.JsonRender(w, true, "Data Found", result, 200)
	return 200, nil
}
