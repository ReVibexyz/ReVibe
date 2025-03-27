package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/yourusername/revibe/backend/config"
)

// TestEnv represents the test environment
type TestEnv struct {
	TempDir string
}

// SetupTestEnv creates a new test environment
func SetupTestEnv(t *testing.T) *TestEnv {
	// Create temporary test directory
	tempDir, err := os.MkdirTemp("", "revibe_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}

	// Create subdirectories
	dirs := []string{
		"uploads",
		"logs",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(tempDir, dir), 0755); err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}

	// Update config for testing
	config.AppConfig.UploadDir = filepath.Join(tempDir, "uploads")
	config.AppConfig.LogDir = filepath.Join(tempDir, "logs")
	config.AppConfig.BaseURL = "http://localhost:8080"
	config.AppConfig.LogLevel = "debug"

	return &TestEnv{
		TempDir: tempDir,
	}
}

// CleanupTestEnv cleans up the test environment
func (env *TestEnv) CleanupTestEnv() {
	os.RemoveAll(env.TempDir)
}

// CreateTestFile creates a test file with the given content
func (env *TestEnv) CreateTestFile(t *testing.T, path, content string) string {
	fullPath := filepath.Join(env.TempDir, path)
	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		t.Fatalf("Failed to create directory for test file: %v", err)
	}

	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	return fullPath
}

// GetTestFilePath returns the full path for a test file
func (env *TestEnv) GetTestFilePath(path string) string {
	return filepath.Join(env.TempDir, path)
} 