package middleware

import (
	"go-xray-config/pkg/api"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			api.Unauthorized(c)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			api.Unauthorized(c)
			c.Abort()
			return
		}

		tokenString := parts[1]

		if tokenString != token {
			api.Unauthorized(c)
			c.Abort()
			return
		}

		c.Next()
	}
}
