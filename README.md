# Go-Zero Clean Architecture Template

A production-ready template repository for building microservices APIs using [go-zero](https://github.com/zeromicro/go-zero) framework with Clean Architecture principles.

## Features

- 🏗️ **Clean Architecture** - Well-organized layers (Domain, Use Case, Repository, Handler)
- 🚀 **Microservices** - API Gateway pattern with independent gRPC services
- 🔐 **OAuth2 Authentication** - Integration with Zitadel for AAA (Authentication, Authorization, Accounting)
- 💾 **PostgreSQL Database** - PostgreSQL database support
- ⚡ **Redis Caching** - Built-in Redis client for caching
- 📨 **RabbitMQ** - Message queue support for async processing
- 🐳 **Docker Support** - Docker Compose setup for local development
- 🛠️ **goctl Integration** - Code generation using goctl CLI
- 📝 **Comprehensive Documentation** - Architecture and getting started guides

## Project Structure

```
go-zero-template/
├── api/                    # API Gateway service
│   ├── internal/
│   │   ├── handler/        # HTTP handlers
│   │   ├── logic/          # Business logic
│   │   ├── middleware/     # Middleware (auth, logging, etc.)
│   │   ├── svc/            # Service context
│   │   └── types/          # Generated types
│   ├── api/                # API definition files (.api)
│   └── main.go
├── service/                # Microservices
│   └── user/               # User service example
│       ├── internal/
│       │   ├── domain/     # Domain layer (clean arch)
│       │   ├── usecase/    # Use case layer
│       │   ├── repository/ # Repository implementation
│       │   ├── handler/    # gRPC handlers
│       │   ├── logic/      # Business logic
│       │   └── svc/        # Service context
│       ├── user.proto      # gRPC proto definition
│       └── main.go
├── common/                 # Shared code
│   ├── auth/               # Authentication utilities
│   ├── cache/              # Redis cache utilities
│   ├── database/           # Database connection & migrations
│   ├── queue/              # RabbitMQ utilities
│   ├── logger/             # Logging utilities
│   ├── errors/             # Error handling
│   └── validator/          # Request validation
├── deployments/            # Deployment configs
│   └── docker-compose.yml   # Docker Compose cho local dev
└── docs/                   # Documentation
```

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- goctl CLI tool
- PostgreSQL
- Redis
- RabbitMQ
- Zitadel instance (for OAuth2)

## Quick Start

### 1. Install goctl

```bash
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

### 2. Clone and Setup

```bash
git clone <your-repo-url>
cd go-zero-template
go mod download
```

### 3. Configure Environment

**This project uses `.env` files for configuration. YAML files are no longer needed.**

Copy the example environment file and update with your values:

```bash
# For local development (outside Docker)
cp .env.example .env
# Edit .env with your actual values

# For Docker Compose
cd deployments
cp env.example .env
# Edit .env with your actual values
```

Update the configuration in `.env` with your database, Redis, RabbitMQ, and Zitadel credentials.

**Note**: All configuration is loaded from environment variables. See `.env.example` for all available configuration options.

### 4. Start All Services with Docker Compose (Recommended)

```bash
cd deployments
docker-compose up -d
```

This will start all services including:
- PostgreSQL database
- Redis cache
- RabbitMQ message queue
- User Service (gRPC)
- API Gateway

See [deployments/README.md](deployments/README.md) for detailed Docker deployment guide.

### 5. Generate Code (if needed)

```bash
# Generate API Gateway code
goctl api go -api api/api/api.api -dir api --style gozero

# Generate User service code
goctl rpc protoc service/user/user.proto \
    --go_out=service/user \
    --go-grpc_out=service/user \
    --zrpc_out=service/user \
    --style gozero
```

See [docs/CODE_GENERATION.md](docs/CODE_GENERATION.md) for detailed code generation guide.

### 6. Run Services Locally (Alternative to Docker)

If you prefer to run services locally instead of Docker:

```bash
# Make sure you have .env file in root directory
cp .env.example .env
# Edit .env with your values

# Run User service
cd service/user
go run main.go

# Run API Gateway (in another terminal)
cd api
go run main.go
```

## Development

### Adding a New Service

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
5. Add service to docker-compose.yml

### Code Generation

Generate code using `goctl` CLI tool:

```bash
# Generate API Gateway from .api file
goctl api go -api api/api/api.api -dir api --style gozero

# Generate gRPC service from .proto file
goctl rpc protoc service/<service-name>/<service-name>.proto \
    --go_out=service/<service-name> \
    --go-grpc_out=service/<service-name> \
    --zrpc_out=service/<service-name> \
    --style gozero

# Generate models from PostgreSQL database
goctl model pg datasource \
    -url "host=localhost port=5432 user=postgres password=postgres dbname=gozero_template sslmode=disable" \
    -table "*" \
    -dir service/user/internal/model \
    -cache
```

See [docs/CODE_GENERATION.md](docs/CODE_GENERATION.md) for detailed guide.

## Architecture

This template follows Clean Architecture principles:

- **Domain Layer**: Core business entities and repository interfaces
- **Use Case Layer**: Business logic and application rules
- **Repository Layer**: Data access implementation
- **Handler Layer**: gRPC/HTTP handlers

See [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) for detailed architecture documentation.

## Configuration

### Environment Variables

All configuration is done via environment variables. See `.env.example` for all available options.

**Setup:**
```bash
# For local development
cp .env.example .env
# Edit .env with your values
```

## Docker

See [deployments/README.md](deployments/README.md) for detailed Docker deployment guide.

**Quick Start:**
```bash
cd deployments
docker-compose up -d
```

## Authentication

The template integrates with Zitadel for OAuth2 authentication. Configure your Zitadel instance in `.env` file.

Protected endpoints require a Bearer token in the Authorization header:
```
Authorization: Bearer <your-token>
```

## Documentation

- [Architecture Documentation](docs/ARCHITECTURE.md)
- [Getting Started Guide](docs/GETTING_STARTED.md)
- [Code Generation Guide](docs/CODE_GENERATION.md)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.

