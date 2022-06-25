package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-news-api/config"
	"go-news-api/controllers"
	"go-news-api/database"
	"log"
	"net/http"
)

func main() {

	config.LoadAppConfig()

	database.Connect(config.AppConfig.ConnectionString)
	database.Migrate()

	router := mux.NewRouter().StrictSlash(true)

	RegisterProductRoutes(router)

	log.Println(fmt.Sprintf("Starting Server on port %s", config.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), router))

}


func RegisterProductRoutes(router *mux.Router) {
	// NEWS
	router.HandleFunc("/api/v1/news", controllers.GetNews).Methods("GET")
	router.HandleFunc("/api/v1/news", controllers.CreateNews).Methods("POST")

	// COMMENTS
	router.HandleFunc("/api/v1/comments", controllers.GetComments).Methods("GET")
	router.HandleFunc("/api/v1/comments", controllers.CreateComments).Methods("POST")
}