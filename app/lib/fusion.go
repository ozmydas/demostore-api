package lib

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/thedevsaddam/govalidator"
)

func JsonResult(status bool, code int, msg string, data interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	result["status"] = status
	result["message"] = msg
	result["data"] = data
	result["code"] = code

	return result
}

func JsonOutput(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func JsonRender(w http.ResponseWriter, status bool, msg string, data interface{}, code int) {
	result := JsonResult(status, code, msg, data)
	JsonOutput(w, 200, result)
}

// rand
func SetSeed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandomString(l int) string {
	bytes := make([]byte, l)

	for i := 0; i < l; i++ {
		bytes[i] = byte(RandomInt(65, 90))
	}

	return string(bytes)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)

}

// Define your Error struct
type MyError struct {
	msg string
}

// Create a function Error() string and associate it to the struct.
func (error *MyError) Error() string {
	return error.msg
}

// for validation of a form

func ExecuteValidator(rules map[string][]string, r *http.Request) url.Values {
	opts := govalidator.Options{
		Request: r,     // request object
		Rules:   rules, // rules map
	}

	v := govalidator.New(opts)

	return v.Validate()
} // end func

// str padd zero on left
func StrPadZero(value int) string {
	return fmt.Sprintf("%02d", value)
} // end func

// this will turn integer to alphabet. ex : 1 to A, 2 to B, or 27 to AA
func toCharStrArr(i int) string {
	var arr = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if i > 26 {
		puluhan := i / 26
		selisih := i % 26

		if selisih == 0 {
			puluhan = puluhan - 1
			selisih = 26
		}

		depan := toCharStrArr(puluhan)
		belakang := toCharStrArr(selisih)
		return depan + belakang
	}

	return arr[i-1]
} // end func
