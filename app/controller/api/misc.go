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
	_ "github.com/gorilla/mux"
)

func MiscBanner(e *app.Env, w http.ResponseWriter, r *http.Request) (int, error) {
	// lookup ke table
	result, count, err := model.ModBannerList(e.DB, w, r)

	// err found
	if err != nil {
		lib.JsonRender(w, false, err.Error(), data_prototype, 102)
		return 200, nil
	}

	jml := strconv.FormatInt(count, 10)

	lib.JsonRender(w, true, jml+" Data Found", result, 200)
	return 200, nil
}
