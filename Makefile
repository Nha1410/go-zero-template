.PHONY: help install generate build run test clean docker-up docker-down

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	go mod download
	go install github.com/zeromicro/go-zero/tools/goctl@latest

generate-api: ## Generate API Gateway code
	./scripts/generate-api.sh

generate-service: ## Generate service code (usage: make generate-service SERVICE=user)
	./scripts/generate-service.sh $(SERVICE)

generate: generate-api generate-service ## Generate all code

build-api: ## Build API Gateway
	cd api && go build -o ../bin/api main.go

build-user: ## Build User service
	cd service/user && go build -o ../../bin/user main.go

build: build-api build-user ## Build all services

run-api: ## Run API Gateway
	cd api && go run main.go

run-user: ## Run User service
	cd service/user && go run main.go

docker-up: ## Start Docker Compose services
	cd deployments && docker-compose up -d

docker-down: ## Stop Docker Compose services
	cd deployments && docker-compose down

docker-logs: ## View Docker Compose logs
	cd deployments && docker-compose logs -f

test: ## Run tests
	go test ./...

clean: ## Clean build artifacts
	rm -rf bin/
	rm -rf api/logs/
	rm -rf service/*/logs/

