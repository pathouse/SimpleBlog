package modcon

import (
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

// TODO - grab from ENV
var (
	sessionKey string = "5a79dfc4cce73aa90bd47fa649b1f6f6906eac8d9e41e20f8412469c9ec3060f6dd056fa935ef3a21ef9cc31e19005c44c97dc65c1d24994dbbf5f16997da466"
)

func NewAppContext(db *gorm.DB) *appContext {
	// compile templates
	appTemplates := gotemp.New("layout")
	for _, str := range templates.ParseTemplateBinaries() {
		gotemp.Must(appTemplates.Parse(str))
	}

	store := sessions.NewCookieStore([]byte(sessionKey))

	return &appContext{
		db:           db,
		appTemplates: appTemplates,
		store:        store,
	}
}
