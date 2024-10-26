package controllers

import (
	"encoding/json"
	"errors"
	"fairground-backend/models"
	"fairground-backend/utils"
	"net/http"
	"strconv"
)

func ListAllTokens(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get all tokens from the database
	tokens, err := models.GetAllTokens()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get tokens")
		return
	}

	// Step: 2 - Return the tokens
	utils.JSONDataResponse(w, http.StatusOK, tokens)
}

func ListParentsTokens(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get email from the request header
	email := r.Header.Get("email")

	// Step: 2 - Get User from the database
	user, err := models.GetUserByEmail(email)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get user")
		return
	}

	// Step: 3 - Return Tokens with the user ID
	tokens, err := models.GetTokensByParentID(user.UserID)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get tokens")
		return
	}

	// Step: 4 - Return the tokens
	utils.JSONDataResponse(w, http.StatusOK, tokens)
}

func PurchaseToken(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get email from the request header
	email := r.Header.Get("email")

	// Step: 2 - Check if the request body is valid
	var token models.Token
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err, "invalid input")
		return
	}

	// Step: 3 - Check if all the required fields are present
	if token.Amount == 0 {
		utils.JSONResponse(w, http.StatusBadRequest, errors.New("missing required fields"))
		return
	}

	// Step: 4 - Get User from the database
	user, err := models.GetUserByEmail(email)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get user")
		return
	}

	// Step: 5 - Create Token
	token.PurchasedBy = user.UserID
	_, err = models.Purchase(&token)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not create token")
		return
	}

	// Step: 6 - Return the token
	utils.JSONDataResponse(w, http.StatusOK, token)
}

func TransferToken(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get user_id from the request header
	var transaction models.Transaction
	userID := r.Header.Get("user_id")
	transaction.FromUserID, _ = strconv.Atoi(userID)

	// Step: 2 - Check if the request body is valid
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err, "invalid input")
		return
	}

	// Step: 3 - Check if all the required fields are present
	if transaction.FromUserID == 0 || transaction.ToUserID == 0 || transaction.Tokens == 0 {
		utils.JSONResponse(w, http.StatusBadRequest, errors.New("missing required fields"))
		return
	}

	// Step: 4 - Transfer tokens
	_, err := models.Transfer(&transaction)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Step: 5 - Return the transaction
	utils.JSONDataResponse(w, http.StatusOK, transaction)

}
