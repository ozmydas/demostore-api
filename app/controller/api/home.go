package api

import (
	_ "log"
	"net/http"

	"../../../app"
	"../../lib"
	"../../model"
	"github.com/thedevsaddam/govalidator"
)

var (
	data_prototype = make(map[string]interface{})
)

/* inisialisasi, catat log dll */
func HomeStart(e *app.Env, w http.ResponseWriter, r *http.Request) (int, error) {
	// validasi
	rules := govalidator.MapData{
		"device_id":    []string{"required"},
		"os_name":      []string{"required"},
		"os_version":   []string{"required"},
		"manufacture":  []string{"required"},
		"model":        []string{"required"},
		"ip_address":   []string{"required"},
		"provider":     []string{"required"},
		"network_type": []string{"required"},
		"device_cpu":   []string{"required"},
		"device_ram":   []string{"required"},
	}

	opts := govalidator.Options{
		Request: r,     // request object
		Rules:   rules, // rules map
	}

	v := govalidator.New(opts)
	validatError := v.Validate()

	if len(validatError) != 0 {
		// log.Printf("%+v", validatError)
		lib.JsonRender(w, false, "FAIL", validatError, 101)
		return 200, nil
	}

	// catat ke log model
	result, err := model.ModLogDeviceSave(e.DB, w, r)

	// jika error
	if err != nil {
		lib.JsonRender(w, false, "FAIL", result, 102)
		return 200, nil
	}

	// done
	lib.JsonRender(w, true, "OK", result, 200)
	return 200, nil
} // end func
