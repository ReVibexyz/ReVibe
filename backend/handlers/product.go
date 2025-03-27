package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Price       float64   `json:"price" binding:"required"`
	Images      []string  `json:"images" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Condition   string    `json:"condition" binding:"required"`
	SellerID    string    `json:"sellerId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func HandleGetProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []Product
		query := db.Model(&Product{})

		// Apply filters
		if category := c.Query("category"); category != "" {
			query = query.Where("category = ?", category)
		}
		if minPrice := c.Query("minPrice"); minPrice != "" {
			query = query.Where("price >= ?", minPrice)
		}
		if maxPrice := c.Query("maxPrice"); maxPrice != "" {
			query = query.Where("price <= ?", maxPrice)
		}
		if search := c.Query("search"); search != "" {
			query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
		}

		if err := query.Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
			return
		}

		c.JSON(http.StatusOK, products)
	}
}

func HandleGetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product Product

		if err := db.First(&product, "id = ?", id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

func HandleCreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Get user ID from context (set by auth middleware)
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		product.ID = uuid.New().String()
		product.SellerID = userID.(string)
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		if err := db.Create(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
			return
		}

		c.JSON(http.StatusCreated, product)
	}
}

func HandleUpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product Product

		if err := db.First(&product, "id = ?", id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
			return
		}

		// Check if user is the seller
		userID, exists := c.Get("userID")
		if !exists || userID.(string) != product.SellerID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to update this product"})
			return
		}

		var updateData Product
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update fields
		product.Name = updateData.Name
		product.Description = updateData.Description
		product.Price = updateData.Price
		product.Images = updateData.Images
		product.Category = updateData.Category
		product.Condition = updateData.Condition
		product.UpdatedAt = time.Now()

		if err := db.Save(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

func HandleDeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product Product

		if err := db.First(&product, "id = ?", id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
			return
		}

		// Check if user is the seller
		userID, exists := c.Get("userID")
		if !exists || userID.(string) != product.SellerID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to delete this product"})
			return
		}

		if err := db.Delete(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
	}
}

func HandleAuthenticateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product Product

		if err := db.First(&product, "id = ?", id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
			return
		}

		// TODO: Implement product authentication logic
		// This should interact with the smart contract to verify the product's authenticity

		c.JSON(http.StatusOK, gin.H{"message": "Product authentication initiated"})
	}
} 