# API Gateway

A high-performance API Gateway built with Go that routes requests to various microservices in the system.

## Features

- Service routing and load balancing
- Configuration-based service management
- Docker containerization support
- Hot-reload development environment (using Air)
- Multiple service endpoints support

## Prerequisites

- Go 1.24 or higher
- Docker and Docker Compose
- Make (optional, for using Makefile commands)

## Project Structure

```
api-gateway/
├── cmd/                    # Main application entry point
│   └── main.go            # Application entry point and server setup
│
├── internal/              # Internal packages and business logic
│   ├── breaker/          # Circuit breaker implementation
│   ├── config/           # Configuration loading and parsing
│   ├── lb/               # Load balancing strategies
│   └── proxy/            # Reverse proxy implementation
│
├── api-config.toml        # Service configuration file
├── Dockerfile            # Container build configuration
├── docker-compose.yml    # Service orchestration
├── .air.toml             # Hot-reload configuration
├── .dockerignore         # Docker ignore patterns
├── .gitignore            # Git ignore patterns
├── go.mod                # Go module definition
└── go.sum                # Go module checksums
```

## Configuration

The API Gateway is configured using `api-config.toml`. Here's an example configuration:

```toml
[[service]]
name = "master-api"
url  = "http://master-service:8082"

[[service]]
name = "product-api"
url  = "http://product-service:8081"

[[service]]
name = "user-api"
url  = "http://user-service:8083"
```

## Getting Started

### Local Development

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Start the development server with hot-reload:
   ```bash
   air
   ```

### Using Docker

1. Build and run using Docker Compose:
   ```bash
   docker-compose up --build
   ```

2. The API Gateway will be available at `http://localhost:8090`

## Environment Variables

- `PORT`: The port the API Gateway listens on (default: 8090)
- `CONFIG_PATH`: Path to the configuration file (default: ./api-config.toml)

## API Endpoints

The API Gateway routes requests to the following services:

- Master API: `http://master-service:8082`
- Product API: `http://product-service:8081`
- User API: `http://user-service:8083`

## Development

### Hot Reload

The project uses [Air](https://github.com/cosmtrek/air) for hot-reload during development. Configuration is in `.air.toml`.

## Deployment

### Docker Deployment

1. Build the Docker image:
   ```bash
   docker build -t api-gateway .
   ```

2. Run the container:
   ```bash
   docker run -p 8090:8090 -v $(pwd)/api-config.toml:/api-config.toml api-gateway
   ```

### Docker Compose Deployment

The `docker-compose.yml` file includes configuration for all services:

```bash
docker-compose up -d
```