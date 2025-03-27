# ReVibe Backend Development Guide

## Project Structure
```
backend/
├── main.go           # Application entry point
├── models/           # Database models
├── routes/           # API routes
├── handlers/         # Request handlers
├── middleware/       # Custom middleware
├── services/         # Business logic
├── utils/           # Utility functions
└── config/          # Configuration
```

## Technology Stack

### Core
- Go 1.19+
- Gin web framework
- GORM ORM
- PostgreSQL
- JWT authentication

### Additional
- Redis for caching
- AWS S3 for storage
- Zap for logging
- Swagger for API docs

## Database Models

### User Model
```go
type User struct {
    gorm.Model
    WalletAddress string    `gorm:"unique;not null"`
    Name          string    `gorm:"size:255"`
    Avatar        string    `gorm:"size:255"`
    Products      []Product `gorm:"foreignKey:SellerID"`
    Purchases     []Order   `gorm:"foreignKey:BuyerID"`
}
```

### Product Model
```go
type Product struct {
    gorm.Model
    Name         string    `gorm:"size:255;not null"`
    Description  string    `gorm:"type:text"`
    Price        string    `gorm:"not null"`
    Images       []Image   `gorm:"foreignKey:ProductID"`
    Category     string    `gorm:"size:100;not null"`
    Condition    string    `gorm:"size:50;not null"`
    SellerID     uint      `gorm:"not null"`
    Seller       User      `gorm:"foreignKey:SellerID"`
    Authenticity string    `gorm:"size:255"`
    Orders       []Order   `gorm:"foreignKey:ProductID"`
}
```

## API Routes

### Route Setup
```go
func setupRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()
    
    // Middleware
    router.Use(corsMiddleware())
    router.Use(authMiddleware())
    
    // Routes
    api := router.Group("/api")
    {
        setupAuthRoutes(api, db)
        setupProductRoutes(api, db)
        setupUserRoutes(api, db)
    }
    
    return router
}
```

### Route Handlers
```go
func handleGetProducts(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var products []Product
        if err := db.Find(&products).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, products)
    }
}
```

## Middleware

### Authentication
```go
func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            c.Abort()
            return
        }
        
        // Validate token
        claims, err := validateToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }
        
        c.Set("user", claims)
        c.Next()
    }
}
```

### CORS
```go
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
```

## Services

### Product Service
```go
type ProductService struct {
    db *gorm.DB
}

func (s *ProductService) CreateProduct(product *Product) error {
    return s.db.Create(product).Error
}

func (s *ProductService) GetProduct(id uint) (*Product, error) {
    var product Product
    if err := s.db.First(&product, id).Error; err != nil {
        return nil, err
    }
    return &product, nil
}
```

### Authentication Service
```go
type AuthService struct {
    db *gorm.DB
}

func (s *AuthService) Login(walletAddress, signature string) (*User, string, error) {
    // Verify signature
    if err := verifySignature(walletAddress, signature); err != nil {
        return nil, "", err
    }
    
    // Get or create user
    user, err := s.getOrCreateUser(walletAddress)
    if err != nil {
        return nil, "", err
    }
    
    // Generate token
    token, err := generateToken(user)
    if err != nil {
        return nil, "", err
    }
    
    return user, token, nil
}
```

## Database Operations

### Migrations
```go
func autoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &User{},
        &Product{},
        &Image{},
        &Order{},
        &Authentication{},
    )
}
```

### Queries
```go
// Find with relations
db.Preload("Seller").Preload("Images").Find(&products)

// Complex queries
db.Where("price >= ?", minPrice).
   Where("category = ?", category).
   Order("created_at desc").
   Find(&products)
```

## Error Handling

### Custom Errors
```go
type AppError struct {
    Code    int
    Message string
    Err     error
}

func (e *AppError) Error() string {
    return e.Message
}
```

### Error Handler
```go
func errorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        
        if len(c.Errors) > 0 {
            err := c.Errors.Last().Err
            if appErr, ok := err.(*AppError); ok {
                c.JSON(appErr.Code, gin.H{"error": appErr.Message})
                return
            }
            c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
        }
    }
}
```

## Configuration

### Environment Variables
```go
type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    JWTSecret  string
    RedisURL   string
    S3Bucket   string
}

func LoadConfig() (*Config, error) {
    return &Config{
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
        JWTSecret:  os.Getenv("JWT_SECRET"),
        RedisURL:   os.Getenv("REDIS_URL"),
        S3Bucket:   os.Getenv("S3_BUCKET"),
    }, nil
}
```

## Logging

### Logger Setup
```go
func setupLogger() *zap.Logger {
    config := zap.NewProductionConfig()
    config.OutputPaths = []string{"stdout", "logs/app.log"}
    config.ErrorOutputPaths = []string{"stderr", "logs/error.log"}
    
    logger, err := config.Build()
    if err != nil {
        panic(err)
    }
    
    return logger
}
```

### Logging Middleware
```go
func loggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        
        c.Next()
        
        logger.Info("request completed",
            zap.String("method", c.Request.Method),
            zap.String("path", c.Request.URL.Path),
            zap.Int("status", c.Writer.Status()),
            zap.Duration("duration", time.Since(start)),
        )
    }
}
```

## Testing

### Unit Tests
```go
func TestProductService_CreateProduct(t *testing.T) {
    db := setupTestDB()
    service := NewProductService(db)
    
    product := &Product{
        Name: "Test Product",
        Price: "0.5",
    }
    
    err := service.CreateProduct(product)
    assert.NoError(t, err)
    assert.NotZero(t, product.ID)
}
```

### Integration Tests
```go
func TestProductAPI(t *testing.T) {
    router := setupTestRouter()
    
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/api/products", nil)
    router.ServeHTTP(w, req)
    
    assert.Equal(t, 200, w.Code)
    
    var response map[string]interface{}
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.NoError(t, err)
}
```

## Deployment

### Docker Setup
```dockerfile
FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
```

### Kubernetes Deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: revibe-backend
spec:
  replicas: 3
  selector:
    matchLabels:
      app: revibe-backend
  template:
    metadata:
      labels:
        app: revibe-backend
    spec:
      containers:
      - name: revibe-backend
        image: revibe-backend:latest
        ports:
        - containerPort: 8080
        env:
        - name: DB_HOST
          valueFrom:
            secretKeyRef:
              name: revibe-secrets
              key: db-host
```

## Monitoring

### Health Check
```go
func healthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "healthy",
        "time":   time.Now(),
    })
}
```

### Metrics
```go
func setupMetrics() {
    prometheus.Register(httpRequestsTotal)
    prometheus.Register(httpRequestDuration)
}
```

## Security

### Input Validation
```go
func validateProduct(product *Product) error {
    if product.Name == "" {
        return errors.New("name is required")
    }
    if product.Price == "" {
        return errors.New("price is required")
    }
    return nil
}
```

### Rate Limiting
```go
func rateLimitMiddleware() gin.HandlerFunc {
    limiter := rate.NewLimiter(rate.Every(time.Second), 10)
    return func(c *gin.Context) {
        if !limiter.Allow() {
            c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
            c.Abort()
            return
        }
        c.Next()
    }
}
``` 