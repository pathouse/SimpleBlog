package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"simpleblog/app/controllers"
	"simpleblog/app/support"
)

func main() {
	// parses the command-line flags from os.Args[1:]
	flag.Parse()
	// ensure everything in the log buffer is printed before main func returns
	defer glog.Flush()

	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Home)

	serverLogger := support.NewServerLogger(router, os.Stderr)
	server := &http.Server{
		Addr:    ":4000",
		Handler: serverLogger,
	}
	server.ListenAndServe()
}
