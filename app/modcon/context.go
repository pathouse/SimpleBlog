package modcon

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	gotemp "html/template"
	"simpleblog/app/templates"
)

type appContext struct {
	db           *gorm.DB
	appTemplates *gotemp.Template
	store        sessions.CookieStore
}

func NewAppContext(db *gorm.DB) *appContext {
	// compile templates
	appTemplates := gotemp.New("layout")
	for _, str := range templates.ParseTemplateBinaries() {
		gotemp.Must(appTemplates.Parse(str))
	}

	return &appContext{
		db:           db,
		appTemplates: appTemplates,
	}
}
