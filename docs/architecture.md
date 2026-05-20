# Architecture Plan

**Status:** *Pre-implementation planning phase. No files have been created yet.*

## System Overview
The Gowit Mini Ad Platform utilizes a simple layered Go backend (`handler -> service -> repository`) paired with a Vite + React + TypeScript frontend, backed by PostgreSQL. The architecture strictly rejects CQRS, event sourcing, microservices, and any unnecessary abstraction.

## Planned Backend Structure
```text
backend/
  cmd/api/main.go
  internal/config/config.go
  internal/db/postgres.go
  internal/campaign/model.go
  internal/campaign/dto.go
  internal/campaign/repository.go
  internal/campaign/service.go
  internal/campaign/handler.go
  internal/campaign/routes.go
  internal/http/response.go
  internal/errors/errors.go
  migrations/
  tests/
  go.mod
```

## Technology Choices & Direction
- **Router:** `chi` is preferred for simple routing.
- **Database Library:** `pgxpool` is preferred for robust PostgreSQL connection pooling and context-aware queries.
- **Dependencies:** Keep dependencies minimal. Avoid heavy ORMs like GORM and dependency injection frameworks.

## Layer Responsibilities
- **Handler:** Decodes HTTP requests and writes HTTP responses.
- **Service:** Owns business decisions and validation outcomes.
- **Repository:** Owns PostgreSQL queries. Atomic impression SQL belongs here. Raw DB errors must not leak to clients.

## Testing Plan
- `backend/tests/impression_concurrency_test.go` is the most critical planned test.
- Additional tests may include CRUD, soft delete, stats calculation, and single impression behavior.
