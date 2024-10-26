package routes

import (
	"fairground-backend/controllers"
	"github.com/gorilla/mux"
)

// AuthRoutes defines routes for authentication-related endpoints.
func AuthRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/auth/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/v1/auth/login", controllers.Login).Methods("POST")
}
