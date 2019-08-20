package app

import (
	"log"
	"net/http"
)

/****
* middleware untuk keperluan auth dan aksi lain sebelum method controller utama dipanggil (atau sesudahnya)
****/

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before AUTH") // we can auth session here
		next.ServeHTTP(w, r)       // call original
		log.Println("After AUTH")
	})
} // end func

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before JWT") // we can validate JWT here
		next.ServeHTTP(w, r)      // call original
		log.Println("After JWT")
	})
} // end func

func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before JWT") // we can validate JWT for auth user here
		next.ServeHTTP(w, r)      // call original
		log.Println("After JWT")
	})
} // end func
