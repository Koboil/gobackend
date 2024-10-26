package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"fairground-backend/utils"
)

func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Step: 1 - Get the token from the request header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.JSONResponse(w, http.StatusUnauthorized, errors.New("authorization header not found"))
			return
		}

		// Step: 2 - Split the token "Bearer <token>"
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		// Step: 3 - Validate the token
		userId, email, role, err := utils.ValidateJWT(tokenString)
		if err != nil {
			utils.JSONResponse(w, http.StatusUnauthorized, err, "invalid token")
			return
		}

		// Step: 4 - Add the user information to the request header
		r.Header.Set("user_id", userId)
		r.Header.Set("email", email)
		r.Header.Set("role", role)
		next.ServeHTTP(w, r)
	})
}
