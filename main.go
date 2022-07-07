package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/luanbe/golang-web-app-structure/app/delivery"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	delivery.NewStaticDelivery(router)
	fmt.Printf("Server START on port%v\n", viper.GetString("server.address"))
	log.Fatal(http.ListenAndServe(viper.GetString("server.address"), router))
}
