#!/bin/bash

# Quick script to generate API Gateway code
# Usage: ./scripts/generate-api.sh

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
"$SCRIPT_DIR/generate.sh" api

