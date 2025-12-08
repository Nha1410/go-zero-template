# Code Generation Guide

This guide explains how to generate code using `goctl` CLI tool.

## Prerequisites

Install `goctl` CLI tool:

```bash
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

Verify installation:

```bash
goctl --version
```

## Quick Reference

### Generate API Gateway Code

```bash
goctl api go -api api/api/api.api -dir api --style gozero
```

### Generate gRPC Service Code

```bash
goctl rpc protoc service/<service-name>/<service-name>.proto \
    --go_out=service/<service-name> \
    --go-grpc_out=service/<service-name> \
    --zrpc_out=service/<service-name> \
    --style gozero
```

### Generate Database Models

```bash
goctl model pg datasource \
    -url "host=localhost port=5432 user=postgres password=postgres dbname=gozero_template sslmode=disable" \
    -table "*" \
    -dir service/user/internal/model \
    -cache
```

## Detailed Guide

### 1. Generate API Gateway Code

Generate HTTP handlers, logic, and types from `.api` definition file.

#### Basic Usage

```bash
# Generate from default file (api/api/api.api)
goctl api go -api api/api/api.api -dir api --style gozero
```

#### Custom API File

```bash
# Generate from custom .api file
goctl api go -api api/custom/api.api -dir api --style gozero
```

#### What Gets Generated

- `api/internal/handler/` - HTTP request handlers
- `api/internal/logic/` - Business logic layer
- `api/internal/types/` - Request/Response type definitions
- Route registration code

#### Example API Definition

Create `api/api/api.api`:

```go
syntax = "v1"

info (
    title: "API Gateway"
    desc: "API Gateway for microservices"
    version: "1.0"
)

type (
    CreateUserRequest {
        Email string `json:"email"`
        Name  string `json:"name"`
    }

    BaseResponse {
        Code    int         `json:"code"`
        Message string      `json:"message"`
        Data    interface{} `json:"data,optional"`
    }
)

service api {
    @handler CreateUser
    post /api/v1/users (CreateUserRequest) returns (BaseResponse)
}
```

After running the command, you'll get:
- `CreateUserHandler` in `api/internal/handler/`
- `CreateUserLogic` in `api/internal/logic/`
- Type definitions in `api/internal/types/`

---

### 2. Generate gRPC Service Code

Generate gRPC server code, handlers, and client stubs from `.proto` file.

#### Basic Usage

```bash
# Generate user service
goctl rpc protoc service/user/user.proto \
    --go_out=service/user \
    --go-grpc_out=service/user \
    --zrpc_out=service/user \
    --style gozero
```

#### Generate Other Services

```bash
# Generate order service
goctl rpc protoc service/order/order.proto \
    --go_out=service/order \
    --go-grpc_out=service/order \
    --zrpc_out=service/order \
    --style gozero

# Generate product service
goctl rpc protoc service/product/product.proto \
    --go_out=service/product \
    --go-grpc_out=service/product \
    --zrpc_out=service/product \
    --style gozero
