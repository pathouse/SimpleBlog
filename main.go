package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"simpleblog/app/assets"
	"simpleblog/app/modcon"
	"simpleblog/app/support"
)

func main() {
	db := OpenDB()
	defer db.Close()

	modcon.AutoMigrate(&db)

	context := modcon.NewAppContext(&db)

	router := mux.NewRouter()

	router.Handle("/",
		modcon.AppHandler{
			Context: context,
			Handler: modcon.IndexHandler}).
		Methods("GET")

	router.Handle("/about",
		modcon.AppHandler{
			Context: context,
			Handler: modcon.AboutHandler}).
		Methods("GET")

	router.Handle("/register",
		modcon.AppHandler{
			Context: context,
			Handler: modcon.RegistrationHandler}).
		Methods("GET")

	router.Handle("/register",
		modcon.RedirectHandler{
			Context: context,
			Handler: modcon.CreateUserHandler}).
		Methods("POST")

	fileServer := http.StripPrefix("/assets/", http.FileServer(assets.NewAssetFileSys()))
	router.PathPrefix("/").Handler(fileServer)

	serverLogger := support.NewServerLogger(router, os.Stderr)

	http.Handle("/", serverLogger)
	http.ListenAndServe(":4000", nil)
}
