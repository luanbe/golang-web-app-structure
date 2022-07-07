package delivery

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luanbe/golang-web-app-structure/helper"
	"github.com/luanbe/golang-web-app-structure/templates"
)

type UserAdminDelivery struct {
	tpl helper.Template
}

func NewUserAdminDelivery(router *httprouter.Router) {
	uad := UserAdminDelivery{}
	router.GET("/admin/user/add", uad.New)
}

func (uad UserAdminDelivery) New(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uad.tpl = helper.TplMust(uad.tpl.TplParseFS(templates.FS, "admin/user/user_add.gohtml"))
	uad.tpl.Execute(w, nil)
}
