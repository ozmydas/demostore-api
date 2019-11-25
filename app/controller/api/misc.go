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
)

func MiscBanner(e *app.Env, w http.ResponseWriter, r *http.Request) (int, error) {
	// lookup ke table
	result, err := model.ModBannerList(e.DB, w, r)

	// err found
	if err != nil {
		lib.JsonRender(w, false, err.Error(), data_prototype, 102)
		return 200, nil
	}

	lib.JsonRender(w, true, "Data Found", result, 200)
	return 200, nil
}
