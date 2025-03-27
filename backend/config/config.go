package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	Port string

	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// JWT
	JWTSecret string

	// Web3
	InfuraID        string
	ContractAddress string
	ChainID         string

	// Storage
	UploadDir string
}

var AppConfig Config

func LoadConfig() error {
	// Load .env file if it exists
	_ = godotenv.Load()

	AppConfig = Config{
		// Server
		Port: getEnvOrDefault("PORT", "8080"),

		// Database
		DBHost:     getEnvOrDefault("DB_HOST", "localhost"),
		DBPort:     getEnvOrDefault("DB_PORT", "5432"),
		DBUser:     getEnvOrDefault("DB_USER", "postgres"),
		DBPassword: getEnvOrDefault("DB_PASSWORD", "postgres"),
		DBName:     getEnvOrDefault("DB_NAME", "revibe"),

		// JWT
		JWTSecret: getEnvOrDefault("JWT_SECRET", "your-secret-key"),

		// Web3
		InfuraID:        getEnvOrDefault("INFURA_ID", ""),
		ContractAddress: getEnvOrDefault("CONTRACT_ADDRESS", ""),
		ChainID:         getEnvOrDefault("CHAIN_ID", "1"),

		// Storage
		UploadDir: getEnvOrDefault("UPLOAD_DIR", "uploads"),
	}

	return nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetDSN returns the database connection string
func (c *Config) GetDSN() string {
	return "host=" + c.DBHost +
		" port=" + c.DBPort +
		" user=" + c.DBUser +
		" password=" + c.DBPassword +
		" dbname=" + c.DBName +
		" sslmode=disable"
} 