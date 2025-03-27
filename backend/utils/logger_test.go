package utils

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/revibe/backend/test"
)

func TestLogger(t *testing.T) {
	// Setup test environment
	env := test.SetupTestEnv(t)
	defer env.CleanupTestEnv()

	// Test logger initialization
	t.Run("InitLogger", func(t *testing.T) {
		err := InitLogger()
		assert.NoError(t, err)
		assert.NotNil(t, Logger)
	})

	// Test log levels
	t.Run("LogLevels", func(t *testing.T) {
		// Test error logging
		LogError(nil, map[string]interface{}{
			"test": "error",
		})

		// Test info logging
		LogInfo("test info", map[string]interface{}{
			"test": "info",
		})

		// Test debug logging
		LogDebug("test debug", map[string]interface{}{
			"test": "debug",
		})

		// Test warning logging
		LogWarning("test warning", map[string]interface{}{
			"test": "warning",
		})

		// Verify log file exists and contains entries
		logFile := env.GetTestFilePath("logs/app.log")
		content, err := os.ReadFile(logFile)
		assert.NoError(t, err)
		assert.Contains(t, string(content), "test error")
		assert.Contains(t, string(content), "test info")
		assert.Contains(t, string(content), "test debug")
		assert.Contains(t, string(content), "test warning")
	})

	// Test log rotation
	t.Run("RotateLogs", func(t *testing.T) {
		// Create old log file
		oldDate := time.Now().AddDate(0, 0, -8)
		oldLogFile := env.CreateTestFile(t, 
			filepath.Join("logs", oldDate.Format("2006-01-02")+".log"),
			"old log content",
		)

		// Create recent log file
		recentDate := time.Now().AddDate(0, 0, -1)
		recentLogFile := env.CreateTestFile(t,
			filepath.Join("logs", recentDate.Format("2006-01-02")+".log"),
			"recent log content",
		)

		// Rotate logs
		err := RotateLogs()
		assert.NoError(t, err)

		// Verify old log is deleted
		_, err = os.Stat(oldLogFile)
		assert.True(t, os.IsNotExist(err))

		// Verify recent log exists
		_, err = os.Stat(recentLogFile)
		assert.NoError(t, err)
	})

	// Test log with context
	t.Run("LogWithContext", func(t *testing.T) {
		fields := map[string]interface{}{
			"test": "context",
		}
		logger := LogWithContext(fields)
		assert.NotNil(t, logger)

		// Log with context
		logger.Info("test context log")

		// Verify log file contains context
		logFile := env.GetTestFilePath("logs/app.log")
		content, err := os.ReadFile(logFile)
		assert.NoError(t, err)
		assert.Contains(t, string(content), "test context log")
		assert.Contains(t, string(content), "test\":\"context")
	})
} 