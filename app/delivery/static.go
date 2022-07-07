package delivery

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luanbe/golang-web-app-structure/helper"
	"github.com/luanbe/golang-web-app-structure/templates"
)

type StaticDelivery struct {
	Tpl helper.Template
}

func NewStaticDelivery(router *httprouter.Router) {
	sd := StaticDelivery{}
	router.GET("/", sd.HomePage)

}

func (sd StaticDelivery) HomePage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sd.Tpl = helper.TplMust(sd.Tpl.TplParseFS(templates.FS, "home.gohtml"))
	sd.Tpl.Execute(w, nil)
}
