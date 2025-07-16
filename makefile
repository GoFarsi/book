.PHONY: test wk build build-desktop version clean

# Get the version from the script
VERSION := $(shell ./scripts/version.sh)

test:
	hugo server

wk:
	git submodule update --remote --recursive
	cd app/worker && wrangler publish

# Build the desktop application with version injection
build-desktop:
	@echo "Building desktop app with version: $(VERSION)"
	cd app/desktop && go build -ldflags "-X main.VERSION=$(VERSION)" -o ../../build/book .

# General build target
build: build-desktop

# Show current version
version:
	@echo $(VERSION)

# Clean build artifacts
clean:
	rm -rf build/
