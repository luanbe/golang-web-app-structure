package delivery

import (
	"fmt"
	"github.com/luanbe/golang-web-app-structure/app/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luanbe/golang-web-app-structure/helper"
	"github.com/luanbe/golang-web-app-structure/templates"
)

type UserDelivery struct {
	Tpl     helper.Template
	Service service.UserService
}

func NewUserDelivery(router *httprouter.Router, s service.UserService) {
	ud := UserDelivery{Service: s}
	router.GET("/signup", ud.Signup)
	router.POST("/users", ud.NewUser)
}

func (ud UserDelivery) Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	ud.Tpl = helper.TplMust(ud.Tpl.TplParseFS(templates.FS, "signup.gohtml"))
	ud.Tpl.Execute(w, nil)
}

func (ud UserDelivery) NewUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := struct {
		Email    string
		UserName string
	}{
		r.FormValue("email"),
		r.FormValue("email"),
	}
	result := ud.Service.AddUser(user.UserName, user.Email)
	fmt.Fprint(w, "New user:", result)

}
