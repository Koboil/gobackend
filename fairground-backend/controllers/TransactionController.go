package controllers

import (
	"fairground-backend/models"
	"fairground-backend/utils"
	"net/http"
	"strconv"
)

func ListMyTransactions(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get user_id from the request header
	userID, _ := strconv.Atoi(r.Header.Get("user_id"))

	// Step: 2 - Get Transactions from the database
	transactions, err := models.GetTransactionsByUserID(userID)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get transactions")
		return
	}

	// Step: 3 - Return the transactions
	utils.JSONDataResponse(w, http.StatusOK, transactions)

}
