package middleware

import (
	"net/http"
	"strconv"
	"talacert-api/internal/auth"

	"github.com/gin-gonic/gin"
)

func RequireOwnership() gin.HandlerFunc {

	return func(c *gin.Context) {

		userID := auth.GetUserID(c)

		paramID := c.Param("id")

		resourceID, err := strconv.Atoi(
			paramID,
		)

		if err != nil {

			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": "invalid id",
				},
			)
			return
		}

		role := auth.GetRole(c)

		if role == "ADMIN" {
			c.Next()
			return
		}

		if uint(resourceID) != userID {

			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{
					"message": "access denied",
				},
			)
			return
		}

		c.Next()
	}
}
