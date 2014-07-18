package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"simpleblog/app/assets"
	"simpleblog/app/modcon"
	"simpleblog/app/support"
)

func main() {
	db, err := gorm.Open("postgres", "dbname=blog_dev sslmode=disable")
	if err != nil {
		support.LogStacktrace(err)
	}
	defer db.Close()

	//defaults
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().Ping()

	modcon.AutoMigrate(&db)

	context := modcon.NewAppContext(&db)

	router := mux.NewRouter()

	router.Handle("/", modcon.NewPageHandler(context, "Home", "Welcome Home Son"))
	router.Handle("/about", modcon.NewPageHandler(context, "About", "What's this all about?"))

	fileServer := http.StripPrefix("/assets/", http.FileServer(assets.NewAssetFileSys()))
	router.PathPrefix("/").Handler(fileServer)

	serverLogger := support.NewServerLogger(router, os.Stderr)

	http.Handle("/", serverLogger)
	http.ListenAndServe(":4000", nil)
}
