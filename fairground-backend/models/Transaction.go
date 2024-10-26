package models

import "fairground-backend/db"

type Transaction struct {
	TransactionID int    `json:"transaction_id"`
	FromUserID    int    `json:"from_user_id"`
	ToUserID      int    `json:"to_user_id"`
	Tokens        int    `json:"tokens"`
	TransferDate  string `json:"transfer_date"`
}

func GetTransactionsByUserID(userID int) ([]Transaction, error) {

	// Step 1: Retrieve all transactions
	rows, err := db.DB.Query("SELECT transaction_id, from_user_id, to_user_id, tokens, transfer_date FROM transactions WHERE from_user_id = $1 OR to_user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Step 2: Iterate over the rows and create a slice of transactions
	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		if err := rows.Scan(&transaction.TransactionID, &transaction.FromUserID, &transaction.ToUserID, &transaction.Tokens, &transaction.TransferDate); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Step 3: Return empty slice if there are no transactions
	if len(transactions) == 0 {
		return []Transaction{}, nil
	}

	// Step 4: Return the slice of transactions
	return transactions, nil
}
