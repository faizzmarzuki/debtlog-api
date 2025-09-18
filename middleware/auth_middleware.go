package middleware

import (
	"net/http"

	"github.com/faizzmarzuki/debtlog-api/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates the JWT and injects user_id into context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization") // expect: Bearer <token>
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}
		tokenStr := utils.ExtractTokenFromHeader(auth)
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		// attach user id to context for handlers
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
