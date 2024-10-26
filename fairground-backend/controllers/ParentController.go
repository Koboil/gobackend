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

func ListAllParents(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get all parents from the database
	parents, err := models.GetAllParents()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get parents")
		return
	}

	// Step: 2 - Return the parents
	utils.JSONDataResponse(w, http.StatusOK, parents)
}

func ListStudentsByParentID(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get email from the request header
	email := r.Header.Get("email")

	// Step: 2 - Get User from the database
	user, err := models.GetUserByEmail(email)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get user")
		return
	}

	// Step: 3 - Return Students from student_parent table with the parent ID
	students, err := models.GetStudentsByParentID(user.UserID)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get students")
		return
	}

	// Step: 4 - Return the students
	utils.JSONDataResponse(w, http.StatusOK, students)

}

func DeleteStudentFromParent(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get student ID from the request parameters using mux
	params := mux.Vars(r)
	studentId := params["studentId"]

	// Step: 2 - Get parent ID from the request header
	parentId := r.Header.Get("user_id")

	// Step: 2 - Delete student from parent_student table
	err := models.DeleteStudentFromParent(studentId, parentId)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not delete student")
		return
	}

	// Step: 3 - Return success message
	utils.JSONResponse(w, http.StatusOK, "Student removed successfully")
}

func AddStudentToParent(w http.ResponseWriter, r *http.Request) {

	// Step:1 - Check if the request body is valid
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err, "invalid input")
		return
	}

	// Step:2 - Check if the required fields are present
	if student.Email == "" {
		utils.JSONResponse(w, http.StatusBadRequest, errors.New("missing required fields"))
		return
	}

	// Step:3 - Get student from the database using email
	user_student, err := models.GetUserByEmail(student.Email)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get user")
		return
	}

	// Step:4 - If the role is not student, return error
	if user_student.Role != models.RoleStudent {
		utils.JSONResponse(w, http.StatusBadRequest, errors.New("user is not a student"))
		return
	}

	// Step:4 - Add student to parent_student table
	err = models.AddStudentToParent(strconv.Itoa(user_student.UserID), r.Header.Get("user_id"))
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Step:5 - Return success message
	utils.JSONResponse(w, http.StatusOK, "Student added successfully")
}

func GetStudentByUserID(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get student ID from the request parameters using mux
	params := mux.Vars(r)
	studentId, _ := strconv.Atoi(params["studentId"])

	// Step: 2 - Get user from the database
	student, err := models.GetStudentByUserID(studentId)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get student")
		return
	}

	// Step: 3 - Get User from the database
	user, err := models.GetUserByID(strconv.Itoa(student.UserID))
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get user")
		return
	}

	// Step: 4 - Remove the password from the user
	user.Password = ""

	// Step: 5 - student.User = user
	student.User = *user

	// Step: 5 - Return the student
	utils.JSONDataResponse(w, http.StatusOK, student)

}
