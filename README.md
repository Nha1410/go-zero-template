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
├── docker/                 # Docker configurations
│   ├── devbox/            # Development container (goctl, lint, etc.)
│   ├── api/               # API Gateway Dockerfile
│   └── user/              # User service Dockerfile
├── deployments/            # Deployment configs
│   └── docker-compose.yml   # Docker Compose cho local dev
└── docs/                   # Documentation
```

## Prerequisites

- Docker and Docker Compose
- Make
- Zitadel instance (for OAuth2)

**Note**:
- All development tools (goctl, golangci-lint, etc.) are included in the devbox Docker image
- PostgreSQL, Redis, and RabbitMQ are provided via Docker Compose
- You don't need to install Go, goctl, or any other tools locally

## Quick Start

### 1. Setup Development Environment

Build the devbox Docker image (contains goctl, golangci-lint, and other dev tools):

```bash
make setup-devbox
```

This only needs to be done once. The devbox image will be used for all development commands.

### 2. Clone and Setup

```bash
git clone <your-repo-url>
cd go-zero-template

# Download dependencies (runs in devbox)
make mod
```

### 3. Configure Environment

**This project uses `.env` files for configuration. YAML files are no longer needed.**

Copy the example environment file in the root directory:

```bash
cp .env.example .env
# Edit .env with your actual values
```

Update the configuration in `.env` with your database, Redis, RabbitMQ, and Zitadel credentials.

**Note**: All configuration is loaded from environment variables. The `.env` file in the root directory will be automatically loaded by Docker Compose. See `.env.example` for all available configuration options.

### 4. Start All Services

```bash
make docker-up
```

Or directly:

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
make generate-api

# Generate User service code
make generate-service SERVICE=user
```

See [docs/CODE_GENERATION.md](docs/CODE_GENERATION.md) for detailed code generation guide.

## Development

### Adding a New Service

1. Create service directory: `service/your-service/`
2. Create proto file: `service/your-service/your-service.proto`
3. Generate code:
   ```bash
   make generate-service SERVICE=your-service
   ```
4. Implement clean architecture layers
5. Add service to docker-compose.yml
6. Create Dockerfile in `docker/your-service/Dockerfile` (copy from `docker/user/Dockerfile` and adjust)

### Code Generation

Generate code using Makefile commands (runs in devbox container):

```bash
# Generate API Gateway from .api file
make generate-api

# Generate gRPC service from .proto file
make generate-service SERVICE=<service-name>

# Format code
make fmt

# Run linter
make lint

# Update dependencies
make update

# Run go mod tidy
make mod
```

All development commands run in the devbox Docker container, ensuring consistent environment across all developers.

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
cp .env.example .env
# Edit .env with your values
```

The `.env` file in the root directory will be automatically loaded by Docker Compose.

## Docker

All services run in Docker containers. See [deployments/README.md](deployments/README.md) for detailed Docker deployment guide.

**Quick Start:**
```bash
make docker-up
```

**View Logs:**
```bash
make docker-logs
```

**Stop Services:**
```bash
make docker-down
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

