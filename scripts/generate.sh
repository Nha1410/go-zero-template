#!/bin/bash

# Script to generate code using goctl
# Usage: ./scripts/generate.sh [api|service|model]

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
        goctl api go -api api/api/api.api -dir api --style gozero
        echo "API Gateway code generated successfully!"
        ;;
    service)
        SERVICE_NAME=${2:-user}
        echo "Generating gRPC service code for: $SERVICE_NAME..."

        if [ ! -f "service/$SERVICE_NAME/$SERVICE_NAME.proto" ]; then
            echo "Error: proto file not found: service/$SERVICE_NAME/$SERVICE_NAME.proto"
            exit 1
        fi

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
            echo "Usage: ./scripts/generate.sh model 'host=localhost port=5432 user=postgres password=postgres dbname=gozero_template sslmode=disable'"
            exit 1
        fi
        goctl model pg datasource -url "$2" -table "*" -dir service/user/internal/model -cache

        echo "Model code generated successfully!"
        ;;
    *)
        echo "Usage: $0 [api|service|model]"
        echo ""
        echo "Commands:"
        echo "  api                    Generate API Gateway code from .api file"
        echo "  service <name>         Generate gRPC service code from .proto file"
        echo "  model <dsn>           Generate model code from PostgreSQL database"
        echo ""
        echo "Examples:"
        echo "  $0 api"
        echo "  $0 service user"
        echo "  $0 model 'host=localhost port=5432 user=postgres password=postgres dbname=gozero_template sslmode=disable'"
        exit 1
        ;;
esac

