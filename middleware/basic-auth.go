package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Hardcoded credentials
const (
	adminUsername = "admin"
	adminPassword = "admin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header required"})
			c.Abort()
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Basic" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization header format"})
			c.Abort()
			return
		}

		user, pass, ok := c.Request.BasicAuth()
		if !ok || !validateCredentials(user, pass) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// validateCredentials checks the credentials against hardcoded values
func validateCredentials(username, password string) bool {
	return username == adminUsername && password == adminPassword
}
