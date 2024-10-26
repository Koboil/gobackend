package middlewares

import (
	"net/http"
	_ "strings"
)

// RoleAuthMiddleware checks if the user has one of the allowed roles to access the API
func RoleAuthMiddleware(allowedRoles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Step: 1 - Get the role from the request header
		role := r.Header.Get("role")
		if role == "" {
			http.Error(w, "Role not found in request", http.StatusForbidden)
			return
		}

		// Step: 2 - Check if the role is allowed
		if !isRoleAllowed(role, allowedRoles) {
			http.Error(w, "Access denied: insufficient role", http.StatusForbidden)
			return
		}

		// Step: 3 - Call the next handler
		next.ServeHTTP(w, r)
	})
}

// isRoleAllowed checks if the user's role is in the allowed roles
func isRoleAllowed(userRole string, allowedRoles []string) bool {
	for _, allowedRole := range allowedRoles {
		if userRole == allowedRole {
			return true
		}
	}
	return false
}