```

#### What Gets Generated

- `service/<name>/userclient/` - gRPC client code (if service name is "user")
- `service/<name>/internal/handler/` - gRPC request handlers
- `service/<name>/*.pb.go` - Generated protobuf code
- `service/<name>/*_grpc.pb.go` - Generated gRPC code

#### Example Proto Definition

Create `service/user/user.proto`:

```protobuf
syntax = "proto3";

package user;

option go_package = "./userclient";

service User {
    rpc CreateUser(CreateUserReq) returns (CreateUserResp);
    rpc GetUser(GetUserReq) returns (GetUserResp);
    rpc GetUsers(GetUsersReq) returns (GetUsersResp);
}

message CreateUserReq {
    string email = 1;
    string name = 2;
}

message CreateUserResp {
    int64 id = 1;
    string email = 2;
    string name = 3;
}

message GetUserReq {
    int64 id = 1;
}

message GetUserResp {
    int64 id = 1;
    string email = 2;
    string name = 3;
}

message GetUsersReq {
    int64 page = 1;
    int64 page_size = 2;
}

message GetUsersResp {
    repeated GetUserResp users = 1;
    int64 total = 2;
}
```

After running the command, you'll get:
- `UserHandler` with methods: `CreateUser`, `GetUser`, `GetUsers`
- Client interface in `userclient` package
- All protobuf generated code

---

### 3. Generate Database Models

Generate Go models from PostgreSQL database schema.

#### Basic Usage

```bash
# Generate all tables
goctl model pg datasource \
    -url "host=localhost port=5432 user=postgres password=postgres dbname=gozero_template sslmode=disable" \
    -table "*" \
    -dir service/user/internal/model \
    -cache
```

#### Generate Specific Table

```bash
# Generate only users table
goctl model pg datasource \
    -url "host=localhost port=5432 user=postgres password=postgres dbname=gozero_template sslmode=disable" \
    -table "users" \
    -dir service/user/internal/model \
    -cache
```

#### Custom Output Directory

```bash
# Generate to custom directory
goctl model pg datasource \
    -url "host=localhost port=5432 user=postgres password=postgres dbname=gozero_template sslmode=disable" \
    -table "*" \
    -dir service/order/internal/model \
    -cache
```

#### What Gets Generated

- Model structs matching database tables
- CRUD methods with caching support
- Query builders
- Type-safe database operations

#### Example Output

For a `users` table:

```go
type Users struct {
    Id        int64     `db:"id"`
    Email     string    `db:"email"`
    Name      string    `db:"name"`
    CreatedAt time.Time `db:"created_at"`
    UpdatedAt time.Time `db:"updated_at"`
}
```

With methods like:
- `FindOne(id int64)` - Find by ID
- `FindOneByEmail(email string)` - Find by email
- `Insert(data Users)` - Insert new record
- `Update(data Users)` - Update record
- `Delete(id int64)` - Delete record

---

### 4. Generate Dockerfile

Generate Dockerfile for a service.

```bash
# Generate Dockerfile for API Gateway
goctl docker -go api/main.go

# Generate Dockerfile for user service
goctl docker -go service/user/main.go
```

---

### 5. Generate Kubernetes Files

Generate Kubernetes deployment files.

```bash
# Generate for API Gateway
goctl kube deploy \
    -name api-gateway \
    -namespace default \
    -image api-gateway:latest \
    -port 8888 \
    -o k8s/

# Generate for user service
goctl kube deploy \
    -name user-service \
    -namespace default \
    -image user-service:latest \
    -port 9000 \
    -o k8s/
```

---

## Common Workflows

### Workflow 1: Create New API Endpoint

1. Edit `api/api/api.api` to add new endpoint
2. Generate code:
   ```bash
   goctl api go -api api/api/api.api -dir api --style gozero
   ```
3. Implement business logic in `api/internal/logic/`
4. Test the endpoint

### Workflow 2: Create New gRPC Service

1. Create service directory:
   ```bash
   mkdir -p service/order
   ```

2. Create proto file: `service/order/order.proto`

3. Generate code:
   ```bash
   goctl rpc protoc service/order/order.proto \
       --go_out=service/order \
       --go-grpc_out=service/order \
       --zrpc_out=service/order \
       --style gozero
   ```

4. Implement clean architecture layers:
   - Domain entities in `service/order/internal/domain/entity/`
   - Repository interface in `service/order/internal/domain/repository/`
   - Use case in `service/order/internal/usecase/`
   - Repository implementation in `service/order/internal/repository/`

5. Register service in `service/order/main.go`

### Workflow 3: Generate Models from Database

1. Create database tables

2. Generate models:
   ```bash
   goctl model pg datasource \
       -url "host=localhost port=5432 user=postgres password=postgres dbname=gozero_template sslmode=disable" \
       -table "*" \
       -dir service/user/internal/model \
       -cache
   ```

3. Use models in repository layer:
   ```go
   import "github.com/Nha1410/go-zero-template/service/user/internal/model"

   func (r *userRepo) GetByID(ctx context.Context, id int64) (*entity.User, error) {
       // Use generated model
       userModel := model.NewUsersModel(r.db)
       result, err := userModel.FindOne(ctx, id)
       // Convert to domain entity
       return convertToEntity(result), err
   }
   ```

---

## goctl Commands Reference

### API Generation

```bash
goctl api go -api <api-file> -dir <output-dir> --style gozero
```

### RPC Generation

```bash
goctl rpc protoc <proto-file> \
    --go_out=<output-dir> \
    --go-grpc_out=<output-dir> \
    --zrpc_out=<output-dir> \
    --style gozero
```

### Model Generation

```bash
goctl model pg datasource \
    -url "<dsn>" \
    -table "<table-name>" \
    -dir <output-dir> \
    -cache
```

---

## Troubleshooting

### Error: goctl is not installed

```bash
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

### Error: proto file not found

Make sure the proto file exists at the expected path:
- For user service: `service/user/user.proto`
- For other services: `service/<service-name>/<service-name>.proto`

### Error: API file not found

Make sure the API file exists:
- Default: `api/api/api.api`
- Or specify custom path in the command

### Generated code conflicts with existing code

1. Backup your custom code
2. Regenerate code
3. Merge your customizations back

### Database connection failed during model generation

- Verify database is running
- Check connection string format
- Ensure database credentials are correct
- Test connection: `psql -h localhost -U postgres -d gozero_template`

---

## Best Practices

1. **Version Control**: Commit `.api` and `.proto` files, but consider ignoring generated code (add to `.gitignore`)

2. **Regenerate After Changes**: Always regenerate code after modifying `.api` or `.proto` files

3. **Don't Edit Generated Code**: Never edit generated code directly. Instead:
   - Modify source files (`.api`, `.proto`)
   - Regenerate code
   - Add custom logic in separate files

4. **Use goctl Commands**: Use `goctl` commands directly for full control and flexibility

5. **Test After Generation**: Always test your code after regeneration to ensure nothing broke

---

## Additional Resources

- [go-zero Documentation](https://go-zero.dev/)
- [goctl CLI Reference](https://go-zero.dev/docs/tasks/goctl)
- [Protocol Buffers Guide](https://developers.google.com/protocol-buffers)
- [go-zero Examples](https://github.com/zeromicro/go-zero/tree/master/example)

