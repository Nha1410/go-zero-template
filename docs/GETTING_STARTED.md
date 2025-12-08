# Getting Started Guide

This guide will help you set up and run the go-zero template project.

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.21+**: [Download Go](https://golang.org/dl/)
- **Docker & Docker Compose**: [Install Docker](https://docs.docker.com/get-docker/)
- **goctl CLI**: Install with `go install github.com/zeromicro/go-zero/tools/goctl@latest`
- **PostgreSQL**: For database
- **Redis**: For caching
- **RabbitMQ**: For message queue
- **Zitadel**: OAuth2 provider (or use a test instance)

## Installation

### 1. Clone the Repository

```bash
git clone <your-repo-url>
cd go-zero-template
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Install goctl

```bash
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

Verify installation:
```bash
goctl --version
```

## Configuration

### 1. Database Setup

#### Option A: Using Docker Compose (Recommended for Development)

```bash
cd deployments
docker-compose up -d postgres redis rabbitmq
```

This will start:
- PostgreSQL on port 5432
- Redis on port 6379
- RabbitMQ on port 5672 (Management UI on 15672)

#### Option B: Local Installation

Install and configure PostgreSQL, Redis, and RabbitMQ locally.

### 2. Configure Environment Variables

#### For Local Development (without Docker)

Copy the example environment file:

```bash
cp env.example .env
```

Edit `.env` with your actual values for database, Redis, RabbitMQ, and Zitadel.

#### For Docker Compose

Copy the example environment file in deployments directory:

```bash
cd deployments
cp env.example .env
```

Edit `deployments/.env` with your actual values. Docker Compose will automatically load this file.

### 3. Initialize Database

#### Option A: Using Docker Compose

The database will be automatically initialized when you start the containers:

```bash
cd deployments
docker-compose up -d postgres
```

The initialization script (`init.sql`) will run automatically on first start.

#### Option B: Manual Setup

Connect to PostgreSQL and create the database:

```bash
psql -U postgres -h localhost
CREATE DATABASE gozero_template;
```

Run the initialization script:

```bash
psql -U postgres -h localhost -d gozero_template -f deployments/init.sql
```

### 4. Configure Services

#### Environment Variables

This project uses `.env` files for configuration. All configuration is loaded from environment variables.

**Setup:**
1. Copy the example environment file:
```bash
cp .env.example .env
```

2. Edit `.env` with your actual values (database, Redis, RabbitMQ, Zitadel credentials)

3. Services will automatically load from `.env` file when running locally

**For Docker:**
- Environment variables are set in `docker-compose.yml`
- You can also use `.env` file which will be loaded automatically

## Code Generation

### Generate API Gateway Code

```bash
goctl api go -api api/api/api.api -dir api --style gozero
```

This will generate:
- Handler code from `.api` file
- Type definitions
- Route handlers

### Generate User Service Code

```bash
goctl rpc protoc service/user/user.proto \
    --go_out=service/user \
    --go-grpc_out=service/user \
    --zrpc_out=service/user \
    --style gozero
```

This will generate:
- gRPC server code from `.proto` file
- Client stubs
- Service registration code

### Generate Models from Database

```bash
goctl model pg datasource \
    -url "host=localhost port=5432 user=postgres password=postgres dbname=gozero_template sslmode=disable" \
    -table "*" \
    -dir service/user/internal/model \
    -cache
```

See [CODE_GENERATION.md](CODE_GENERATION.md) for detailed code generation guide.

## Running the Services

### Development Mode

#### 1. Start Infrastructure

```bash
cd deployments
docker-compose up -d postgres redis rabbitmq
```

#### 2. Run User Service

```bash
cd service/user
go run main.go
```

The service will start on `localhost:9000`.

#### 3. Run API Gateway

In another terminal:

```bash
cd api
go run main.go
```

The API Gateway will start on `localhost:8888`.

### Using Docker Compose

Run everything with Docker Compose:

```bash
cd deployments
docker-compose up --build
```

This will build and start all services.

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
goctl api go -api api/api/api.api -dir api --style gozero
```

### 2. Modify Proto Definition

Edit `service/user/user.proto` and regenerate:

```bash
goctl rpc protoc service/user/user.proto \
    --go_out=service/user \
    --go-grpc_out=service/user \
    --zrpc_out=service/user \
    --style gozero
```

### 3. Add New Service

1. Create service directory: `service/your-service/`
2. Create proto file: `service/your-service/your-service.proto`
3. Generate code:
   ```bash
   goctl rpc protoc service/your-service/your-service.proto \
       --go_out=service/your-service \
       --go-grpc_out=service/your-service \
       --zrpc_out=service/your-service \
       --style gozero
   ```
4. Implement clean architecture layers
5. Add to `deployments/docker-compose.yml`

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
# Docker logs
docker-compose -f deployments/docker-compose.yml logs -f

# Service logs
tail -f service/user/logs/*.log
tail -f api/logs/*.log
```

### Database Migrations

The template includes a migration script placeholder. You can integrate with:
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [dbmate](https://github.com/amacneil/dbmate)

### Debugging

1. **Enable debug logging**: Set log level to `debug` in config files
2. **Use gRPC reflection**: Enable in service config (`Mode: dev`)
3. **Use grpcurl**: Test gRPC services directly

```bash
# Install grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# List services
grpcurl -plaintext localhost:9000 list

# Call service
grpcurl -plaintext localhost:9000 user.User/GetUser
```

## Troubleshooting

### Port Already in Use

If you get port conflicts:
- Change ports in configuration files
- Or stop existing services using those ports

### Database Connection Failed

- Check database is running: `docker ps`
- Verify credentials in config files
- Check network connectivity

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

