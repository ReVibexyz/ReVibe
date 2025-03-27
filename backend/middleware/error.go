package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			// Handle different types of errors
			switch e := err.(type) {
			case *gin.Error:
				c.JSON(http.StatusBadRequest, ErrorResponse{
					Error:   "Bad Request",
					Message: e.Error(),
				})
			case gorm.ErrRecordNotFound:
				c.JSON(http.StatusNotFound, ErrorResponse{
					Error:   "Not Found",
					Message: "Resource not found",
				})
			default:
				c.JSON(http.StatusInternalServerError, ErrorResponse{
					Error:   "Internal Server Error",
					Message: "An unexpected error occurred",
				})
			}
		}
	}
}

// Custom error types
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type AuthenticationError struct {
	Message string
}

func (e *AuthenticationError) Error() string {
	return e.Message
}

type AuthorizationError struct {
	Message string
}

func (e *AuthorizationError) Error() string {
	return e.Message
}

// Helper functions for creating errors
func NewValidationError(message string) *ValidationError {
	return &ValidationError{Message: message}
}

func NewAuthenticationError(message string) *AuthenticationError {
	return &AuthenticationError{Message: message}
}

func NewAuthorizationError(message string) *AuthorizationError {
	return &AuthorizationError{Message: message}
} 