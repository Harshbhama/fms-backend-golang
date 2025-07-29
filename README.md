# Go Microservices with Docker, PostgreSQL, and Redis

This project demonstrates a microservice architecture using Go, Docker, PostgreSQL, and Redis. It consists of the following components:

1. **API Gateway (NGINX)**: Routes requests to appropriate services
2. **Auth Service**: Handles user authentication and authorization
3. **Client Service**: A protected service that requires authentication
4. **PostgreSQL**: Database for persistent storage
5. **Redis**: For caching and session management

## Prerequisites

- Docker and Docker Compose
- Go 1.21 or later (for local development)

## Getting Started

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd backend-fms
   ```

2. **Start the services**
   ```bash
   docker-compose up --build
   ```
   This will start:
   - PostgreSQL database
   - Redis cache
   - Auth Service (port 8080)
   - Client Service (port 8081)

## API Endpoints

All endpoints are accessed through the API Gateway (port 80)

### Auth Service
- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Authenticate a user
- `GET /api/auth/health` - Health check

### Client Service
- `GET /api/client/protected` - Protected endpoint (requires authentication)
- `GET /api/client/health` - Health check

### API Gateway
- `GET /health` - API Gateway health check

## Environment Variables

### API Gateway
- `NGINX_PORT` - Port to run the API Gateway on (default: 80)

### Auth Service
- `DB_HOST` - PostgreSQL host (default: postgres)
- `DB_USER` - Database username (default: admin)
- `DB_PASSWORD` - Database password (default: admin123)
- `DB_NAME` - Database name (default: auth_db)
- `DB_PORT` - Database port (default: 5432)
- `REDIS_ADDR` - Redis address (default: redis:6379)
- `REDIS_PASSWORD` - Redis password (default: redispass)

### Client Service
- `AUTH_SERVICE_URL` - URL of the auth service (default: http://auth-service:8080)
- `REDIS_ADDR` - Redis address (default: redis:6379)
- `REDIS_PASSWORD` - Redis password (default: redispass)

## Development

### Running locally (without Docker)

1. Start PostgreSQL and Redis using Docker Compose:
   ```bash
   docker-compose up -d postgres redis
   ```

2. Run the auth service:
   ```bash
   cd auth-service
   go run main.go
   ```

3. Run the client service:
   ```bash
   cd client-service
   go run main.go
   ```

## Project Structure

```
backend-fms/
├── auth-service/         # Authentication service
│   ├── main.go          # Entry point
│   ├── Dockerfile       # Docker configuration
│   └── go.mod          # Go dependencies
├── client-service/      # Client service
│   ├── main.go         # Entry point
│   ├── Dockerfile      # Docker configuration
│   └── go.mod         # Go dependencies
├── nginx/              # API Gateway configuration
│   ├── nginx.conf      # Main NGINX configuration
│   └── conf.d/         # Server blocks
│       └── api-gateway.conf  # API Gateway routes
├── docker-compose.yml  # Docker Compose configuration
└── README.md          # This file
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
