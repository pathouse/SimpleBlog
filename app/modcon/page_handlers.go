package modcon

import (
	"net/http"
	"simpleblog/app/support"
)

type PageHandler struct {
	context *appContext
	Title   string
	Body    string
}

func NewPageHandler(context *appContext, title, body string) http.Handler {
	return &PageHandler{
		context: context,
		Title:   title,
		Body:    body}
}

func (p *PageHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if err := p.context.appTemplates.ExecuteTemplate(resp, "layout", p); err != nil {
		support.LogStacktrace(err)
	}
}
