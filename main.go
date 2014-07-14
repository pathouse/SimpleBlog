package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"simpleblog/app/controllers"
	"simpleblog/app/support"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Home)

	// assets in our templates are not loaded like other files
	// they're accessed via HTTP, so we need to serve them as static files
	// StripPrefix returns a handler that serves HTTP requests by removing the
	// given prefix from the request URL's Path and invoking the handler h.
	// basically, we strip the name of the folder so we can serve the whole folder
	// and know that the remaining path works inside that folder
	fileServer := http.StripPrefix("/assets/", http.FileServer(http.Dir("/Users/patsicle/workspace/go/src/simpleblog/assets")))
	router.PathPrefix("/").Handler(fileServer)

	serverLogger := support.NewServerLogger(router, os.Stderr)

	http.Handle("/", serverLogger)
	http.ListenAndServe(":4000", nil)
}
