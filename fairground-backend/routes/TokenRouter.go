package routes

import (
	"fairground-backend/controllers"
	"fairground-backend/middlewares"
	"fairground-backend/models"
	"github.com/gorilla/mux"
	"net/http"
)

// TokenRoutes defines routes for token-related endpoints.
func TokenRoutes(router *mux.Router) {
	// Define the allowed roles for the Token API

	// List all token history(admin only)
	router.Handle("/api/v1/tokens",
		middlewares.JwtAuthMiddleware(
			middlewares.RoleAuthMiddleware([]string{string(models.RoleFairgroundManagement)},
				http.HandlerFunc(controllers.ListAllTokens)))).
		Methods("GET")

	// List particular parents token history
	router.Handle("/api/v1/tokens/me",
		middlewares.JwtAuthMiddleware(
			middlewares.RoleAuthMiddleware([]string{string(models.RoleParent)},
				http.HandlerFunc(controllers.ListParentsTokens)))).
		Methods("GET")

	// Purchase token
	router.Handle("/api/v1/tokens/purchase",
		middlewares.JwtAuthMiddleware(
			middlewares.RoleAuthMiddleware([]string{string(models.RoleParent)},
				http.HandlerFunc(controllers.PurchaseToken)))).
		Methods("POST")

	// Transfer token
	router.Handle("/api/v1/tokens/transfer",
		middlewares.JwtAuthMiddleware(
			middlewares.RoleAuthMiddleware([]string{string(models.RoleParent)},
				http.HandlerFunc(controllers.TransferToken)))).
		Methods("POST")
}
