package modcon

import (
	"net/http"
	"simpleblog/app/support"
)

type AppHandler struct {
	Context *appContext
	Handler func(*appContext, http.ResponseWriter, *http.Request) (int, error)
}

func (ah AppHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	status, err := ah.Handler(ah.Context, resp, req)
	if err != nil {
		support.LogStacktrace(err)
		switch status {
		case http.StatusNotFound:
			// replace with custom 404
			http.NotFound(resp, req)
		default:
			http.Error(resp, http.StatusText(status), status)
		}
	}
}
