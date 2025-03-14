# Product Service API

This repository contains a RESTful API service for managing products, built with Go and MongoDB.

## Table of Contents

- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Running with Docker](#running-with-docker)
    - [Running Locally](#running-locally)
- [Environment Variables](#environment-variables)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)


## Getting Started

### Prerequisites

- Go 1.24.0 or higher
- Docker and Docker Compose (for containerized setup)
- MongoDB (if running locally)

### Running with Docker

The easiest way to run the application is using Docker Compose:

```bash
# Clone the repository
git clone https://github.com/Ahmad-mufied/aws-academy-group-2.git
cd aws-academy-group-2/product-service

# Start the application with Docker Compose
docker-compose up -d
```

This will start both the application and MongoDB. The API will be available at http://localhost:8081.

### Running Locally

If you prefer to run the application locally:

```bash
# Clone the repository
git clone https://github.com/Ahmad-mufied/aws-academy-group-2.git
cd aws-academy-group-2/product-service

# Set up environment variables (see Environment Variables section)
cp .env.example .env
# Edit .env with your MongoDB connection string

# Run the application
go run cmd/main.go
```

The API will be available at http://localhost:8080 (or the port specified in your .env file).

## Environment Variables

The application uses the following environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| MONGO_URI | MongoDB connection string | mongodb://admin:password@mongodb:27017/ |
| WEB_SERVER_PORT | Port for the web server | 8080 |

## API Endpoints

### Get All Products

```
GET /products
```

**Response:**
```json
[
  {
    "product_id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Product Name",
    "is_active": "active",
    "created_by": "123e4567-e89b-12d3-a456-426614174001",
    "updated_by": "123e4567-e89b-12d3-a456-426614174001",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
]
```

### Get Product by ID

```
GET /products/:id
```

**Response:**
```json
{
  "product_id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "Product Name",
  "is_active": "active",
  "created_by": "123e4567-e89b-12d3-a456-426614174001",
  "updated_by": "123e4567-e89b-12d3-a456-426614174001",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### Create Product

```
POST /products
```

**Request Body:**
```json
{
  "name": "New Product",
  "is_active": "active",
  "created_by": "123e4567-e89b-12d3-a456-426614174001",
  "updated_by": "123e4567-e89b-12d3-a456-426614174001"
}
```

**Response:**
```json
{
  "product_id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "New Product",
  "is_active": "active",
  "created_by": "123e4567-e89b-12d3-a456-426614174001",
  "updated_by": "123e4567-e89b-12d3-a456-426614174001",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### Update Product

```
PUT /products/:product_id
```

**Request Body:**
```json
{
  "name": "Updated Product",
  "is_active": "inactive",
  "updated_by": "123e4567-e89b-12d3-a456-426614174002"
}
```

**Response:**
```json
{
  "product_id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "Updated Product",
  "is_active": "inactive",
  "created_by": "123e4567-e89b-12d3-a456-426614174001",
  "updated_by": "123e4567-e89b-12d3-a456-426614174002",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### Delete Product

```
DELETE /products/:product_id
```

**Response:**
```
204 No Content
```

## Project Structure

```
product-service/
├── cmd/
│   └── main.go                 # Application entry point
├── config/
│   ├── mongodb.go              # MongoDB configuration
│   └── viper.go                # Environment configuration
├── domain/
│   └── product.go              # Domain models and interfaces
├── dto/
│   └── product.go              # Data transfer objects
├── logger/
│   └── logger.go               # Logging configuration
├── repository/
│   └── mongoProductRepository.go # MongoDB implementation
├── server/
│   ├── productHandler.go       # HTTP handlers
│   └── routes.go               # Route definitions
├── service/
│   └── productService.go       # Business logic
├── utils/
│   └── uuidcodec.go            # UUID utilities
├── .dockerignore
├── .env.example
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```
