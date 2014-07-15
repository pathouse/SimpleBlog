package templates

import (
	"html/template"
	"net/http"
	"simpleblog/app/support"
)

var (
	SiteLayout *template.Template = template.New("layout")
)

func init() {
	for _, name := range AssetNames() {
		asset, err := Asset(name)
		if err != nil {
			support.LogStacktrace(err)
		}
		template.Must(SiteLayout.Parse(string(asset[:])))
	}
}

func Execute(resp http.ResponseWriter, args interface{}) {
	if err := SiteLayout.ExecuteTemplate(resp, "layout", args); err != nil {
		support.LogStacktrace(err)
	}
}
