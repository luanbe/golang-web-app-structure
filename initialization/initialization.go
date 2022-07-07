package initialization

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luanbe/golang-web-app-structure/app/delivery"
	"github.com/luanbe/golang-web-app-structure/app/models/entity"
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
func InitRouting(db *gorm.DB) *httprouter.Router {
	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
	// Service registry

	// Routing to deliveries with other paths
	// Front End site
	delivery.NewStaticDelivery(router)
	delivery.NewUserDelivery(router)

	// API site
	// deliver.NewStaticEndpointDelivery(router)

	// Admin site
	delivery.NewStaticAdminDelivery(router)
	delivery.NewUserAdminDelivery(router)

	return router
}
