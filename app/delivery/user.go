package delivery

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/luanbe/golang-web-app-structure/app/service"
	"github.com/luanbe/golang-web-app-structure/templates"

	"github.com/luanbe/golang-web-app-structure/helper"
)

type UserDelivery struct {
	Tpl     helper.Template
	Service service.UserService
}

func NewUserDelivery(s service.UserService) *UserDelivery {
	return &UserDelivery{Service: s}
}

func (ud *UserDelivery) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/signup", ud.Signup)
	r.Post("/", ud.NewUser)

	return r
}

func (ud *UserDelivery) Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	ud.Tpl = helper.TplMust(ud.Tpl.TplParseFS(templates.FS, "signup.gohtml"))
	ud.Tpl.Execute(w, nil)
}

func (ud *UserDelivery) NewUser(w http.ResponseWriter, r *http.Request) {
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
