# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go SDK for QYWX (Enterprise WeChat) API. The project is in early development stage with basic authentication and API framework implemented.

## Development Commands

### Building and Running
```bash
# Build the binary
go build -o go-qywx.exe

# Run with help
go run main.go

# Start API server (listens on :8080)
go run main.go server

# Run application demo (shows token fetching)
go run main.go app
```

### Dependency Management
```bash
# Download dependencies
go mod download

# Tidy dependencies
go mod tidy

# Verify dependencies
go mod verify
```

### Testing
Currently no test files exist. When tests are added:
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## Architecture

### Package Structure
- `api/` - API constants and base response models
- `auth/` - Authentication and token management
- `cmd/` - CLI command implementations
  - `api/` - HTTP API server using Gin framework
  - `app/` - Application demo with hardcoded QYWX config

### Key Dependencies
- `github.com/spf13/cobra` - CLI framework
- `github.com/gin-gonic/gin` - Web framework for API server
- `github.com/carlmjohnson/requests` - HTTP client for API calls

### Current Implementation Notes
1. **Hardcoded Configuration**: Enterprise WeChat credentials are hardcoded in `cmd/app/index.go:27-38`
2. **Basic Authentication**: Token fetching is implemented in `auth/token.go`
3. **API Server**: Simple Gin server with `/ping` endpoint in `cmd/api/server.go`
4. **CLI Structure**: Uses Cobra with `server` and `app` subcommands

### Important Files
- `main.go` - Entry point, calls `cmd.Execute()`
- `cmd/root.go` - Root command definition and subcommand registration
- `auth/Model.go` - Configuration structs for enterprise setup
- `api/Model.go` - Base API response structure

## Development Notes

### Configuration Management
Currently uses hardcoded values. Future improvements should:
- Move configuration to external files (YAML/JSON)
- Support environment variables
- Add configuration validation

### API Implementation Pattern
The current pattern for API calls:
1. Define constants in `api/constants.go`
2. Create request/response structs in respective packages
3. Implement API calls using `github.com/carlmjohnson/requests`
4. Handle errors and response parsing

### Security Considerations
- **Hardcoded Secrets**: Current implementation has hardcoded CorpId and Secret
- **Token Management**: Basic token fetching implemented, needs token caching and refresh logic
- **No Encryption**: No encryption for configuration or sensitive data

### Extending the SDK
To add new QYWX API endpoints:
1. Add API path constants to `api/constants.go`
2. Create request/response models in appropriate package
3. Implement API call function following `auth/token.go` pattern
4. Add CLI command if needed in `cmd/` package

## Project State
- **Early Development**: Basic framework with minimal functionality
- **No Tests**: Test files need to be created
- **Limited Documentation**: Only basic README exists
- **Example Code**: Current implementation serves as example/demo

When working on this project, focus on:
1. Implementing missing QYWX API endpoints
2. Adding proper configuration management
3. Creating comprehensive tests
4. Improving documentation and examples