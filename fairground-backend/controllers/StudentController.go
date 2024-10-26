package controllers

import (
	"encoding/json"
	"errors"
	"fairground-backend/models"
	"fairground-backend/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ListAllStudents(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get all students from the database
	students, err := models.GetAllStudents()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get students")
		return
	}

	// Step: 2 - Return the students
	utils.JSONDataResponse(w, http.StatusOK, students)
}

func ListParentsByStudentID(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get user_id from the request header
	userID, _ := strconv.Atoi(r.Header.Get("user_id"))

	// Step: 3 - Return Students from student_parent table with the parent ID
	parents, err := models.GetParentsByStudentID(userID)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get students")
		return
	}

	// Step: 4 - Return the parents
	utils.JSONDataResponse(w, http.StatusOK, parents)

}

func DeleteParentFromStudent(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get parent ID from the request parameters using mux
	params := mux.Vars(r)
	parentId := params["parentId"]

	// Step: 2 - Get student ID from the request header
	studentId := r.Header.Get("user_id")

	// Step: 2 - Delete student from parent_student table
	err := models.DeleteStudentFromParent(studentId, parentId)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not delete parent")
		return
	}

	// Step: 3 - Return success message
	utils.JSONResponse(w, http.StatusOK, "Parent removed successfully")
}
func AddParentToStudent(w http.ResponseWriter, r *http.Request) {

	// Step:1 - Check if the request body is valid
	var parent models.Parent
	if err := json.NewDecoder(r.Body).Decode(&parent); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err, "invalid input")
		return
	}

	// Step:2 - Check if the required fields are present
	if parent.Email == "" {
		utils.JSONResponse(w, http.StatusBadRequest, errors.New("missing required fields"))
		return
	}

	// Step:3 - Get student from the database using email
	user_parent, err := models.GetUserByEmail(parent.Email)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get user")
		return
	}

	// Step:4 - If the role is not student, return error
	if user_parent.Role != models.RoleParent {
		utils.JSONResponse(w, http.StatusBadRequest, errors.New("user is not a parent"))
		return
	}

	// Step:4 - Add student to parent_student table
	err = models.AddParentToStudent(strconv.Itoa(user_parent.UserID), r.Header.Get("user_id"))
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Step:5 - Return success message
	utils.JSONResponse(w, http.StatusOK, "Parent added successfully")
}
