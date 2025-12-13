package middleware

import (
	"AdvAuthGo/internal/models"
	"AdvAuthGo/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRoles(authService services.AuthService, allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("refresh_token")
		if err != nil || refreshToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		user, err := authService.GetUserByToken(refreshToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		if !userHasRole(user, allowedRoles) {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func userHasRole(user *models.User, allowed []string) bool {
	roleMap := make(map[string]struct{}, len(user.Roles))
	for _, role := range user.Roles {
		roleMap[role.Name] = struct{}{}
	}

	for _, required := range allowed {
		if _, ok := roleMap[required]; ok {
			return true
		}
	}
	return false
}
