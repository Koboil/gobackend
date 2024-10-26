package controllers

import (
	"encoding/json"
	"errors"
	"fairground-backend/models"
	"fairground-backend/utils"
	"net/http"
	"strconv"
)

// Register handles the user registration process.
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Step: 1 - Check if the request body is valid
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err, "invalid input")
		return
	}

	// Step: 2 - Check if the required fields are present
	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" || user.Role == "" {
		utils.JSONResponse(w, http.StatusBadRequest, errors.New("missing required fields"))
		return
	}

	// Step: 3 - Check if the role is valid
	if !utils.IsValidRole(user.Role) {
		utils.JSONResponse(w, http.StatusBadRequest, errors.New("invalid role"))
		return
	}

	// Step: 4 - Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not hash password")
		return
	}
	user.Password = hashedPassword

	// Step: 5 - Create user in the database
	if err := models.CreateUser(&user); err != nil {
		if err.Error() == "user already exists" {
			utils.JSONResponse(w, http.StatusConflict, err)
		} else {
			utils.JSONResponse(w, http.StatusInternalServerError, err)
		}
		return
	}

	// Step: 6 - Get the user from the database
	var storedUser *models.User

	storedUser, err = models.GetUserByEmail(user.Email)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get user")
		return
	}

	// Step: 6 - Generate token
	token, err := utils.GenerateJWT(strconv.Itoa(storedUser.UserID), user.Email, string(user.Role))
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not generate token")
		return
	}

	// Step: 7 - Return token
	w.Header().Set("Authorization", token)
	utils.JSONResponse(w, http.StatusCreated, map[string]string{"token": token})

}

// Login handles the user login process.
func Login(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Check if the request body is valid
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err, "invalid input")
		return
	}

	// Step: 2 - Get user from the database
	storedUser, err := models.GetUserByEmail(user.Email)
	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, errors.New("invalid credentials"))
		return
	}

	// Step: 3 - Check if the password is correct
	isValidPassword := utils.CheckPasswordHash(user.Password, storedUser.Password)
	if !isValidPassword {
		utils.JSONResponse(w, http.StatusUnauthorized, errors.New("invalid credentials"))
		return
	}

	// Step: 4 - Generate token
	token, err := utils.GenerateJWT(strconv.Itoa(storedUser.UserID), storedUser.Email, string(storedUser.Role))
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not generate token")
		return
	}

	// Step: 5 - Return token
	w.Header().Set("Authorization", token)
	utils.JSONResponse(w, http.StatusOK, map[string]string{"token": token})
}
