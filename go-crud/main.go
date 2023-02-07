package main

import (
	"fmt"
	"go-crud/controllers"
	"go-crud/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")
}

func main() {

	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterUserRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}
