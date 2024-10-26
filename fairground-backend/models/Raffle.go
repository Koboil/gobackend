package models

import "fairground-backend/db"

type Raffle struct {
	RaffleID     int    `json:"raffle_id"`     // Primary key for the raffle
	OrganizerID  int    `json:"organizer_id"`  // Foreign key to Organizer
	RaffleName   string `json:"raffle_name"`   // Name of the raffle
	TicketPrice  int    `json:"ticket_price"`  // Cost of each ticket
	PrizeDetails string `json:"prize_details"` // Prize description
	StartDate    string `json:"start_date"`    // Start date of the raffle
	EndDate      string `json:"end_date"`      // End date of the raffle
}

func GetAllRaffles() ([]Raffle, error) {
	// Step 1: Retrieve all raffles
	rows, err := db.DB.Query("SELECT raffle_id, organizer_id, raffle_name, ticket_price, prize_details, start_date, end_date FROM raffles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Step 2: Iterate over the rows and create a slice of raffles
	var raffles []Raffle
	for rows.Next() {
		var raffle Raffle
		if err := rows.Scan(&raffle.RaffleID, &raffle.OrganizerID, &raffle.RaffleName, &raffle.TicketPrice, &raffle.PrizeDetails, &raffle.StartDate, &raffle.EndDate); err != nil {
			return nil, err
		}
		raffles = append(raffles, raffle)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Step 3: If there are no raffles, return an empty slice
	if len(raffles) == 0 {
		return []Raffle{}, nil
	}

	// Step 3: Return the slice of raffles
	return raffles, nil
}

func GetRaffleById(id string) (*Raffle, error) {
	// Step 1: Retrieve raffle by ID
	var raffle Raffle
	err := db.DB.QueryRow("SELECT raffle_id, organizer_id, raffle_name, ticket_price, prize_details, start_date, end_date FROM raffles WHERE raffle_id = $1", id).Scan(&raffle.RaffleID, &raffle.OrganizerID, &raffle.RaffleName, &raffle.TicketPrice, &raffle.PrizeDetails, &raffle.StartDate, &raffle.EndDate)
	if err != nil {
		return nil, err
	}

	// Step 2: Return the raffle
	return &raffle, nil

}
