package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"simpleblog/app/assets"
	"simpleblog/app/controllers"
	"simpleblog/app/support"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/", controllers.NewPageHandler("Home", "Welcome Home Son"))
	router.Handle("/about", controllers.NewPageHandler("About", "What's this all about?"))

	fileServer := http.StripPrefix("/assets/", http.FileServer(assets.NewAssetFileSys()))
	router.PathPrefix("/").Handler(fileServer)

	serverLogger := support.NewServerLogger(router, os.Stderr)

	http.Handle("/", serverLogger)
	http.ListenAndServe(":4000", nil)
}
