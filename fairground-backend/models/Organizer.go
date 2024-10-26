package models

import "fairground-backend/db"

type Organizer struct {
	User
	TokensBalance int `json:"tokens_balance"`
}

func GetOrganizerByUserID(userID int) (*Organizer, error) {
	// Step 1: Retrieve organizer by user ID
	var organizer Organizer
	err := db.DB.QueryRow("SELECT organizer_id, tokens_balance FROM organizers WHERE organizer_id = $1", userID).Scan(&organizer.UserID, &organizer.TokensBalance)
	if err != nil {
		return nil, err
	}

	// Step 2: Return the organizer
	return &organizer, nil
}
