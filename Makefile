MAKEFILE_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
DEVBOX := go-zero-template/devbox
INDEVBOX := docker run --network host --rm -i -v $(MAKEFILE_DIR):/go/src $(DEVBOX)

INAPI := docker compose -f deployments/docker-compose.yml exec api-gateway
INUSER := docker compose -f deployments/docker-compose.yml exec user-service

.PHONY: help setup-devbox vet fmt imports mod update lint generate build run test clean docker-up docker-down docker-logs

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

setup-devbox: ## Build devbox Docker image for development commands
	docker build --no-cache -t $(DEVBOX) -f $(MAKEFILE_DIR)/docker/devbox/Dockerfile $(MAKEFILE_DIR)

vet: ## Run go vet
	$(INAPI) go vet ./...

fmt: ## Format Go code
	$(INAPI) gofmt -d -s ./...

imports: ## Format imports
	$(INAPI) goimports -w ./...

mod: ## Run go mod tidy, verify and download
	$(INAPI) go mod tidy
	$(INAPI) go mod verify
	$(INAPI) go mod download

update: ## Update all Go dependencies
	$(INDEVBOX) go get -u ./...

lint: ## Run golangci-lint
	$(INAPI) golangci-lint run

generate-api: ## Generate API Gateway code
	$(INDEVBOX) goctl api go -api api/api/api.api -dir api --style gozero

generate-service: ## Generate service code (usage: make generate-service SERVICE=user)
	$(INDEVBOX) goctl rpc protoc service/$(SERVICE)/$(SERVICE).proto \
		--go_out=service/$(SERVICE) \
		--go-grpc_out=service/$(SERVICE) \
		--zrpc_out=service/$(SERVICE) \
		--style gozero

generate: generate-api generate-service ## Generate all code

build-api: ## Build API Gateway
	$(INAPI) go build -o bin/api ./api/main.go

build-user: ## Build User service
	$(INUSER) go build -o bin/user ./service/user/main.go

build: build-api build-user ## Build all services

run-api: ## Run API Gateway locally
	cd api && go run main.go

run-user: ## Run User service locally
	cd service/user && go run main.go

test: ## Run tests
	$(INAPI) go test ./...

clean: ## Clean build artifacts
	rm -rf bin/
	rm -rf api/logs/
	rm -rf service/*/logs/

docker-up: ## Start Docker Compose services
	cd deployments && docker-compose up -d

docker-down: ## Stop Docker Compose services
	cd deployments && docker-compose down

docker-logs: ## View Docker Compose logs
	cd deployments && docker-compose logs -f

