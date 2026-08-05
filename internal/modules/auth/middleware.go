package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/response"
)

// AuthMiddleware protects routes by validating Bearer JWT tokens in the Authorization header.
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "Authorization header missing")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "Invalid authorization header format")
			c.Abort()
			return
		}

		claims, err := ValidateToken(parts[1], jwtSecret)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Next()
	}
}
