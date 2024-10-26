package controllers

import (
	"errors"
	"fairground-backend/models"
	"fairground-backend/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func ListAllUsers(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get all users from the database
	users, err := models.GetAllUsers()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get users")
		return
	}

	// Step: 2 - Return the users
	utils.JSONDataResponse(w, http.StatusOK, users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get the user ID from the request parameters using mux
	params := mux.Vars(r)
	userID := params["id"]

	// Step: 2 - Get the user from the database
	user, err := models.GetUserByID(userID)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get user")
		return
	}

	// Step: 3 - Return the user
	utils.JSONDataResponse(w, http.StatusOK, user)
}

func GetMe(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get the email from the request header
	email := r.Header.Get("email")

	// Step: 2 - Get the user from the database
	user, err := models.GetUserByEmail(email)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get user")
		return
	}

	// Step: 3 - Remove the password from the user
	user.Password = ""

	// Step: 4 - If the role is Parent, get the parent information
	if user.Role == models.RoleParent {

		parent, err := models.GetParentByUserID(user.UserID)
		if err != nil {
			utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get parent")
			return
		}

		parent.User = *user

		utils.JSONDataResponse(w, http.StatusOK, parent)
		return
	}

	// Step: 5 - If the role is Student, get the student information
	if user.Role == models.RoleStudent {

		student, err := models.GetStudentByUserID(user.UserID)
		if err != nil {
			utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get student")
			return
		}

		student.User = *user

		utils.JSONDataResponse(w, http.StatusOK, student)
		return
	}

	// Step: 6 - If the role is Organizer, get the organizer information
	if user.Role == models.RoleOrganizer {

		organizer, err := models.GetOrganizerByUserID(user.UserID)
		if err != nil {
			utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get organizer")
			return
		}

		organizer.User = *user

		utils.JSONDataResponse(w, http.StatusOK, organizer)
		return
	}

	// Step: 7 - If the role is FairgroundManagement, get the fairgroundmanagement information
	if user.Role == models.RoleFairgroundManagement {

		fairgroundmanagement, err := models.GetFairgroundManagementByUserID(user.UserID)
		if err != nil {
			utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get fairgroundmanagement")
			return
		}

		fairgroundmanagement.User = *user

		utils.JSONDataResponse(w, http.StatusOK, fairgroundmanagement)
		return
	}

	// Step: 8 - If the role is StandHolder, get the standholder information
	if user.Role == models.RoleStandHolder {

		standholder, err := models.GetStandHolderByUserID(user.UserID)
		if err != nil {
			utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get standholder")
			return
		}

		standholder.User = *user

		utils.JSONDataResponse(w, http.StatusOK, standholder)
		return
	}

	// Step: 4 - Return Error
	utils.JSONResponse(w, http.StatusInternalServerError, errors.New("could not get user"))

}
