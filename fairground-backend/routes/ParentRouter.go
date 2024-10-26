package routes

import (
	"fairground-backend/controllers"
	"fairground-backend/middlewares"
	"fairground-backend/models"
	"github.com/gorilla/mux"
	"net/http"
)

// ParentRoutes defines routes for parent-related endpoints.
func ParentRoutes(router *mux.Router) {

	// Get all parents
	router.Handle("/api/v1/parents", middlewares.JwtAuthMiddleware(http.HandlerFunc(controllers.ListAllParents))).Methods("GET")

	// Get Parent's Student
	router.Handle("/api/v1/parents/students",
		middlewares.JwtAuthMiddleware(
			middlewares.RoleAuthMiddleware([]string{string(models.RoleParent)},
				http.HandlerFunc(controllers.ListStudentsByParentID)))).
		Methods("GET")

	// Delete a student from a parent
	router.Handle("/api/v1/parents/student/{studentId}",
		middlewares.JwtAuthMiddleware(
			middlewares.RoleAuthMiddleware([]string{string(models.RoleParent)},
				http.HandlerFunc(controllers.DeleteStudentFromParent)))).
		Methods("DELETE")

	// Add a student to a parent using email
	router.Handle("/api/v1/parents/student",
		middlewares.JwtAuthMiddleware(
			middlewares.RoleAuthMiddleware([]string{string(models.RoleParent)},
				http.HandlerFunc(controllers.AddStudentToParent)))).
		Methods("POST")

	// Get a particular student
	router.Handle("/api/v1/parents/student/{studentId}",
		middlewares.JwtAuthMiddleware(
			middlewares.RoleAuthMiddleware([]string{string(models.RoleParent)},
				http.HandlerFunc(controllers.GetStudentByUserID)))).
		Methods("GET")

}
