package delivery

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luanbe/golang-web-app-structure/helper"
	"github.com/luanbe/golang-web-app-structure/templates"
)

type UserDelivery struct {
	Tpl helper.Template
}

func NewUserDelivery(router *httprouter.Router) {
	ud := UserDelivery{}
	router.GET("/signup", ud.Signup)
	// router.POST("/users", ud.New)
}

func (ud UserDelivery) Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	ud.Tpl = helper.TplMust(ud.Tpl.TplParseFS(templates.FS, "signup.gohtml"))
	ud.Tpl.Execute(w, nil)
}

// func (ud UserDelivery) New(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	fmt.Fprint(w, "this is test new user")
// }
