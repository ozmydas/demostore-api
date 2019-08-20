package controller

import (
	_ "log"
	"net/http"

	"../../app"
	"../lib"
)

var (
	data_prototype = make(map[string]interface{})
)

func HomeWelcome(e *app.Env, w http.ResponseWriter, r *http.Request) (int, error) {
	data := data_prototype
	lib.TemplateRender(w, lib.TemplatePath(), data)
	return 200, nil
} // end func
