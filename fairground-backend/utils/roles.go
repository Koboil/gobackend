package utils

import "fairground-backend/models"

// IsValidRole Function to check if a role is valid
func IsValidRole(role models.Role) bool {
	switch role {
	case models.RoleStudent, models.RoleParent, models.RoleStandHolder, models.RoleOrganizer, models.RoleFairgroundManagement:
		return true
	}
	return false
}
