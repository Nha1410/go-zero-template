# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/user ./service/user/main.go

# Final stage
FROM golang:1.24-alpine

RUN apk --no-cache add ca-certificates tzdata git

# Install goimports
RUN go install golang.org/x/tools/cmd/goimports@latest

# Install golangci-lint
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/bin/user ./user

# Copy source code for development commands
COPY --from=builder /app/go.mod /app/go.sum ./
COPY --from=builder /app/service ./service
COPY --from=builder /app/common ./common

EXPOSE 9000

CMD ["./user"]

