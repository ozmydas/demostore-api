package app

import (
	"log"
	"net/http"
)

type Handler struct {
	*Env
	H func(e *Env, w http.ResponseWriter, r *http.Request) (int, error)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := h.H(h.Env, w, r)

	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
			// err := ah.renderTemplate(w, "http_404.tmpl", nil)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}
