package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var CONFIG *Config

// Config holds the configuration for the application
type Config struct {
	DB_HOST      string
	DB_PORT      string
	DB_USER      string
	DB_PASSWORD  string
	DB_NAME      string
	JWT_SECRET   string
	BACKEND_PORT string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	CONFIG = &Config{
		DB_HOST:      os.Getenv("DB_HOST"),
		DB_PORT:      os.Getenv("DB_PORT"),
		DB_USER:      os.Getenv("DB_USER"),
		DB_PASSWORD:  os.Getenv("DB_PASSWORD"),
		DB_NAME:      os.Getenv("DB_NAME"),
		JWT_SECRET:   os.Getenv("JWT_SECRET"),
		BACKEND_PORT: os.Getenv("BACKEND_PORT"),
	}
}
