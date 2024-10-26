package controllers

import (
	"fairground-backend/models"
	"fairground-backend/utils"
	"net/http"
)

func ListAllRaffles(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get all raffles from the database
	raffles, err := models.GetAllRaffles()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get raffles")
		return
	}

	// Step: 2 - Return the raffles
	utils.JSONDataResponse(w, http.StatusOK, raffles)

}
