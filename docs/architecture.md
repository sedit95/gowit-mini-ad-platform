# Architecture Plan

**Status:** *Pre-implementation planning phase. Implementation has not started yet.*

## System Overview
The Gowit Mini Ad Platform is designed as a simple, narrow-scoped web application. It explicitly rejects over-engineering patterns such as microservices, CQRS, event sourcing, Kubernetes, and cloud deployment for this case study.

## Components
- **Backend:** Go application structured with simple layering (`handler -> service -> repository`).
- **Frontend:** Vite + React + TypeScript, organized by `pages`, `components`, `api`, and `types`.
- **Database:** PostgreSQL.

## Validation & Orchestration
- **Docker Compose:** Planned to orchestrate the backend, frontend, and database locally.
- **Testing:** Basic Go tests and k6 load tests are planned to validate race condition safety.
