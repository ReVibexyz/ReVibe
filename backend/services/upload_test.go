package services

import (
	"bytes"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/revibe/backend/config"
	"github.com/yourusername/revibe/backend/test"
)

func TestUploadService(t *testing.T) {
	// Setup test environment
	env := test.SetupTestEnv(t)
	defer env.CleanupTestEnv()

	// Create upload service
	service, err := NewUploadService()
	assert.NoError(t, err)
	assert.NotNil(t, service)

	// Test file upload
	t.Run("UploadFile", func(t *testing.T) {
		// Create test file
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", "test.jpg")
		assert.NoError(t, err)
		part.Write([]byte("test image content"))
		writer.Close()

		// Create file header
		fileHeader := &multipart.FileHeader{
			Filename: "test.jpg",
			Size:     int64(len(body.Bytes())),
		}

		// Upload file
		url, err := service.UploadFile(fileHeader, "products")
		assert.NoError(t, err)
		assert.NotEmpty(t, url)

		// Verify file exists
		filePath := env.GetTestFilePath(filepath.Join("uploads", "products", filepath.Base(url)))
		_, err = os.Stat(filePath)
		assert.NoError(t, err)
	})

	// Test file deletion
	t.Run("DeleteFile", func(t *testing.T) {
		// Create test file
		testFile := env.CreateTestFile(t, "uploads/products/test_delete.jpg", "test content")

		// Delete file
		err := service.DeleteFile("/uploads/products/test_delete.jpg")
		assert.NoError(t, err)

		// Verify file is deleted
		_, err = os.Stat(testFile)
		assert.True(t, os.IsNotExist(err))
	})

	// Test file type validation
	t.Run("ValidateFileType", func(t *testing.T) {
		// Test valid file type
		validFile := &multipart.FileHeader{
			Filename: "test.jpg",
			Size:     1024 * 1024, // 1MB
		}
		err := service.validateFileType(validFile)
		assert.NoError(t, err)

		// Test invalid file type
		invalidFile := &multipart.FileHeader{
			Filename: "test.txt",
			Size:     1024,
		}
		err = service.validateFileType(invalidFile)
		assert.Error(t, err)

		// Test file size limit
		largeFile := &multipart.FileHeader{
			Filename: "test.jpg",
			Size:     11 * 1024 * 1024, // 11MB
		}
		err = service.validateFileType(largeFile)
		assert.Error(t, err)
	})

	// Test file URL generation
	t.Run("GetFileURL", func(t *testing.T) {
		url := service.GetFileURL("products/test.jpg")
		expectedURL := config.AppConfig.BaseURL + "/uploads/products/test.jpg"
		assert.Equal(t, expectedURL, url)
	})

	// Test cleanup unused files
	t.Run("CleanupUnusedFiles", func(t *testing.T) {
		// Create test files
		usedFile := env.CreateTestFile(t, "uploads/products/used.jpg", "used content")
		unusedFile := env.CreateTestFile(t, "uploads/products/unused.jpg", "unused content")

		// Create product image record
		productImage := models.ProductImage{
			URL: "products/used.jpg",
		}

		// Cleanup unused files
		err := service.CleanupUnusedFiles(database.DB)
		assert.NoError(t, err)

		// Verify used file exists
		_, err = os.Stat(usedFile)
		assert.NoError(t, err)

		// Verify unused file is deleted
		_, err = os.Stat(unusedFile)
		assert.True(t, os.IsNotExist(err))
	})
} 