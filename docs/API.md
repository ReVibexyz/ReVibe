# ReVibe API Documentation

## Base URL
```
http://localhost:8080/api
```

## Authentication

### Login
```http
POST /auth/login
```

Request body:
```json
{
  "walletAddress": "0x...",
  "signature": "0x..."
}
```

Response:
```json
{
  "token": "jwt_token",
  "user": {
    "walletAddress": "0x...",
    "name": "John Doe",
    "avatar": "https://..."
  }
}
```

### Verify
```http
POST /auth/verify
```

Request body:
```json
{
  "token": "jwt_token"
}
```

Response:
```json
{
  "valid": true,
  "user": {
    "walletAddress": "0x...",
    "name": "John Doe",
    "avatar": "https://..."
  }
}
```

## Products

### Get Products
```http
GET /products
```

Query parameters:
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)
- `category` (optional): Filter by category
- `sort` (optional): Sort by field (price, date, etc.)
- `order` (optional): Sort order (asc, desc)

Response:
```json
{
  "products": [
    {
      "id": "1",
      "name": "Limited Edition Sneaker",
      "description": "Exclusive limited edition sneaker",
      "price": "0.5",
      "image": "https://...",
      "category": "Footwear",
      "seller": {
        "walletAddress": "0x...",
        "name": "John Doe"
      }
    }
  ],
  "total": 100,
  "page": 1,
  "limit": 10
}
```

### Get Product
```http
GET /products/:id
```

Response:
```json
{
  "id": "1",
  "name": "Limited Edition Sneaker",
  "description": "Exclusive limited edition sneaker",
  "price": "0.5",
  "images": [
    "https://...",
    "https://..."
  ],
  "category": "Footwear",
  "condition": "New",
  "authenticity": "Verified",
  "seller": {
    "walletAddress": "0x...",
    "name": "John Doe"
  }
}
```

### Create Product
```http
POST /products
```

Request body:
```json
{
  "name": "Limited Edition Sneaker",
  "description": "Exclusive limited edition sneaker",
  "price": "0.5",
  "images": [
    "https://...",
    "https://..."
  ],
  "category": "Footwear",
  "condition": "New"
}
```

Response:
```json
{
  "id": "1",
  "name": "Limited Edition Sneaker",
  "description": "Exclusive limited edition sneaker",
  "price": "0.5",
  "images": [
    "https://...",
    "https://..."
  ],
  "category": "Footwear",
  "condition": "New",
  "seller": {
    "walletAddress": "0x...",
    "name": "John Doe"
  }
}
```

### Update Product
```http
PUT /products/:id
```

Request body:
```json
{
  "name": "Updated Name",
  "description": "Updated description",
  "price": "0.6",
  "images": [
    "https://...",
    "https://..."
  ],
  "category": "Footwear",
  "condition": "New"
}
```

Response:
```json
{
  "id": "1",
  "name": "Updated Name",
  "description": "Updated description",
  "price": "0.6",
  "images": [
    "https://...",
    "https://..."
  ],
  "category": "Footwear",
  "condition": "New",
  "seller": {
    "walletAddress": "0x...",
    "name": "John Doe"
  }
}
```

### Delete Product
```http
DELETE /products/:id
```

Response:
```json
{
  "success": true
}
```

### Authenticate Product
```http
POST /products/:id/authenticate
```

Request body:
```json
{
  "images": [
    "https://...",
    "https://..."
  ]
}
```

Response:
```json
{
  "authenticated": true,
  "score": 0.95,
  "details": "Product authenticity verified with high confidence"
}
```

## Users

### Get User
```http
GET /users/:address
```

Response:
```json
{
  "walletAddress": "0x...",
  "name": "John Doe",
  "avatar": "https://...",
  "listings": 12,
  "sales": 8,
  "purchases": 15
}
```

### Update User
```http
PUT /users/:address
```

Request body:
```json
{
  "name": "John Doe",
  "avatar": "https://..."
}
```

Response:
```json
{
  "walletAddress": "0x...",
  "name": "John Doe",
  "avatar": "https://...",
  "listings": 12,
  "sales": 8,
  "purchases": 15
}
```

### Get User Products
```http
GET /users/:address/products
```

Query parameters:
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)
- `status` (optional): Filter by status (active, sold)

Response:
```json
{
  "products": [
    {
      "id": "1",
      "name": "Limited Edition Sneaker",
      "description": "Exclusive limited edition sneaker",
      "price": "0.5",
      "image": "https://...",
      "category": "Footwear",
      "status": "active"
    }
  ],
  "total": 12,
  "page": 1,
  "limit": 10
}
```

### Get User Orders
```http
GET /users/:address/orders
```

Query parameters:
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)
- `status` (optional): Filter by status (pending, completed, cancelled)

Response:
```json
{
  "orders": [
    {
      "id": "1",
      "product": {
        "id": "1",
        "name": "Limited Edition Sneaker",
        "image": "https://..."
      },
      "price": "0.5",
      "status": "completed",
      "completedAt": "2024-03-23T12:00:00Z"
    }
  ],
  "total": 15,
  "page": 1,
  "limit": 10
}
```

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid input",
  "message": "Detailed error message"
}
```

### 401 Unauthorized
```json
{
  "error": "Unauthorized",
  "message": "Invalid or expired token"
}
```

### 403 Forbidden
```json
{
  "error": "Forbidden",
  "message": "Insufficient permissions"
}
```

### 404 Not Found
```json
{
  "error": "Not Found",
  "message": "Resource not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "Internal Server Error",
  "message": "An unexpected error occurred"
}
``` 