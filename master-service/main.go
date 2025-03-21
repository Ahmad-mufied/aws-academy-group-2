package main

import (
	"fmt"
	"log"
	"master-service/config"
	"master-service/handler"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	app := config.InitializeApp()

	route := mux.NewRouter()

	route.HandleFunc("/", handler.Hello)
	route.HandleFunc("/roles", app.RoleHandler.CreateNewRole).Methods("POST")
	route.HandleFunc("/roles", app.RoleHandler.GetAllRoles).Methods("GET")
	route.HandleFunc("/roles/{id}", app.RoleHandler.GetRoleByID).Methods("GET")
	route.HandleFunc("/status", app.StatusHandler.CreateNewStatus).Methods("POST")
	route.HandleFunc("/status", app.StatusHandler.GetAllStatus).Methods("GET")
	route.HandleFunc("/status/{id}", app.StatusHandler.GetStatusByID).Methods("GET")

	fmt.Println("Server is listening on port 8001...")

	err = http.ListenAndServe(":8001", route)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
