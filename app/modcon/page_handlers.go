package modcon

import (
	"net/http"
	"simpleblog/app/support"
)

type Page struct {
	Title       string
	Bodyclass   string
	MoreStyles  []string
	MoreScripts []string
}

func IndexHandler(context *appContext, resp http.ResponseWriter, req *http.Request) (int, error) {
	p := &Page{
		Title:     "Home",
		Bodyclass: "index-page",
	}
	if err := context.appTemplates.ExecuteTemplate(resp, "indexPage", p); err != nil {
		support.LogStacktrace(err)
		return http.StatusInternalServerError, err
	}
	return 200, nil
}

// func AboutHandler(context *appContext, resp http.ResponseWriter, req *http.Request) (int, error) {
// 	p := &Page{
// 		Title:     "About",
// 		Bodyclass: "about-page",
// 	}
// 	if err := context.appTemplates.ExecuteTemplate(resp, "aboutPage", p); err != nil {
// 		support.LogStacktrace(err)
// 		return http.StatusInternalServerError, err
// 	}
// 	return 200, nil
// }
