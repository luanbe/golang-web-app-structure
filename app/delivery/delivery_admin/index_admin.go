package delivery_admin

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/luanbe/golang-web-app-structure/helper"
	"github.com/luanbe/golang-web-app-structure/templates"
)

type IndexAdminDelivery struct {
	Tpl helper.Template
}

func NewIndexAdminDelivery() *IndexAdminDelivery {
	return &IndexAdminDelivery{}
}

func (iad *IndexAdminDelivery) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", iad.HomeAdmin)

	return r
}

func (iad *IndexAdminDelivery) HomeAdmin(w http.ResponseWriter, r *http.Request) {
	iad.Tpl = helper.TplMust(iad.Tpl.TplParseFS(templates.FS, "admin/home.gohtml"))
	iad.Tpl.Execute(w, nil)
}
