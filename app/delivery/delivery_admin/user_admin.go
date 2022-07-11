package delivery_admin

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/luanbe/golang-web-app-structure/helper"
	"github.com/luanbe/golang-web-app-structure/templates"
)

type UserAdminDelivery struct {
	tpl helper.Template
}

func NewUserAdminDelivery() *UserAdminDelivery {
	return &UserAdminDelivery{}
}

func (uad *UserAdminDelivery) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/add", uad.NewUser)

	return r
}

func (uad *UserAdminDelivery) NewUser(w http.ResponseWriter, r *http.Request) {
	uad.tpl = helper.TplMust(uad.tpl.TplParseFS(templates.FS, "admin/user/user_add.gohtml"))
	uad.tpl.Execute(w, nil)
}
