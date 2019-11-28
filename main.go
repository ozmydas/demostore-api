package main

import (
	"log"
	"net/http"

	"./app"
	"./app/controller"
	"./app/controller/api"
	"./app/lib"
	"github.com/gorilla/mux"
)

func main() {
	lib.SetSeed()

	context := &app.Env{
		DB:          app.InitDB(),
		PORT:        ":6060",
		UPLOAD_PATH: "public/files",
	}

	/******/

	mx := mux.NewRouter()
	mx.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	/** HOME **/

	// index
	mx.Handle("/", app.Handler{context, controller.HomeWelcome}).Methods("GET")

	// startup : saat aplikasi pertama dijalankan, cek koneksi ke server serta kirim informasi device user
	mx.Handle("/api/start", app.JwtMiddleware(app.Handler{context, api.HomeStart})).Methods("POST")

	// login user
	mx.Handle("/api/login", app.JwtMiddleware(app.Handler{context, api.MemberLogin})).Methods("POST")

	// register user
	mx.Handle("/api/register", app.JwtMiddleware(app.Handler{context, api.MemberRegister})).Methods("POST")

	// list product
	mx.Handle("/api/product", app.JwtMiddleware(app.Handler{context, api.ProductList})).Methods("POST")

	// detail product
	mx.Handle("/api/product/{id}", app.JwtMiddleware(app.Handler{context, api.ProductDetail})).Methods("GET")

	// misc. list banner
	mx.Handle("/api/banner", app.JwtMiddleware(app.Handler{context, api.MiscBanner})).Methods("GET")

	// profile
	mx.Handle("/api/profile/update", app.JwtMiddleware(app.Handler{context, api.MemberUpdate})).Methods("POST")
	mx.Handle("/api/profile/{id}", app.JwtMiddleware(app.Handler{context, api.MemberView})).Methods("POST")

	// cart
	mx.Handle("/api/cart/add", app.JwtMiddleware(app.Handler{context, api.CartAdd})).Methods("POST")
	mx.Handle("/api/cart/view/{id}", app.JwtMiddleware(app.Handler{context, api.CartView})).Methods("GET")

	/******/

	// go app.HandleMessages()
	log.Println("running on", context.PORT)
	log.Fatal(http.ListenAndServe(context.PORT, mx))
}
