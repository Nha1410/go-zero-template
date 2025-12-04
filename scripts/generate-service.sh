#!/bin/bash

# Quick script to generate gRPC service code
# Usage: ./scripts/generate-service.sh [service-name]

set -e

SERVICE_NAME=${1:-user}
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
"$SCRIPT_DIR/generate.sh" rpc "$SERVICE_NAME"

