package delivery

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luanbe/golang-web-app-structure/helper"
	"github.com/luanbe/golang-web-app-structure/templates"
)

type StaticAdminDelivery struct {
	Tpl helper.Template
}

func NewStaticAdminDelivery(router *httprouter.Router) {
	sad := StaticAdminDelivery{}
	router.GET("/admin", sad.HomePageAdmin)
}

func (sad StaticAdminDelivery) HomePageAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sad.Tpl = helper.TplMust(sad.Tpl.TplParseFS(templates.FS, "admin/home.gohtml"))
	sad.Tpl.Execute(w, nil)
}
