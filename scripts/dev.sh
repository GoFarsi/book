#!/bin/bash

# Development utility script
# Usage: ./scripts/dev.sh [command]

set -e

PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$PROJECT_ROOT"

COMMAND=${1:-help}

case $COMMAND in
    "version")
        echo "Current version: $(./scripts/version.sh)"
        ;;
    "build")
        echo "Building application..."
        ./scripts/build.sh
        ;;
    "run")
        echo "Building and running application..."
        ./scripts/build.sh --run
        ;;
    "hugo")
        echo "Starting Hugo development server..."
        hugo server
        ;;
    "test")
        echo "Running tests..."
        cd app/desktop
        go test ./...
        ;;
    "clean")
        echo "Cleaning build artifacts..."
        rm -rf build/
        echo "Clean completed"
        ;;
    "deps")
        echo "Installing/updating dependencies..."
        cd app/desktop
        go mod tidy
        go mod download
        ;;
    "help"|*)
        echo "GoFarsi Book Development Script"
        echo ""
        echo "Usage: ./scripts/dev.sh [command]"
        echo ""
        echo "Commands:"
        echo "  version  - Show current version"
        echo "  build    - Build desktop application"
        echo "  run      - Build and run desktop application"
        echo "  hugo     - Start Hugo development server"
        echo "  test     - Run tests"
        echo "  clean    - Clean build artifacts"
        echo "  deps     - Install/update dependencies"
        echo "  help     - Show this help message"
        ;;
esac
