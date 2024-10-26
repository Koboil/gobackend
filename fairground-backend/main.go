package main

import (
	"fairground-backend/config"
	"fairground-backend/db"
	"fairground-backend/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// Connect to the database
	db.ConnectDB()

	// Create a new mux router
	router := mux.NewRouter()

	// Register routes
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.TokenRoutes(router)
	routes.ParentRoutes(router)
	routes.StudentRoutes(router)
	routes.TransactionRoutes(router)
	routes.RaffleRotuer(router)
	routes.RaffleTicketRoutes(router)

	// Start the server and pass the router
	log.Println("Starting server on :" + config.CONFIG.BACKEND_PORT)
	if err := http.ListenAndServe(":"+config.CONFIG.BACKEND_PORT, router); err != nil {
		log.Fatal(err)
	}
}
