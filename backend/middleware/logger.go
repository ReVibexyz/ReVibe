package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/revibe/backend/utils"
)

// Logger middleware logs HTTP requests
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log request details
		utils.LogInfo("HTTP Request", map[string]interface{}{
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"status":     c.Writer.Status(),
			"latency":    time.Since(start).String(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		})
	}
} 