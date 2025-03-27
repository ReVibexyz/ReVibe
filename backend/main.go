package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yourusername/revibe/backend/config"
	"github.com/yourusername/revibe/backend/database"
	"github.com/yourusername/revibe/backend/handlers"
	"github.com/yourusername/revibe/backend/middleware"
	"github.com/yourusername/revibe/backend/services"
	"github.com/yourusername/revibe/backend/utils"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		utils.LogWarning("Error loading .env file", map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Initialize logger
	if err := utils.InitLogger(); err != nil {
		utils.LogFatal(err, nil)
	}

	// Load configuration
	if err := config.LoadConfig(); err != nil {
		utils.LogFatal(err, nil)
	}

	// Initialize database
	if err := database.InitDB(); err != nil {
		utils.LogFatal(err, nil)
	}
	defer database.CloseDB()

	// Initialize Web3 service
	web3Service, err := services.NewWeb3Service()
	if err != nil {
		utils.LogFatal(err, nil)
	}
	defer web3Service.Close()

	// Initialize upload service
	uploadService, err := services.NewUploadService()
	if err != nil {
		utils.LogFatal(err, nil)
	}

	// Initialize metrics service
	metricsService, err := services.NewMetricsService()
	if err != nil {
		utils.LogFatal(err, nil)
	}

	// Create context for event listeners and metrics collector
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start event listeners
	if err := web3Service.StartEventListeners(ctx); err != nil {
		utils.LogError(err, map[string]interface{}{
			"component": "event_listeners",
		})
	}

	// Start metrics collector
	go metricsService.StartMetricsCollector(ctx)

	// Create Gin router
	router := gin.Default()

	// Add middleware
	router.Use(gin.Recovery())
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.Cors())
	router.Use(middleware.Logger())
	router.Use(middleware.Metrics(metricsService))

	// Setup routes
	setupRoutes(router, web3Service, uploadService, metricsService)

	// Start server
	server := &http.Server{
		Addr:    ":" + config.AppConfig.Port,
		Handler: router,
	}

	// Start log rotation
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := utils.RotateLogs(); err != nil {
					utils.LogError(err, map[string]interface{}{
						"component": "log_rotation",
					})
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	// Handle graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		utils.LogInfo("Shutting down server...", nil)
		cancel() // Cancel context to stop event listeners and metrics collector
		if err := server.Close(); err != nil {
			utils.LogError(err, map[string]interface{}{
				"component": "server_shutdown",
			})
		}
	}()

	utils.LogInfo("Server starting", map[string]interface{}{
		"port": config.AppConfig.Port,
	})
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		utils.LogFatal(err, nil)
	}
}

func initDB() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=revibe port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate models
	if err := autoMigrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func setupRoutes(router *gin.Engine, web3Service *services.Web3Service, uploadService *services.UploadService, metricsService *services.MetricsService) {
	// Health check and metrics routes
	router.GET("/health", handlers.HandleHealthCheck())
	router.GET("/metrics", handlers.HandleMetrics())

	// Auth routes
	auth := router.Group("/auth")
	{
		auth.POST("/login", handlers.HandleLogin(database.DB, web3Service))
		auth.POST("/verify", handlers.HandleVerify(database.DB, web3Service))
	}

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Product routes
		products := protected.Group("/products")
		{
			products.GET("", handlers.HandleGetProducts(database.DB, web3Service))
			products.GET("/:id", handlers.HandleGetProduct(database.DB, web3Service))
			products.POST("", handlers.HandleCreateProduct(database.DB, web3Service))
			products.PUT("/:id", handlers.HandleUpdateProduct(database.DB, web3Service))
			products.DELETE("/:id", handlers.HandleDeleteProduct(database.DB, web3Service))
			products.POST("/:id/authenticate", handlers.HandleAuthenticateProduct(database.DB, web3Service))
		}

		// User routes
		users := protected.Group("/users")
		{
			users.GET("/:walletAddress", handlers.HandleGetUser(database.DB, web3Service))
			users.PUT("/:walletAddress", handlers.HandleUpdateUser(database.DB, web3Service))
			users.GET("/:walletAddress/products", handlers.HandleGetUserProducts(database.DB, web3Service))
			users.GET("/:walletAddress/orders", handlers.HandleGetUserOrders(database.DB, web3Service))
		}

		// Upload routes
		uploads := protected.Group("/uploads")
		{
			uploads.POST("", handlers.HandleUpload(uploadService))
			uploads.DELETE("", handlers.HandleDeleteFile(uploadService))
			uploads.POST("/cleanup", handlers.HandleCleanupUnusedFiles(uploadService, database.DB))
		}
	}

	// Public file routes
	router.GET("/uploads/*path", handlers.HandleGetFile(uploadService))
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func autoMigrate(db *gorm.DB) error {
	// Add models to migrate
	return nil
} 