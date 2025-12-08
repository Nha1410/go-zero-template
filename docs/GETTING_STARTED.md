# Getting Started Guide

This guide will help you set up and run the go-zero template project.

## Prerequisites

Before you begin, ensure you have the following installed:

- **Docker & Docker Compose**: [Install Docker](https://docs.docker.com/get-docker/)
- **Make**: Usually pre-installed on Unix systems
- **Zitadel**: OAuth2 provider (or use a test instance)

**Note**: All development tools (Go, goctl, golangci-lint, etc.) are included in the devbox Docker image. PostgreSQL, Redis, and RabbitMQ are provided via Docker Compose.

## Installation

### 1. Clone the Repository

```bash
git clone <your-repo-url>
cd go-zero-template
```

### 2. Setup Development Environment

Build the devbox Docker image (contains all development tools):

```bash
make setup-devbox
```

This only needs to be done once.

### 3. Install Dependencies

```bash
make mod
```

This runs `go mod tidy`, `go mod verify`, and `go mod download` in the devbox container.

## Configuration

### 1. Database Setup

Start infrastructure services with Docker Compose:

```bash
make docker-up
```

This will start:
- PostgreSQL on port 5432
- Redis on port 6379
- RabbitMQ on port 5672 (Management UI on 15672)
- User Service (gRPC) on port 9000
- API Gateway on port 8888

### 2. Configure Environment Variables

Copy the example environment file in the root directory:

```bash
cp .env.example .env
```

Edit `.env` with your actual values for database, Redis, RabbitMQ, and Zitadel. Docker Compose will automatically load this file.

### 3. Initialize Database

The database will be automatically initialized when you start the containers:

```bash
make docker-up
```

The initialization script (`deployments/init.sql`) will run automatically on first start.

### 4. Configure Services

This project uses `.env` files for configuration. All configuration is loaded from environment variables.

**Setup:**
1. Copy the example environment file:
```bash
cp .env.example .env
```

2. Edit `.env` with your actual values (database, Redis, RabbitMQ, Zitadel credentials)

3. Docker Compose will automatically load the `.env` file from the root directory

## Code Generation

All code generation commands run in the devbox Docker container via Makefile.

### Generate API Gateway Code

```bash
make generate-api
```

This will generate:
- Handler code from `.api` file
- Type definitions
- Route handlers

### Generate User Service Code

```bash
make generate-service SERVICE=user
```

This will generate:
- gRPC server code from `.proto` file
- Client stubs
- Service registration code

### Other Development Commands

```bash
make fmt              # Format code
make lint             # Run linter
make mod              # Update dependencies
make update           # Update all dependencies
```

See [CODE_GENERATION.md](CODE_GENERATION.md) for detailed code generation guide.

## Running the Services

All services run in Docker containers. Start all services:

```bash
make docker-up
```

This will start:
- PostgreSQL database
- Redis cache
- RabbitMQ message queue
- User Service (gRPC) on port 9000
- API Gateway on port 8888

**View Logs:**
```bash
make docker-logs
```

**Stop Services:**
```bash
make docker-down
```

**Rebuild and Restart:**
```bash
cd deployments
docker-compose up -d --build
```

## Testing the API

### Health Check

```bash
curl http://localhost:8888/health
```

Expected response:
```json
{
  "code": 200,
  "message": "OK",
  "data": {
    "status": "healthy"
  }
}
```

### Create User (Requires Authentication)

```bash
curl -X POST http://localhost:8888/api/v1/users \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-zitadel-token>" \
  -d '{
    "email": "user@example.com",
    "name": "John Doe"
  }'
```

### Get User

```bash
curl http://localhost:8888/api/v1/users/1 \
  -H "Authorization: Bearer <your-zitadel-token>"
```

## Development Workflow

### 1. Modify API Definition

Edit `api/api/api.api` and regenerate:

```bash
make generate-api
```

### 2. Modify Proto Definition

Edit `service/user/user.proto` and regenerate:

```bash
make generate-service SERVICE=user
```

### 3. Add New Service

1. Create service directory: `service/your-service/`
2. Create proto file: `service/your-service/your-service.proto`
3. Generate code:
   ```bash
   make generate-service SERVICE=your-service
   ```
4. Implement clean architecture layers
5. Create Dockerfile: `docker/your-service/Dockerfile` (copy from `docker/user/Dockerfile` and adjust)
6. Add to `deployments/docker-compose.yml`

## Project Structure

```
go-zero-template/
├── api/                    # API Gateway
│   ├── api/                # API definitions (.api files)
│   ├── internal/           # Internal code
│   │   ├── handler/        # HTTP handlers
│   │   ├── logic/          # Business logic
│   │   ├── middleware/     # Middleware
│   │   └── svc/            # Service context
│   └── etc/                # Configuration
├── service/                # Microservices
│   └── user/               # User service example
│       ├── internal/
│       │   ├── domain/     # Domain layer
│       │   ├── usecase/    # Use case layer
│       │   ├── repository/ # Repository layer
│       │   └── handler/    # gRPC handlers
│       └── etc/            # Configuration
├── common/                 # Shared packages
└── deployments/            # Docker configs
```

## Common Tasks

### View Logs

```bash
# All services
make docker-logs

# Specific service
cd deployments
docker-compose logs -f api-gateway
docker-compose logs -f user-service
```

### Database Migrations

The template includes a migration script placeholder. You can integrate with:
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [dbmate](https://github.com/amacneil/dbmate)

### Debugging

1. **Enable debug logging**: Set `LOG_LEVEL=debug` in `.env` file
2. **Use gRPC reflection**: Already enabled in dev mode
3. **Use grpcurl**: Test gRPC services directly (install locally or use devbox)

```bash
# List services (from host)
grpcurl -plaintext localhost:9000 list

# Call service
grpcurl -plaintext localhost:9000 user.User/GetUser
```

**Note**: grpcurl can be installed locally or you can run it in the devbox container.

## Troubleshooting

### Port Already in Use

If you get port conflicts:
- Change ports in configuration files
- Or stop existing services using those ports

### Database Connection Failed

- Check database is running: `docker ps`
- Verify credentials in `.env` file
- Check network connectivity: `docker network ls`
- View database logs: `cd deployments && docker-compose logs postgres`

### gRPC Connection Failed

- Verify service is running
- Check service discovery configuration
- Ensure ports are accessible

### Authentication Errors

- Verify Zitadel configuration
- Check token validity
- Ensure token has required scopes

## Next Steps

1. Read [Architecture Documentation](ARCHITECTURE.md)
2. Explore the code structure
3. Add your own services
4. Customize authentication
5. Add more features

## Resources

- [go-zero Documentation](https://go-zero.dev/)
- [go-zero Examples](https://github.com/zeromicro/go-zero/tree/master/example)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## Support

For issues and questions:
- Check existing documentation
- Review go-zero documentation
- Open an issue in the repository

