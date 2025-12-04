# Docker Deployment Guide

This guide explains how to run the go-zero template services using Docker Compose.

## Prerequisites

- Docker and Docker Compose installed
- `.env` file in the root directory (optional, can use env vars directly)

## Quick Start

### 1. Setup Environment (Optional)

If you want to use a `.env` file, create one in the root directory:

```bash
cd ..
cp .env.example .env
# Edit .env with your values
```

### 2. Start All Services

```bash
cd deployments
docker-compose up -d
```

This will start:
- PostgreSQL database
- Redis cache
- RabbitMQ message queue
- User Service (gRPC)
- API Gateway

### 3. Check Service Status

```bash
docker-compose ps
```

### 4. View Logs

```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f api-gateway
docker-compose logs -f user-service
```

### 5. Stop Services

```bash
docker-compose down
```

### 6. Rebuild and Restart

```bash
docker-compose up -d --build
```

## Configuration

All services are configured via environment variables. You can:

1. **Use .env file**: Create `.env` in root directory and docker-compose will load it
2. **Set directly in docker-compose.yml**: Environment variables are already defined
3. **Override via command line**:
   ```bash
   DATABASE_PASSWORD=mypassword docker-compose up -d
   ```

## Service Endpoints

- **API Gateway**: http://localhost:8888
- **User Service (gRPC)**: localhost:9000
- **PostgreSQL**: localhost:5432
- **Redis**: localhost:6379
- **RabbitMQ Management UI**: http://localhost:15672 (guest/guest)

## Environment Variables

All configuration is done via environment variables. See `.env.example` in the root directory for all available options.

Key variables:
- `DATABASE_HOST`, `DATABASE_PORT`, `DATABASE_USER`, `DATABASE_PASSWORD`, `DATABASE_NAME`
- `REDIS_HOST`, `REDIS_PORT`, `REDIS_PASSWORD`
- `RABBITMQ_HOST`, `RABBITMQ_PORT`, `RABBITMQ_USER`, `RABBITMQ_PASSWORD`
- `API_NAME`, `API_HOST`, `API_PORT`
- `USER_SERVICE_NAME`, `USER_SERVICE_LISTEN_ON`, `USER_SERVICE_MODE`

## Troubleshooting

### Services won't start

1. Check logs: `docker-compose logs`
2. Verify environment variables are set correctly
3. Ensure ports are not already in use
4. Check database connection: `docker-compose exec postgres psql -U postgres`

### Rebuild after code changes

```bash
docker-compose up -d --build
```

### Clean restart (removes volumes)

```bash
docker-compose down -v
docker-compose up -d
```

