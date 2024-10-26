package models

import (
	"fairground-backend/db"
	"strconv"
)

type RaffleTicket struct {
	RaffleTicketID int    `json:"raffle_ticket_id"`
	PurchaseDate   string `json:"purchase_date"`
	UserID         int    `json:"user_id"`
	RaffleID       int    `json:"raffle_id"`
}

func GetRaffleTicketsByUserId(userId string) ([]RaffleTicket, error) {

	// Step 1: Retrieve all raffle tickets
	rows, err := db.DB.Query("SELECT raffle_ticket_id, purchase_date, user_id, raffle_id FROM raffletickets WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Step 2: Iterate over the rows and create a slice of raffle tickets
	var raffleTickets []RaffleTicket
	for rows.Next() {
		var raffleTicket RaffleTicket
		if err := rows.Scan(&raffleTicket.RaffleTicketID, &raffleTicket.PurchaseDate, &raffleTicket.UserID, &raffleTicket.RaffleID); err != nil {
			return nil, err
		}
		raffleTickets = append(raffleTickets, raffleTicket)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Step 3: If there are no raffle tickets, return an empty slice
	if len(raffleTickets) == 0 {
		return []RaffleTicket{}, nil
	}

	// Step 4: Return the slice of raffle tickets
	return raffleTickets, nil
}

func PurchaseRaffleTicket(userId string, raffleId string) error {
	// Start a new transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	// Rollback transaction in case of an error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Step 1: Get Raffle by ID
	raffle, err := GetRaffleById(raffleId)
	if err != nil {
		return err
	}

	//// Step 2: Check if the raffle is still active by the end date
	//endDate, err := time.Parse("2006-01-02 15:04:05", raffle.EndDate)
	//if err != nil {
	//	return err
	//}
	//if endDate.Before(time.Now()) {
	//	return errors.New("raffle has ended")
	//}

	// Step 3: Create Raffle Ticket
	_, err = tx.Exec("INSERT INTO raffletickets (user_id, raffle_id) VALUES ($1, $2)", userId, raffleId)
	if err != nil {
		return err
	}

	// Step 4: Prepare a transaction struct for token transfer
	var transaction Transaction
	transaction.Tokens = raffle.TicketPrice
	transaction.FromUserID, _ = strconv.Atoi(userId)
	transaction.ToUserID = raffle.OrganizerID

	// Step 5: Transfer tokens by inserting into transactions table within the transaction
	_, err = Transfer(&transaction)
	if err != nil {
		return err
	}

	// Step 6: Commit the transaction
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
