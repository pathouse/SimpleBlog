package modcon

import (
	"github.com/jinzhu/gorm"
	gotemp "html/template"
	"simpleblog/app/templates"
)

type appContext struct {
	db           *gorm.DB
	appTemplates *gotemp.Template
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
