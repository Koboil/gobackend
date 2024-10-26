package controllers

import (
	"fairground-backend/models"
	"fairground-backend/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func ListMyRaffleTickets(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get user ID from the request header
	userId := r.Header.Get("user_id")

	// Step: 2 - Get raffle tickets from the database
	raffleTickets, err := models.GetRaffleTicketsByUserId(userId)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err, "could not get raffle tickets")
		return
	}

	// Step: 3 - Return the raffle tickets
	utils.JSONDataResponse(w, http.StatusOK, raffleTickets)

}

func PurchaseRaffleTicket(w http.ResponseWriter, r *http.Request) {

	// Step: 1 - Get user ID from the request header
	userId := r.Header.Get("user_id")

	// Step: 2 - Get raffle ID from the request parameters using mux
	params := mux.Vars(r)
	raffleId := params["raffleId"]

	// Step: 3 - Purchase raffle ticket
	err := models.PurchaseRaffleTicket(userId, raffleId)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Step: 4 - Return success
	utils.JSONResponse(w, http.StatusOK, "raffle ticket purchased")
}
