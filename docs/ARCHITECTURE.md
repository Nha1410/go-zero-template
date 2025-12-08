# Architecture Documentation

## Overview

This template implements a microservices architecture using go-zero framework with Clean Architecture principles. The system consists of an API Gateway that routes HTTP requests to various gRPC microservices.

## Architecture Layers

### Clean Architecture

The project follows Clean Architecture principles with clear separation of concerns:

```
┌─────────────────────────────────────┐
│         Handler Layer               │  (gRPC/HTTP Handlers)
├─────────────────────────────────────┤
│         Use Case Layer              │  (Business Logic)
├─────────────────────────────────────┤
│         Repository Interface        │  (Domain Layer)
├─────────────────────────────────────┤
│         Repository Implementation   │  (Data Access)
└─────────────────────────────────────┘
```

#### 1. Domain Layer (`internal/domain/`)

- **Entities**: Core business objects (e.g., `User`)
- **Repository Interfaces**: Contracts for data access

**Example:**
```go
// entity/user.go
type User struct {
    ID        int64
    Email     string
    Name      string
    CreatedAt time.Time
    UpdatedAt time.Time
}

// repository/user_repository.go
type UserRepository interface {
    Create(ctx context.Context, user *User) error
    GetByID(ctx context.Context, id int64) (*User, error)
    // ...
}
```

#### 2. Use Case Layer (`internal/usecase/`)

Contains business logic and application rules. This layer depends only on repository interfaces, not implementations.

**Example:**
```go
type UserUsecase struct {
    userRepo repository.UserRepository
}

func (uc *UserUsecase) CreateUser(ctx context.Context, email, name string) (*entity.User, error) {
    // Business logic here
    // - Validate input
    // - Check business rules
    // - Call repository
}
```

#### 3. Repository Layer (`internal/repository/`)

Implements data access logic. This layer implements the repository interfaces defined in the domain layer.

**Example:**
```go
type userRepo struct {
    db *sql.DB
}

func (r *userRepo) Create(ctx context.Context, user *entity.User) error {
    // Database operations
}
```

#### 4. Handler Layer (`internal/handler/`)

gRPC handlers that receive requests and delegate to use cases.

**Example:**
```go
func (h *UserHandler) CreateUser(ctx context.Context, req *userclient.CreateUserReq) (*userclient.CreateUserResp, error) {
    l := logic.NewCreateUserLogic(ctx, h.svcCtx)
    return l.CreateUser(req)
}
```

## Microservices Architecture

### API Gateway Pattern

```
┌─────────────┐
│   Client    │
└──────┬──────┘
       │ HTTP
       ▼
┌─────────────┐
│ API Gateway │
└──────┬──────┘
       │ gRPC
       ├──────────┬──────────┬──────────┐
       ▼          ▼          ▼          ▼
   ┌──────┐  ┌──────┐  ┌──────┐  ┌──────┐
   │User  │  │Order │  │Product│ │Other │
   │Service│  │Service│ │Service│ │Service│
   └──────┘  └──────┘  └──────┘  └──────┘
```

### API Gateway (`api/`)

- Receives HTTP requests
- Handles authentication/authorization
- Routes requests to appropriate microservices via gRPC
- Aggregates responses
- Handles cross-cutting concerns (logging, metrics, etc.)

### Microservices (`service/`)

Each microservice:
- Is independently deployable
- Has its own database (database per service pattern)
- Communicates via gRPC
- Follows Clean Architecture
- Can scale independently

## Communication Patterns

### Synchronous Communication

- **API Gateway → Services**: gRPC calls
- **Service → Service**: Direct gRPC calls (when needed)

### Asynchronous Communication

- **Message Queue**: RabbitMQ for async processing
- **Events**: Services can publish events to RabbitMQ

## Database Strategy

### Database per Service

Each microservice has its own database:
- **User Service**: `users` table
- **Other Services**: Their own tables

This ensures:
- Service independence
- Data isolation
- Independent scaling

### Supported Databases

- **PostgreSQL**: Database used by the template

## Authentication & Authorization

### OAuth2 with Zitadel

1. Client sends request with Bearer token
2. API Gateway validates token with Zitadel
3. User info is extracted and added to context
4. Request is forwarded to service with user context

### Flow

```
Client → API Gateway (Auth Middleware) → Zitadel (Validate Token)
                                      ↓
                              Service (with user context)
```

## Caching Strategy

### Redis Caching

- **Purpose**: Reduce database load
- **Usage**: Cache frequently accessed data
- **Pattern**: Cache-aside pattern

**Example:**
```go
// Check cache first
user, err := redis.GetJSON("user:" + id, &user)
if err != nil {
    // Cache miss - fetch from database
    user = db.GetUser(id)
    // Store in cache
    redis.SetJSON("user:" + id, user, 1*time.Hour)
}
```

## Message Queue

### RabbitMQ

Used for:
- Async task processing
- Event-driven communication
- Decoupling services

**Example:**
```go
// Publish event
rabbitmq.Publish("user.created", userEvent)

// Consume events
msgs, _ := rabbitmq.Consume("user.created", "consumer", false, false, false, false)
```

## Error Handling

### Error Types

- **Domain Errors**: Business logic errors
- **Infrastructure Errors**: Database, network errors
- **Validation Errors**: Input validation errors

### Error Response Format

```json
{
  "code": "NOT_FOUND",
  "message": "Resource not found",
  "details": "User with ID 123 not found"
}
```

## Logging

### Structured Logging

- Uses go-zero's logging framework
- Logs include context (request ID, user ID, etc.)
- Different log levels (debug, info, warn, error)

## Monitoring & Observability

### Metrics

- Custom business metrics
- System metrics (CPU, memory, etc.)

### Tracing

- Request tracing across services
- Distributed tracing support

## Deployment

### Docker

Each service has its own Dockerfile:
- Multi-stage builds
- Optimized for production
- Minimal image size

### Docker Compose

Local development setup includes:
- All databases
- Redis
- RabbitMQ
- All services

## Best Practices

1. **Dependency Rule**: Dependencies point inward (toward domain)
2. **Interface Segregation**: Small, focused interfaces
3. **Single Responsibility**: Each layer has one responsibility
4. **Testability**: Easy to test with dependency injection
5. **Independence**: Services can be developed/deployed independently

## Adding a New Service

1. Create service directory structure
2. Define proto file
3. Generate code with goctl
4. Implement domain layer (entities, repository interfaces)
5. Implement use case layer
6. Implement repository layer
7. Implement handler layer
8. Add to docker-compose.yml
9. Configure service discovery

## References

- [go-zero Documentation](https://go-zero.dev/)
- [Clean Architecture by Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Microservices Patterns](https://microservices.io/patterns/)

