package delivery

import (
	"github.com/julienschmidt/httprouter"
	"github.com/luanbe/golang-web-app-structure/helper"
	"github.com/luanbe/golang-web-app-structure/templates"
	"net/http"
)

type StaticDelivery struct {
	Tpl helper.Template
}

type StaticAdminDelivery struct {
	Tpl helper.Template
}

func NewStaticDelivery(router *httprouter.Router) {
	sd := StaticDelivery{helper.TplMust(helper.TplParseFS(templates.FS, "home.gohtml"))}
	sad := StaticAdminDelivery{helper.TplMust(helper.TplParseFS(templates.FS, "admin/home.gohtml"))}

	router.GET("/", sd.HomePage)
	router.GET("/admin", sad.HomePageAdmin)
}

func (sd StaticDelivery) HomePage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sd.Tpl.Execute(w, nil)
}

func (sad StaticAdminDelivery) HomePageAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sad.Tpl.Execute(w, nil)
}
