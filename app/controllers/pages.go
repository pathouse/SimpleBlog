package controllers

import (
	"html/template"
	"net/http"
	"simpleblog/app/support"
)

const (
	PagesContentPath string = "/Users/patsicle/workspace/go/src/simpleblog/app/templates/pages/content.html"
)

type Page struct {
	Title string
	Body  string
}

func Home(resp http.ResponseWriter, req *http.Request) {
	p := &Page{
		Title: "Simple Blog - Home",
		Body:  "Welcome to my simple blog.",
	}

	templates := template.Must(template.ParseFiles(PagesContentPath, LayoutPath))
	if err := templates.ExecuteTemplate(resp, "layout", p); err != nil {
		support.LogStacktrace(err)
	}
}
