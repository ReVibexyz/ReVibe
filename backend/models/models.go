package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID            string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	WalletAddress string    `gorm:"uniqueIndex;not null" json:"walletAddress"`
	Name          string    `gorm:"size:255" json:"name"`
	Avatar        string    `gorm:"size:255" json:"avatar"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

// Product represents a product listing
type Product struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Category    string    `gorm:"size:50;not null" json:"category"`
	Condition   string    `gorm:"size:50;not null" json:"condition"`
	SellerID    string    `gorm:"type:uuid;not null" json:"sellerId"`
	Seller      User      `gorm:"foreignKey:SellerID" json:"seller"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// ProductImage represents an image associated with a product
type ProductImage struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	ProductID string    `gorm:"type:uuid;not null" json:"productId"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"-"`
	URL       string    `gorm:"size:255;not null" json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Order represents a product purchase order
type Order struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	ProductID   string    `gorm:"type:uuid;not null" json:"productId"`
	Product     Product   `gorm:"foreignKey:ProductID" json:"product"`
	BuyerID     string    `gorm:"type:uuid;not null" json:"buyerId"`
	Buyer       User      `gorm:"foreignKey:BuyerID" json:"buyer"`
	Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Status      string    `gorm:"size:50;not null;default:'pending'" json:"status"`
	TxHash      string    `gorm:"size:66" json:"txHash"`
	CompletedAt time.Time `json:"completedAt"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// Authentication represents a product authentication record
type Authentication struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	ProductID string    `gorm:"type:uuid;not null" json:"productId"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	Result    bool      `gorm:"not null" json:"result"`
	Score     float64   `gorm:"type:decimal(5,2)" json:"score"`
	Details   string    `gorm:"type:text" json:"details"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// AutoMigrate performs database migrations
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Product{},
		&ProductImage{},
		&Order{},
		&Authentication{},
	)
} 