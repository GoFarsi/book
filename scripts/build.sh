#!/bin/bash

# Local build script for development
# This builds the desktop app with auto-generated version

set -e

# Get project root directory
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$PROJECT_ROOT"

echo "Building GoFarsi Book Desktop Application..."

# Create build directory if it doesn't exist
mkdir -p build

# Generate version
VERSION=$(./scripts/version.sh)
echo "Version: $VERSION"

# Build desktop application
echo "Building desktop application..."
cd app/desktop
go build -ldflags "-X main.VERSION=$VERSION" -o ../../build/book .

echo "Build completed successfully!"
echo "Binary location: $PROJECT_ROOT/build/book"
echo "Version: $VERSION"

# Optionally run the application
if [ "$1" = "--run" ]; then
    echo "Running the application..."
    ../../build/book
fi
