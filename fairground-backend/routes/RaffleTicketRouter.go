package routes

import (
	"fairground-backend/controllers"
	"fairground-backend/middlewares"
	"fairground-backend/models"
	"github.com/gorilla/mux"
	"net/http"
)

// RaffleTicketRoutes defines routes for token-related endpoints.
func RaffleTicketRoutes(router *mux.Router) {

	// List particular currents users raffle tickets
	router.Handle("/api/v1/raffle/tickets/me",
		middlewares.JwtAuthMiddleware(
			middlewares.RoleAuthMiddleware([]string{string(models.RoleParent), string(models.RoleStudent)},
				http.HandlerFunc(controllers.ListMyRaffleTickets)))).
		Methods("GET")

	// Purchase Raffle Ticket
	router.Handle("/api/v1/raffle/ticket/purchase/{raffleId}",
		middlewares.JwtAuthMiddleware(
			middlewares.RoleAuthMiddleware([]string{string(models.RoleParent), string(models.RoleStudent)},
				http.HandlerFunc(controllers.PurchaseRaffleTicket)))).
		Methods("POST")

}
