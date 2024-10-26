package models

import (
	"errors"
	"fairground-backend/db"
)

type User struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      Role   `json:"role"` // 'Student', 'Parent', 'StandHolder', 'Organizer', 'FairgroundManager'
}

func CreateUser(user *User) error {
	// Step 1: Start a transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // Ensure rollback in case of failure

	// Step 2: Check if the user already exists
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", user.Email).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("user already exists")
	}

	// Step 3: Insert into the users table
	_, err = tx.Exec("INSERT INTO users (first_name, last_name, email, password, role) VALUES ($1, $2, $3, $4, $5)",
		user.FirstName, user.LastName, user.Email, user.Password, user.Role)
	if err != nil {
		return err
	}

	// Step 4: Retrieve user ID by email
	var userID int
	err = tx.QueryRow("SELECT user_id FROM users WHERE email = $1", user.Email).Scan(&userID)
	if err != nil {
		return err
	}

	// Step 5: Insert into the relevant role table using the retrieved userID
	switch user.Role {
	case RoleParent:
		_, err = tx.Exec("INSERT INTO parents (parent_id) VALUES ($1)", userID)
	case RoleOrganizer:
		_, err = tx.Exec("INSERT INTO organizers (organizer_id) VALUES ($1)", userID)
	case RoleStudent:
		_, err = tx.Exec("INSERT INTO students (student_id) VALUES ($1)", userID)
	case RoleStandHolder:
		_, err = tx.Exec("INSERT INTO standholders (stand_holder_id) VALUES ($1)", userID)
	default:
		return errors.New("invalid role")
	}

	if err != nil {
		return err
	}

	// Step 6: Commit the transaction
	return tx.Commit()
}

func GetUserByEmail(email string) (*User, error) {

	// Step 1: Retrieve user by email
	var user User
	err := db.DB.QueryRow("SELECT user_id, first_name, last_name, email, password, role FROM users WHERE email = $1", email).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	// Step 2: Return the user
	return &user, nil
}

func GetAllUsers() ([]User, error) {
	// Step 1: Retrieve all users
	rows, err := db.DB.Query("SELECT user_id, first_name, last_name, email, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Step 2: Iterate over the rows and create a slice of users
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Step 3: Return the slice of users
	return users, nil
}

func GetUserByID(id string) (*User, error) {
	// Step 1: Retrieve user by ID
	var user User
	err := db.DB.QueryRow("SELECT user_id, first_name, last_name, email, role FROM users WHERE user_id = $1", id).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Role)
	if err != nil {
		return nil, err
	}

	// Step 2: Return the user
	return &user, nil
}
