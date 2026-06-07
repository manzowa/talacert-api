package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"talacert-api/internal/logger"
)

func AccessLog() gin.HandlerFunc {

	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		logger.AccessLogger.Info(
			"HTTP Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"ip", c.ClientIP(),
			"duration", time.Since(start),
		)
	}
}
