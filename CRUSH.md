# Ollama Development Guide for Agentic Coding Agents

## Build Commands
- Build and run: `go run . serve`
- Clean CGO cache: `go clean -cache`
- Run tests: `go test ./...`
- Enable synctest for development: `GOEXPERIMENT=synctest go test ./...`
- Lint: `golangci-lint run`

## Test Commands
- Run all tests: `go test ./...`
- Run specific test: `go test -run TestName ./path/to/package`
- Run with synctest: `GOEXPERIMENT=synctest go test ./...`
- Run benchmarks: `go test -bench=. ./...`

## Code Style Guidelines

### Go Code
- Use standard Go formatting (gofmt/gofumpt)
- Follow idiomatic Go patterns
- Use descriptive variable and function names
- Prefer explicit error handling with early returns
- Use context.Context for cancellation and timeouts
- Prefer interfaces for dependency injection

### Imports
- Group standard library imports separately from third-party imports
- Use descriptive import aliases when needed
- Avoid unused imports

### Types and Naming
- Use camelCase for variables and functions
- Use PascalCase for exported identifiers
- Prefer descriptive names over abbreviations
- Use context-specific types rather than generic interfaces where possible

### Error Handling
- Always handle errors explicitly with `if err != nil` checks
- Return early on errors to reduce nesting
- Use error wrapping with `fmt.Errorf("description: %w", err)`
- Use `errors.Is()` and `errors.As()` for error checking
- Create custom error types with `errors.New()` for specific conditions

## Project Structure
- CLI code: `cmd/`
- Core server logic: `server/`
- API definitions: `api/`
- Models and conversion: `model/`, `convert/`
- Native libraries: `llama/`, `ml/`
- UI code: `app/ui/` (TypeScript/React)
- Format handling: `format/`

## Linting and Formatting
- Enabled linters: asasalint, bidichk, bodyclose, containedctx, gocheckcompilerdirectives, gofmt, gofumpt, gosimple, govet, ineffassign, intrange, makezero, misspell, nilerr, nolintlint, nosprintfhostport, staticcheck, unconvert, usetesting, wastedassign, whitespace
- Disabled linters: usestdlibvars, errcheck
- Severity rules: gofmt, goimports, intrange are treated as info level

## Backend Libraries
- Native code is in `llama/llama.cpp` and compiled with CGO
- Acceleration libraries are loaded from `./lib/ollama` (Windows), `../lib/ollama` (Linux), `.` (macOS), `build/lib/ollama` (development)

## TOON Format Support
- TOON format is supported by specifying `"format": "toon"` in API requests
- The implementation is in `format/toon.go` which provides encoding/decoding functions
- TOON format responses are converted from JSON to TOON format before being sent to clients
- TOON format is handled in both Generate and Chat endpoints
- Grammar for TOON format is defined in `llm/server.go`

## Adding New Formats
1. Update format validation in `llm/server.go`
2. Add format-specific handling in `server/routes.go`
3. Create format-specific encoding/decoding functions in `format/` package
4. Update tests to include the new format
5. Update documentation and CRUSH.md