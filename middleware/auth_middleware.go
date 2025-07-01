package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	// utils import
	"test-golang-api/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, utils.ApiResponse{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
				Data:    nil,
			})

			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.ApiResponse{
				Status:  http.StatusUnauthorized,
				Message: "Invalid token",
				Data:    nil,
			})

			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
