package utils

import (
	"errors"
	"fairground-backend/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// GenerateJWT generates a JWT token
func GenerateJWT(userID string, email string, role string) (string, error) {

	secretKey := []byte(config.CONFIG.JWT_SECRET)

	// Create JWT claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	return token.SignedString(secretKey)
}

// ValidateJWT validates the JWT token and returns email and role
func ValidateJWT(tokenString string) (string, string, string, error) {

	secretKey := []byte(config.CONFIG.JWT_SECRET)

	// Step: 1 - Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return secretKey, nil
	})

	// Step: 2 - Check if token is valid
	if err != nil || !token.Valid {
		return "", "", "", errors.New("invalid token")
	}

	// Step: 3 - Parse claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", "", errors.New("invalid claims")
	}

	// Step: 4 - Check if email is valid
	email, ok := claims["email"].(string)
	if !ok {
		return "", "", "", errors.New("invalid email")
	}

	// Step: 5 - Check if role is valid
	role, ok := claims["role"].(string)
	if !ok {
		return "", "", "", errors.New("invalid role")
	}

	// Step: 6 - Check if user_id is valid
	userId, ok := claims["user_id"].(string)
	if !ok {
		return "", "", "", errors.New("invalid user_id")
	}

	// Step: 7 - Return email and role
	return userId, email, role, nil
}
