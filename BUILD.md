# Build and Version Management

This project uses an automated datetime-based versioning system that eliminates the need for manual version updates.

## Version Format

The version follows this format:
- **For releases**: `YYYY.MM.DD-HHMMSS-{git_short_hash}`
- **For tagged releases**: Uses the git tag as version
- **For development**: `dev` (when no git info available)

Example: `2025.07.16-094725-c24088e58`

## Building

### Local Development

Use the development script for common tasks:

```bash
# Show current version
./scripts/dev.sh version

# Build the desktop application
./scripts/dev.sh build

# Build and run the application
./scripts/dev.sh run

# Start Hugo development server
./scripts/dev.sh hugo

# Run tests
./scripts/dev.sh test

# Clean build artifacts
./scripts/dev.sh clean

# Install/update dependencies
./scripts/dev.sh deps
```

### Manual Build

You can also build manually:

```bash
# Build with auto-generated version
make build-desktop

# Or use the build script directly
./scripts/build.sh

# Build and run immediately
./scripts/build.sh --run
```

### CI/CD Build

The GitHub Actions workflow automatically:
1. Generates a datetime-based version
2. Creates a git tag for the version
3. Builds and releases the application with GoReleaser
4. Publishes to GitHub Releases

## Version Injection

The version is injected at build time using Go's `-ldflags` flag:

```bash
go build -ldflags "-X main.VERSION=$(./scripts/version.sh)" -o book .
```

This means:
- ✅ No need to manually update version in source code
- ✅ No version-related commits cluttering git history
- ✅ Each build gets a unique, traceable version
- ✅ Supports both development and production builds
- ✅ Compatible with existing GoReleaser workflow

## Scripts

- `scripts/version.sh` - Generates the version string
- `scripts/build.sh` - Builds the desktop application
- `scripts/dev.sh` - Development utility script with common commands

## Hugo Integration

The Hugo static site generation remains unchanged and continues to build and deploy to GitHub Pages automatically on pushes to the main branch.

## Migration from Manual Versioning

The old manual versioning system with `const VERSION = "6.2.0"` has been replaced with:
- `var VERSION = "dev"` - Default development version
- Build-time injection sets the actual version
- No source code changes needed for version updates
