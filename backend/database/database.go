package database

import (
	"fmt"
	"log"
	"time"

	"github.com/yourusername/revibe/backend/config"
	"github.com/yourusername/revibe/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection and performs migrations
func InitDB() error {
	dsn := config.AppConfig.GetDSN()

	// Configure GORM logger
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Get underlying SQL DB
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Perform migrations
	if err := models.AutoMigrate(db); err != nil {
		return fmt.Errorf("failed to perform database migrations: %v", err)
	}

	DB = db
	log.Println("Database connection established successfully")
	return nil
}

// CloseDB closes the database connection
func CloseDB() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return fmt.Errorf("failed to get database instance: %v", err)
		}
		return sqlDB.Close()
	}
	return nil
}

// Transaction executes a function within a database transaction
func Transaction(fc func(tx *gorm.DB) error) error {
	return DB.Transaction(fc)
} 