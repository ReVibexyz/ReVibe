package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yourusername/revibe/backend/test"
)

func TestLoggerMiddleware(t *testing.T) {
	// Setup test environment
	env := test.SetupTestEnv(t)
	defer env.CleanupTestEnv()

	// Create Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(Logger())

	// Add test route
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	// Create test request
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("User-Agent", "test-agent")
	req.RemoteAddr = "127.0.0.1"

	// Create response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verify response
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify log file contains request details
	logFile := env.GetTestFilePath("logs/app.log")
	content, err := os.ReadFile(logFile)
	assert.NoError(t, err)
	assert.Contains(t, string(content), "HTTP Request")
	assert.Contains(t, string(content), "method\":\"GET")
	assert.Contains(t, string(content), "path\":\"/test")
	assert.Contains(t, string(content), "status\":200")
	assert.Contains(t, string(content), "client_ip\":\"127.0.0.1")
	assert.Contains(t, string(content), "user_agent\":\"test-agent")
} 