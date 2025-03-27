package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserProfile struct {
	ID            string    `json:"id"`
	WalletAddress string    `json:"walletAddress"`
	Name          string    `json:"name"`
	Avatar        string    `json:"avatar"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type UpdateUserRequest struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func HandleGetUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		walletAddress := c.Param("walletAddress")
		var user UserProfile

		if err := db.First(&user, "wallet_address = ?", walletAddress).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func HandleUpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		walletAddress := c.Param("walletAddress")
		var user UserProfile

		if err := db.First(&user, "wallet_address = ?", walletAddress).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
			return
		}

		// Check if user is updating their own profile
		userID, exists := c.Get("userID")
		if !exists || userID.(string) != user.ID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to update this profile"})
			return
		}

		var updateData UpdateUserRequest
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update fields
		user.Name = updateData.Name
		user.Avatar = updateData.Avatar
		user.UpdatedAt = time.Now()

		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func HandleGetUserProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		walletAddress := c.Param("walletAddress")
		var user UserProfile

		if err := db.First(&user, "wallet_address = ?", walletAddress).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
			return
		}

		var products []Product
		if err := db.Where("seller_id = ?", user.ID).Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user products"})
			return
		}

		c.JSON(http.StatusOK, products)
	}
}

func HandleGetUserOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		walletAddress := c.Param("walletAddress")
		var user UserProfile

		if err := db.First(&user, "wallet_address = ?", walletAddress).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
			return
		}

		// TODO: Implement order fetching logic
		// This should fetch orders from the smart contract and combine with product data

		c.JSON(http.StatusOK, gin.H{"message": "Orders endpoint not implemented yet"})
	}
} 