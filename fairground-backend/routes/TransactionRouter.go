package routes

import (
	"fairground-backend/controllers"
	"fairground-backend/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

// UserRoutes defines routes for user-related endpoints.
func TransactionRoutes(router *mux.Router) {

	// Get all transactions
	router.Handle("/api/v1/transactions/me", middlewares.JwtAuthMiddleware(http.HandlerFunc(controllers.ListMyTransactions))).Methods("GET")
}
