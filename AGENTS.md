# AGENTS.md - MIAFIN Project Guidelines

## Project Overview

**MIAFIN** is a full-stack security/SaaS application with:
- **Frontend**: SvelteKit 5 + TypeScript (Bun package manager)
- **Backend**: Go 1.25+ with standard project layout
- **Repository**: https://github.com/LuneLoops/MIAFIN

## Build & Development Commands

### Frontend (SvelteKit)

```bash
cd frontend

# Development
bun run dev              # Start dev server
bun run check            # Type check Svelte
bun run check:watch      # Type check in watch mode

# Build
bun run build            # Production build
bun run preview          # Preview production build

# Testing (when configured)
bun test                 # Run all tests
bun test --run SingleTest # Run single test file
bun test --watch         # Run tests in watch mode
```

### Backend (Go)

```bash
cd backend

# Development
go run ./cmd/api         # Run API server
go run ./cmd/worker      # Run background worker
go run .                 # Run main package

# Testing
go test ./...            # Run all tests
go test -v ./...         # Run with verbose output
go test -run TestName ./pkg/models  # Run single test
go test -race ./...      # Run with race detector
go test -cover ./...     # Run with coverage

# Build
go build -o bin/api ./cmd/api
go build -o bin/worker ./cmd/worker

# Dependencies
go mod tidy              # Clean up dependencies
go mod download          # Download dependencies
```

### Lint & Format

```bash
# Go
gofmt -w .               # Format Go code
go vet ./...             # Static analysis
golangci-lint run        # Run all linters (install first)

# Frontend (when configured)
bun run lint             # ESLint
bun run format           # Prettier
```

## Code Style Guidelines

### Go (Backend)

**Imports**: Group imports: stdlib, external, internal
```go
import (
    "context"
    "time"
    
    "github.com/gin-gonic/gin"
    
    "github.com/LuneLoops/MIAFIN/internal/auth"
)
```

**Naming**:
- Exported: PascalCase (`UserService`)
- Unexported: camelCase (`userService`)
- Interfaces: noun ending in -er (`Reader`, `Handler`)
- Packages: short, lowercase, no underscores

**Error Handling**:
- Always check errors explicitly
- Wrap errors with context: `fmt.Errorf("failed to fetch user: %w", err)`
- Use custom error types for domain errors
- Never ignore errors with `_` without comment

**Types**:
- Use structs with field tags for JSON/DB
- Prefer composition over inheritance
- Use interfaces for dependencies (testing)
- Leverage generics when appropriate

**Functions**:
- Accept interfaces, return concrete types
- First parameter should be `ctx context.Context`
- Keep functions small and focused
- Return early to reduce nesting

**Concurrency**:
- Always use `context` for cancellation
- Never leak goroutines - use WaitGroup
- Protect shared state with sync primitives
- Run race detector in tests

### Svelte/TypeScript (Frontend)

**Components**:
- Use Svelte 5 runes: `$state`, `$derived`, `$effect`
- Props: Use runes with TypeScript interfaces
- Events: Use callbacks as props
- Slots: Use named slots for complex layouts

**Naming**:
- Components: PascalCase (`UserCard.svelte`)
- Stores: camelCase with Store suffix (`userStore`)
- Files: PascalCase for components, camelCase for utilities

**TypeScript**:
- Enable strict mode
- Define interfaces for all props
- Use explicit return types on functions
- Avoid `any` - use `unknown` with type guards

**Styling**:
- Use scoped styles in components
- CSS variables for theming
- Tailwind classes when utility classes needed
- BEM or utility-first naming

**Error Handling**:
- Use try/catch for async operations
- Display user-friendly error messages
- Log errors to monitoring service
- Handle loading states explicitly

## Project Structure

### Backend Layout
```
backend/
├── cmd/
│   ├── api/           # API server entry
│   └── worker/        # Background jobs
├── internal/          # Private code
│   ├── auth/          # Authentication
│   ├── tenant/        # Multi-tenancy
│   ├── audit/         # Security logs
│   ├── rbac/          # Permissions
│   └── database/      # DB connections
├── pkg/               # Public libraries
│   ├── models/        # Domain models
│   ├── errors/        # Error types
│   └── crypto/        # Encryption utils
└── api/swagger/       # API docs
```

### Frontend Layout
```
frontend/src/
├── lib/
│   ├── components/    # Reusable UI
│   ├── stores/        # Svelte stores
│   └── api/           # API client
└── routes/            # SvelteKit routes
```

## Testing Standards

### Go
- Table-driven tests
- Mock external dependencies
- 80%+ coverage for critical paths
- Use testdata/ for fixtures
- Run with `-race` flag in CI

### Frontend
- Unit tests for utilities/stores
- Component tests with testing-library
- E2E tests for critical flows
- Mock API calls in tests

## Git Workflow

- Branch: `main` for production
- Feature branches: `feature/description`
- Commit messages: Conventional commits
- Pull requests required for main
- No secrets in commits

## Security Considerations

- Never log sensitive data
- Validate all inputs
- Use parameterized queries
- Implement proper CORS
- Store secrets in environment variables
- Use HTTPS in production

## AI Agent Instructions

When modifying code:
1. Follow existing patterns in the codebase
2. Run tests before suggesting changes
3. Use type-safe code (avoid `any` in TS)
4. Add error handling for all fallible operations
5. Keep functions small and testable
6. Document exported functions/types
7. Respect the `internal/` package boundary in Go
8. Use Svelte 5 runes correctly (check with svelte-autofixer)

## Available Skills

This project has configured skills:
- `golang-pro`: Go development, concurrency, microservices
- `svelte-code-writer`: Svelte 5 documentation and code analysis
- `postgresql-table-design`: PostgreSQL schema design

Load appropriate skill before working on related code.
