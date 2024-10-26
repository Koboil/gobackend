package routes

import (
	"fairground-backend/controllers"
	"fairground-backend/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

// TokenRoutes defines routes for token-related endpoints.
func RaffleRotuer(router *mux.Router) {
	// Define the allowed roles for the Token API

	// List all raffles
	router.Handle("/api/v1/raffles", middlewares.JwtAuthMiddleware(http.HandlerFunc(controllers.ListAllRaffles))).Methods("GET")

}
