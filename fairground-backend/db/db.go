// db/db.go
package db

import (
	"database/sql"
	"fairground-backend/config"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

// ConnectDB initializes the database connection pool.
// It now takes a hostname parameter for local development.
func ConnectDB() {

	// Step: 1 - Load environment variables
	var err error

	// Step: 2 - Connect to the database
	DB, err = sql.Open("postgres", "host="+config.CONFIG.DB_HOST+" user="+config.CONFIG.DB_USER+" password="+config.CONFIG.DB_PASSWORD+" dbname="+config.CONFIG.DB_NAME+" sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Step: 3 - Ping the database
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	log.Println("Connected to the database")
}

//// db/db.go
//package db
//
//import (
//	"database/sql"
//	"fairground-backend/config"
//	"log"
//
//	_ "github.com/lib/pq" // PostgreSQL driver
//)
//
//var DB *sql.DB
//
//// ConnectDB initializes the database connection pool.
//func ConnectDB() {
//
//	// Step: 1 - Load environment variables
//	var err error
//
//	// Step: 2 - Connect to the database
//	DB, err = sql.Open("postgres", "user="+config.CONFIG.DB_USER+" password="+config.CONFIG.DB_PASSWORD+" dbname="+config.CONFIG.DB_NAME+" sslmode=disable")
//	if err != nil {
//		log.Fatal("Failed to connect to the database:", err)
//	}
//
//	// Step: 3 - Ping the database
//	err = DB.Ping()
//	if err != nil {
//		log.Fatal("Failed to ping the database:", err)
//	}
//
//	log.Println("Connected to the database")
//}
