package controllers

import (
	"net/http"
	"simpleblog/app/templates"
)

type PageHandler struct {
	Title string
	Body  string
}

func NewPageHandler(title, body string) http.Handler {
	return &PageHandler{Title: title, Body: body}
}

func (p *PageHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	templates.Execute(resp, p)
}
