package routes

import (
	"fairground-backend/controllers"
	"fairground-backend/middlewares"
	"fairground-backend/models"
	"github.com/gorilla/mux"
	"net/http"
)

// StudentRoutes defines routes for student-related endpoints.
func StudentRoutes(router *mux.Router) {

	// Get all students
	router.Handle("/api/v1/students", middlewares.JwtAuthMiddleware(http.HandlerFunc(controllers.ListAllStudents))).Methods("GET")

	// Get Parent's Student
	router.Handle("/api/v1/students/parents",
		middlewares.JwtAuthMiddleware(
			middlewares.RoleAuthMiddleware([]string{string(models.RoleStudent)},
				http.HandlerFunc(controllers.ListParentsByStudentID)))).
		Methods("GET")

	// Delete a parent from student
	router.Handle("/api/v1/students/parent/{parentId}",
		middlewares.JwtAuthMiddleware(
			middlewares.RoleAuthMiddleware([]string{string(models.RoleStudent)},
				http.HandlerFunc(controllers.DeleteParentFromStudent)))).
		Methods("DELETE")

	// Add a parent to a student using email
	router.Handle("/api/v1/students/parent",
		middlewares.JwtAuthMiddleware(
			middlewares.RoleAuthMiddleware([]string{string(models.RoleStudent)},
				http.HandlerFunc(controllers.AddParentToStudent)))).
		Methods("POST")

}
