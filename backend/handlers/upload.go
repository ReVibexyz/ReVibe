package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/revibe/backend/services"
)

// HandleUpload handles file upload requests
func HandleUpload(uploadService *services.UploadService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get file from request
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
			return
		}

		// Get subdirectory from query parameter
		subDir := c.DefaultQuery("dir", "products")

		// Upload file
		url, err := uploadService.UploadFile(file, subDir)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return file URL
		c.JSON(http.StatusOK, gin.H{
			"url": uploadService.GetFileURL(url),
		})
	}
}

// HandleDeleteFile handles file deletion requests
func HandleDeleteFile(uploadService *services.UploadService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get file URL from request
		url := c.Query("url")
		if url == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file URL provided"})
			return
		}

		// Delete file
		if err := uploadService.DeleteFile(url); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
	}
}

// HandleCleanupUnusedFiles handles cleanup of unused files
func HandleCleanupUnusedFiles(uploadService *services.UploadService, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cleanup unused files
		if err := uploadService.CleanupUnusedFiles(db); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Cleanup completed successfully"})
	}
}

// HandleGetFile handles file retrieval requests
func HandleGetFile(uploadService *services.UploadService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get file path from URL
		path := c.Param("path")
		if path == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file path provided"})
			return
		}

		// Get file path
		filepath := filepath.Join(uploadService.UploadDir, path)

		// Check if file exists
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}

		// Serve file
		c.File(filepath)
	}
} 