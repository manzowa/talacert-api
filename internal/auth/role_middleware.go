package auth

import (
	"net/http"

	"talacert-api/internal/dto"

	"github.com/gin-gonic/gin"
)

func RequireRole(roles ...string) gin.HandlerFunc {

	return func(c *gin.Context) {

		userValue, exists := c.Get(ContextUser)

		if !exists {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{
					"success": false,
					"message": "role not found",
				},
			)
			return
		}

		user, ok := userValue.(dto.UserResponse)

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid user context",
			})
			return
		}

		for _, role := range roles {
			if user.Role == role {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "Access denied",
		})
	}
}
