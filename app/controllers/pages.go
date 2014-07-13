package controllers

import (
	"io"
	"net/http"
)

type Page struct {
	Title string
}

func Home(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Hello World!\n")
}
