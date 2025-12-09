package utils

import "AdvAuthGo/internal/models"

func ExtractRoleNames(roles []models.Role) []string {
	names := make([]string, 0, len(roles))
	for _, r := range roles {
		names = append(names, r.Name)
	}
	return names
}
