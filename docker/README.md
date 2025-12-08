# Docker Configuration

This directory contains Docker configurations for the project.

## Structure

- `devbox/` - Development container with all development tools (goctl, golangci-lint, etc.)
- `api/` - Dockerfile for API Gateway service
- `user/` - Dockerfile for User service

## Devbox

The devbox container is used for all development commands via Makefile. It includes:

- Go 1.24
- goctl CLI tool
- golangci-lint
- git, make, bash, curl

### Setup

Build the devbox image:

```bash
make setup-devbox
```

### Usage

All development commands in Makefile automatically use the devbox container:

```bash
make generate-api      # Runs goctl in devbox
make lint             # Runs golangci-lint in devbox
make mod              # Runs go mod commands in devbox
make fmt              # Formats code in devbox
```

## Service Dockerfiles

Service Dockerfiles are minimal and only include what's needed to run the service:

- Multi-stage builds for smaller images
- Only runtime dependencies (no dev tools)
- Optimized for production use

### API Gateway

Located at `docker/api/Dockerfile`

### User Service

Located at `docker/user/Dockerfile`

