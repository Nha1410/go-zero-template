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
- 🛠️ **goctl Integration** - Helper scripts for code generation
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
│   ├── etc/                # Configuration files
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
│       ├── etc/            # Configuration
│       └── main.go
├── common/                 # Shared code
│   ├── auth/               # Authentication utilities
│   ├── cache/              # Redis cache utilities
│   ├── database/           # Database connection & migrations
│   ├── queue/              # RabbitMQ utilities
│   ├── logger/             # Logging utilities
│   ├── errors/             # Error handling
│   └── validator/          # Request validation
├── scripts/                # Helper scripts
│   ├── generate.sh         # Script generate code với goctl
│   └── migrate.sh          # Database migration script
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

#### Option A: Using Environment Variables (Recommended)

Copy the example environment file and update with your values:

```bash
# For local development (outside Docker)
cp env.example .env
# Edit .env with your actual values

# For Docker Compose
cd deployments
cp env.example .env
# Edit .env with your actual values
```

#### Option B: Direct Configuration Files

Copy and edit configuration files directly:

```bash
cp api/etc/api.yaml api/etc/api.yaml.local
cp service/user/etc/user.yaml service/user/etc/user.yaml.local
```

Update the configuration with your database, Redis, RabbitMQ, and Zitadel credentials.

**Note**: Configuration files support environment variable substitution using `${VARIABLE:-default}` syntax. Environment variables take precedence over default values in config files.

### 4. Start Infrastructure with Docker Compose

```bash
cd deployments
docker-compose up -d postgres redis rabbitmq
```

### 5. Generate Code

```bash
# Generate API Gateway code
./scripts/generate-api.sh

# Generate User service code
./scripts/generate-service.sh user
```

### 6. Run Services

```bash
# Run User service
cd service/user
go run main.go -f etc/user.yaml

# Run API Gateway (in another terminal)
cd api
go run main.go -f etc/api.yaml
```

## Development

### Adding a New Service

1. Create service directory: `service/your-service/`
2. Create proto file: `service/your-service/your-service.proto`
3. Generate code: `./scripts/generate-service.sh your-service`
4. Implement clean architecture layers
5. Add service to docker-compose.yml

### Code Generation

The template includes helper scripts for code generation:

```bash
# Generate API Gateway from .api file
./scripts/generate-api.sh

# Generate gRPC service from .proto file
./scripts/generate-service.sh <service-name>

# Generate models from PostgreSQL database
./scripts/generate.sh model 'host=localhost port=5432 user=postgres password=postgres dbname=gozero_template sslmode=disable'
```

## Architecture

This template follows Clean Architecture principles:

- **Domain Layer**: Core business entities and repository interfaces
- **Use Case Layer**: Business logic and application rules
- **Repository Layer**: Data access implementation
- **Handler Layer**: gRPC/HTTP handlers

See [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) for detailed architecture documentation.

## Configuration

### Environment Variables

The project supports environment variables for configuration. Two example files are provided:

- **`env.example`** - For local development (root directory)
- **`deployments/env.example`** - For Docker Compose (deployments directory)

**Setup:**
```bash
# For local development
cp env.example .env
# Edit .env with your values

# For Docker Compose
cd deployments
cp env.example .env
# Edit .env with your values
```

### Configuration Files

Configuration files use environment variable substitution with `${VARIABLE:-default}` syntax:
- `api/etc/api.yaml` - API Gateway configuration
- `service/user/etc/user.yaml` - User service configuration

Environment variables take precedence over default values in config files.

## Docker

### Build and Run with Docker Compose

```bash
cd deployments
docker-compose up --build
```

This will start:
- PostgreSQL database
- Redis cache
- RabbitMQ message queue
- User service (gRPC)
- API Gateway (HTTP)

### Individual Service Dockerfiles

Each service has its own Dockerfile:
- `api/Dockerfile` - API Gateway
- `service/user/Dockerfile` - User service

## Authentication

The template integrates with Zitadel for OAuth2 authentication. Configure your Zitadel instance in the configuration files.

Protected endpoints require a Bearer token in the Authorization header:
```
Authorization: Bearer <your-token>
```

## Documentation

- [Architecture Documentation](docs/ARCHITECTURE.md)
- [Getting Started Guide](docs/GETTING_STARTED.md)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.

