package models

import "fairground-backend/db"

type FairgroundManagement struct {
	User
}

func GetFairgroundManagementByUserID(userID int) (*FairgroundManagement, error) {
	// Step 1: Retrieve fairground management by user ID
	var fairgroundManagement FairgroundManagement
	err := db.DB.QueryRow("SELECT management_id FROM fairgroundmanagement WHERE management_id = $1", userID).Scan(&fairgroundManagement.UserID)
	if err != nil {
		return nil, err
	}

	// Step 2: Return the fairground management
	return &fairgroundManagement, nil
}
