package auth

import (
	"net/http"
	"strings"

	"talacert-api/internal/dto"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(
	jwtManager *JWTManager,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"success": false,
					"message": "authorization header required",
				},
			)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"success": false,
					"message": "invalid authorization format",
				},
			)
			return
		}

		tokenString := strings.TrimSpace(parts[1])

		if tokenString == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"success": false,
					"message": "token missing",
				},
			)
			return
		}

		claims, err := jwtManager.VerifyAccessToken(tokenString)

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"success": false,
					"message": "invalid token",
				},
			)
			return
		}

		user := dto.UserResponse{
			ID:       claims.UserID,
			Username: claims.Username,
			Email:    claims.Email,
			Role:     claims.Role,
		}

		c.Set(ContextUser, user)

		c.Next()
	}
}
