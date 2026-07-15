package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"talacert-api/internal/logger"
)

func ErrorLog() gin.HandlerFunc {

	return func(c *gin.Context) {

		defer func() {

			if err := recover(); err != nil {

				logger.ErrorLogger.Error(
					"Panic recovered",
					"error", err,
					"stack", string(debug.Stack()),
				)

				c.JSON(
					http.StatusInternalServerError,
					gin.H{
						"error": "Internal Server Error",
					},
				)

				c.Abort()
			}
		}()

		c.Next()
	}
}
