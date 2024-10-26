package routes

import (
	"fairground-backend/controllers"
	"fairground-backend/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

// UserRoutes defines routes for user-related endpoints.
func UserRoutes(router *mux.Router) {
	// Get all users
	router.Handle("/api/v1/users", middlewares.JwtAuthMiddleware(http.HandlerFunc(controllers.ListAllUsers))).Methods("GET")

	// Get user by ID
	router.Handle("/api/v1/users/{id}", middlewares.JwtAuthMiddleware(http.HandlerFunc(controllers.GetUserByID))).Methods("GET")

	// Get me
	router.Handle("/api/v1/user/me", middlewares.JwtAuthMiddleware(http.HandlerFunc(controllers.GetMe))).Methods("GET")

	// Stats
	//router.Handle("/api/v1/stats",
	//	middlewares.JwtAuthMiddleware(
	//		middlewares.RoleAuthMiddleware([]string{string(models.RoleFairgroundManagement)},
	//			http.HandlerFunc(controllers.ListStats)))).
	//	Methods("GET")
}
