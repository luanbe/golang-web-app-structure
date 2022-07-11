package initialization

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/luanbe/golang-web-app-structure/app/delivery"
	"github.com/luanbe/golang-web-app-structure/app/delivery/delivery_admin"
	"github.com/luanbe/golang-web-app-structure/app/models/entity"
	"github.com/luanbe/golang-web-app-structure/app/registry"
	"github.com/luanbe/golang-web-app-structure/helper/database"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

//TODO: add logger later
func InitDb() (*gorm.DB, error) {
	db, err := database.NewConnectionDB(
		viper.GetString("database.driver"),
		viper.GetString("database.dbname"),
		viper.GetString("database.host"),
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetInt("database.port"),
	)
	if err != nil {
		return nil, err
	}

	//run drop table to refresh data.
	// db.Migrator().DropTable(&entity.User{})

	// Define auto migration here
	_ = db.AutoMigrate(&entity.User{})

	//seedingPredefined(db, logger)

	return db, nil
}

// TODO: Add logger later
func InitRouting(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Mount("/", fontEndRouter(db))
	r.Mount("/admin", adminRouter(db))

	return r
}

func fontEndRouter(db *gorm.DB) http.Handler {
	// Service registry
	userService := registry.RegisterUserService(db)

	r := chi.NewRouter()
	index := delivery.NewIndexDelivery()
	r.Mount("/", index.Routes())

	user := delivery.NewUserDelivery(userService)
	r.Mount("/users", user.Routes())
	return r
}

func adminRouter(db *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(AdminOnly)
	indexAdmin := delivery_admin.NewIndexAdminDelivery()
	r.Mount("/", indexAdmin.Routes())

	userAdmin := delivery_admin.NewUserAdminDelivery()
	r.Mount("/users", userAdmin.Routes())

	return r
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		isAdmin, ok := ctx.Value("acl.admin").(bool)
		if !ok || !isAdmin {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
