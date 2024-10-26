package models

import "fairground-backend/db"

type StandHolder struct {
	User
}

func GetStandHolderByUserID(userID int) (*StandHolder, error) {

	// Step 1: Retrieve stand holder by user ID
	var standHolder StandHolder
	err := db.DB.QueryRow("SELECT stand_holder_id FROM standholders WHERE stand_holder_id = $1", userID).Scan(&standHolder.UserID)

	if err != nil {
		return nil, err
	}

	// Step 2: Return the stand holder
	return &standHolder, nil

}
