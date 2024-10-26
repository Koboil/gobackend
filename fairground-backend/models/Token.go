package models

import (
	"errors"
	"fairground-backend/db"
	"fmt"
	"strconv"
	"time"
)

type Token struct {
	TokenId      string `json:"token_id"`
	Amount       int    `json:"amount"`
	PurchasedBy  int    `json:"purchased_by"`
	PurchaseDate string `json:"purchase_date"`
}

func GetAllTokens() ([]Token, error) {
	// Step 1: Retrieve all tokens
	rows, err := db.DB.Query("SELECT token_id, amount, purchased_by, purchase_date FROM tokens")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Step 2: Iterate over the rows and create a slice of tokens
	var tokens []Token
	for rows.Next() {
		var token Token
		if err := rows.Scan(&token.TokenId, &token.Amount, &token.PurchasedBy, &token.PurchaseDate); err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Step 3: Return the slice of tokens
	return tokens, nil
}

func GetTokensByParentID(parentID int) ([]Token, error) {
	// Step 1: Retrieve all tokens
	rows, err := db.DB.Query("SELECT token_id, amount, purchased_by, purchase_date FROM tokens WHERE purchased_by = $1", parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Step 2: Iterate over the rows and create a slice of tokens
	var tokens []Token
	for rows.Next() {
		var token Token
		if err := rows.Scan(&token.TokenId, &token.Amount, &token.PurchasedBy, &token.PurchaseDate); err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Step 3: if there are no tokens, return an empty slice
	if len(tokens) == 0 {
		return []Token{}, nil
	}

	// Step 5: Return the slice of tokens
	return tokens, nil
}

func Purchase(token *Token) (*Token, error) {

	// Step 1: Begin transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Step 2: Add Purchase Date
	token.PurchaseDate = time.Now().Format("2006-01-02 15:04:05")

	// Step 3: Insert token into the database
	_, err = tx.Exec("INSERT INTO tokens (amount, purchased_by, purchase_date) VALUES ($1, $2, $3)", token.Amount, token.PurchasedBy, token.PurchaseDate)
	if err != nil {
		return nil, err
	}

	// Step 4: Update the parent's balance
	_, err = tx.Exec("UPDATE parents SET tokens_balance = tokens_balance + $1 WHERE parent_id = $2", token.Amount, token.PurchasedBy)
	if err != nil {
		return nil, err
	}

	// Step 5: Commit transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// Step 6: Return the token
	return token, nil
}

func Transfer(transaction *Transaction) (*Transaction, error) {

	// Step 1: Begin transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Step 2: Get details of sender
	sender, _ := GetUserByID(strconv.Itoa(transaction.FromUserID))
	var sender_table, sender_column string

	if sender.Role == RoleParent {
		sender_table = "parents"
		sender_column = "parent_id"
	} else {
		sender_table = "students"
		sender_column = "student_id"
	}

	// Step 3: Get details of receiver
	receiver, _ := GetUserByID(strconv.Itoa(transaction.ToUserID))
	var receiver_table, receiver_column string

	if receiver.Role == RoleOrganizer {
		receiver_table = "organizers"
		receiver_column = "organizer_id"
	} else {
		receiver_table = "students"
		receiver_column = "student_id"
	}

	// Step 4: Check if the sender has enough tokens
	var balance int
	query := fmt.Sprintf("SELECT tokens_balance FROM %s WHERE %s = $1", sender_table, sender_column)
	err = tx.QueryRow(query, transaction.FromUserID).Scan(&balance)
	if err != nil {
		return nil, err
	}
	if balance < transaction.Tokens {
		return nil, errors.New("not enough tokens")
	}

	// Step 5: Update the sender's balance
	updateSender := fmt.Sprintf("UPDATE %s SET tokens_balance = tokens_balance - $1 WHERE %s = $2", sender_table, sender_column)
	_, err = tx.Exec(updateSender, transaction.Tokens, transaction.FromUserID)
	if err != nil {
		return nil, err
	}

	// Step 6: Update the receiver's balance
	updateReceiver := fmt.Sprintf("UPDATE %s SET tokens_balance = tokens_balance + $1 WHERE %s = $2", receiver_table, receiver_column)
	_, err = tx.Exec(updateReceiver, transaction.Tokens, transaction.ToUserID)
	if err != nil {
		return nil, err
	}

	// Step 7: Add transfer date
	transaction.TransferDate = time.Now().Format("2006-01-02 15:04:05")

	// Step 8: Insert the transaction record
	_, err = tx.Exec("INSERT INTO transactions (from_user_id, to_user_id, tokens, transfer_date) VALUES ($1, $2, $3, $4)",
		transaction.FromUserID, transaction.ToUserID, transaction.Tokens, transaction.TransferDate)
	if err != nil {
		return nil, err
	}

	// Step 9: Commit transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
