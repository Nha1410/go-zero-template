#!/bin/bash

# Script to generate code using goctl
# Usage: ./scripts/generate.sh [api|rpc|model|docker|kube|migrate|template|quickstart]

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

cd "$PROJECT_ROOT"

# Check if goctl is installed
if ! command -v goctl &> /dev/null; then
    echo "Error: goctl is not installed"
    echo "Install it with: go install github.com/zeromicro/go-zero/tools/goctl@latest"
    exit 1
fi

case "$1" in
    api)
        echo "Generating API Gateway code..."
        if [ -z "$2" ]; then
            # Default: generate from api/api/api.api
            goctl api go -api api/api/api.api -dir api --style gozero
        else
            # Custom .api file
            goctl api go -api "$2" -dir api --style gozero
        fi
        echo "API Gateway code generated successfully!"
        ;;
    rpc|service)
        SERVICE_NAME=${2:-user}
        echo "Generating gRPC service code for: $SERVICE_NAME..."

        if [ ! -f "service/$SERVICE_NAME/$SERVICE_NAME.proto" ]; then
            echo "Error: proto file not found: service/$SERVICE_NAME/$SERVICE_NAME.proto"
            exit 1
        fi

        # Use goctl rpc command
        goctl rpc protoc service/$SERVICE_NAME/$SERVICE_NAME.proto \
            --go_out=service/$SERVICE_NAME \
            --go-grpc_out=service/$SERVICE_NAME \
            --zrpc_out=service/$SERVICE_NAME \
            --style gozero

        echo "Service code generated successfully!"
        ;;
    model)
        echo "Generating model code from PostgreSQL database..."

        if [ -z "$2" ]; then
            echo "Error: PostgreSQL connection string required"
            echo "Usage: ./scripts/generate.sh model 'host=localhost port=5432 user=postgres password=postgres dbname=gozero_template sslmode=disable' [table] [output_dir]"
            exit 1
        fi

        TABLE=${3:-"*"}
        OUTPUT_DIR=${4:-"service/user/internal/model"}

        goctl model pg datasource -url "$2" -table "$TABLE" -dir "$OUTPUT_DIR" -cache

        echo "Model code generated successfully!"
        ;;
    docker)
        echo "Generating Dockerfile..."
        if [ -z "$2" ]; then
            echo "Error: Service name required"
            echo "Usage: ./scripts/generate.sh docker [api|user|service-name]"
            exit 1
        fi
        SERVICE_NAME=$2
        if [ "$SERVICE_NAME" = "api" ]; then
            goctl docker -go api/main.go
        else
            goctl docker -go service/$SERVICE_NAME/main.go
        fi
        echo "Dockerfile generated successfully!"
        ;;
    kube)
        echo "Generating Kubernetes files..."
        if [ -z "$2" ]; then
            echo "Error: Service name required"
            echo "Usage: ./scripts/generate.sh kube [api|user|service-name]"
            exit 1
        fi
        SERVICE_NAME=$2
        if [ "$SERVICE_NAME" = "api" ]; then
            goctl kube deploy -name api-gateway -namespace default -image api-gateway:latest -port 8888 -o k8s/
        else
            goctl kube deploy -name $SERVICE_NAME-service -namespace default -image $SERVICE_NAME-service:latest -port 9000 -o k8s/
        fi
        echo "Kubernetes files generated successfully!"
        ;;
    migrate)
        echo "Migrating from tal-tech to zeromicro..."
        goctl migrate --help
        echo ""
        echo "For more information, run: goctl migrate --help"
        ;;
    template)
        echo "Template operations..."
        if [ -z "$2" ]; then
            echo "Usage: ./scripts/generate.sh template [install|clean|validate]"
            exit 1
        fi
        goctl template $2 ${@:3}
        ;;
    quickstart)
        echo "Quick start a new project..."
        if [ -z "$2" ]; then
            echo "Usage: ./scripts/generate.sh quickstart [project-name]"
            exit 1
        fi
        goctl quickstart -o "$2"
        ;;
    *)
        echo "Usage: $0 [api|rpc|model|docker|kube|migrate|template|quickstart]"
        echo ""
        echo "Commands:"
        echo "  api [file]            Generate API Gateway code from .api file"
        echo "  rpc <name>            Generate gRPC service code from .proto file"
        echo "  model <dsn> [table] [dir]  Generate model code from PostgreSQL database"
        echo "  docker <name>         Generate Dockerfile for service"
        echo "  kube <name>           Generate Kubernetes deployment files"
        echo "  migrate               Migrate from tal-tech to zeromicro"
        echo "  template <cmd>        Template operations (install/clean/validate)"
        echo "  quickstart <name>     Quick start a new project"
        echo ""
        echo "Examples:"
        echo "  $0 api"
        echo "  $0 api api/api/api.api"
        echo "  $0 rpc user"
        echo "  $0 model 'host=localhost port=5432 user=postgres password=postgres dbname=gozero_template sslmode=disable'"
        echo "  $0 docker api"
        echo "  $0 kube user"
        exit 1
        ;;
esac

