#!/bin/bash

# Generate version based on datetime and git info
# Format: YYYY.MM.DD-HHMMSS-{short_commit}

# Get current date and time
DATE=$(date +%Y.%m.%d)
TIME=$(date +%H%M%S)

# Get git short commit hash (fallback to 'unknown' if not in git repo)
if git rev-parse --short HEAD >/dev/null 2>&1; then
    COMMIT=$(git rev-parse --short HEAD)
else
    COMMIT="unknown"
fi

# Check if we're on a tagged commit
if git describe --tags --exact-match >/dev/null 2>&1; then
    # If on a tagged commit, use the tag as version
    VERSION=$(git describe --tags --exact-match)
else
    # Otherwise, use datetime-based version
    VERSION="${DATE}-${TIME}-${COMMIT}"
fi

echo "${VERSION}"
