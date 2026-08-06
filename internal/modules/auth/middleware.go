package auth

import (
	"crypto/rsa"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/response"
)

// AuthMiddleware protects routes by validating Bearer RS256 JWT tokens in the Authorization header.
func AuthMiddleware(publicKey *rsa.PublicKey) gin.HandlerFunc {
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

		claims, err := ValidateAccessToken(parts[1], publicKey)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid or expired access token")
			c.Abort()
			return
		}

		// Store user metadata in Gin context for handlers and future RBAC middleware
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Set("claims", claims)

		c.Next()
	}
}
