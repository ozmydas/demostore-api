package lib

import (
	"encoding/gob"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const cookieName = "aptx-4869"

var Store *sessions.CookieStore
var empty interface{}

func InitSession() {
	Store = sessions.NewCookieStore([]byte("super-secret-key"))

	Store.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 8,
		HttpOnly: true,
	}

	gob.Register(map[string]interface{}{})
}

func SetSession(key string, value interface{}, w http.ResponseWriter, r *http.Request) interface{} {
	if Store == nil {
		InitSession()
	}

	session, err := Store.Get(r, cookieName)

	if err != nil {
		log.Printf("set: %+v", err.Error())
		return empty
	}

	// assign new value to session
	session.Values[key] = value

	// save changes to curent session
	err = session.Save(r, w)
	if err != nil {
		log.Printf("save: %+v", err.Error())
		return empty
	}

	return session.Values[key]
} // end func

func GetSession(key string, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if Store == nil {
		InitSession()
	}

	session, err := Store.Get(r, cookieName)

	if err != nil {
		return empty, err
	}

	if session.Values[key] == nil {
		log.Printf("get: %+v", session.Values)
		return session.Values[key], errors.New("Empty Session")
	}

	// log.Printf("get: %+v", session)
	result := session.Values[key]
	return result, nil
} //end func

func ClearSession(w http.ResponseWriter, r *http.Request) error {
	if Store == nil {
		InitSession()
	}

	session, err := Store.Get(r, cookieName)

	if err != nil {
		return err
	}

	session.Options.MaxAge = -1

	// save changes to curent session
	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}
