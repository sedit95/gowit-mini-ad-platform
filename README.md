# Gowit Mini Ad Platform

## Overview
This is a case-study sized mini ad platform that supports campaign CRUD operations and budget-based impression tracking. It leverages atomic database update logic to strictly prevent campaign overspending. The project includes a modern React frontend and a full Docker Compose runtime environment.

## Tech Stack

**Backend:**
- Go
- `chi` router
- `pgxpool`
- PostgreSQL

**Frontend:**
- Vite
- React
- TypeScript
- `react-router-dom`
- Native `fetch`

**Infrastructure:**
- Docker
- Docker Compose
- PostgreSQL container
- Nginx static frontend serving

**Testing & Validation:**
- Go integration tests
- Real PostgreSQL validation
- Backend concurrency integration test
- Frontend build validation
- Frontend runtime smoke test
- Docker Compose runtime smoke test

## Features

**Backend:**
- Create campaign
- List campaigns
- Get campaign detail
- Update allowed campaign fields
- Soft delete campaign
- Record impression
- Get campaign stats
- Auto-pause when budget reaches zero
- Atomic budget protection

**Frontend:**
- Campaign list page
- Campaign create form
- Campaign detail + stats page
- Record impression button
- Update section without budget field
- Delete action
- Stats polling

## Architecture
- The browser loads the frontend from `http://localhost:5173`.
- The browser calls the backend API at `http://localhost:8080`.
- The backend container connects to PostgreSQL through `postgres:5432` inside the Docker Compose network.
- `VITE_API_BASE_URL` is intentionally set to `http://localhost:8080` because these requests are made from the user's browser, not from within the Docker network.
- Do not use `http://backend:8080` in frontend browser code as it cannot resolve Docker's internal DNS.

## Backend API

**Endpoints:**
- `GET /health`
- `GET /campaigns`
- `POST /campaigns`
- `GET /campaigns/{id}`
- `PUT /campaigns/{id}`
- `DELETE /campaigns/{id}`
- `POST /impression/{id}`
- `GET /stats/{id}`

**Important Contract Notes:**
- `CreateCampaignRequest` includes the campaign budget.
- `UpdateCampaignRequest` does not update the budget.
- `POST /impression/{id}` may return `accepted=false` as a business outcome if the budget is exhausted.
- Soft-deleted campaigns behave as not found (HTTP 404).

## Environment Variables

**Backend:**
- `DATABASE_URL`
- `BACKEND_PORT`

**Frontend:**
- `VITE_API_BASE_URL`

**Docker Defaults:**
- Backend API: `http://localhost:8080`
- Frontend UI: `http://localhost:5173`
- Postgres host debug port: `localhost:5433`

## Local Development

**Backend PowerShell Example:**
```powershell
$env:DATABASE_URL="postgres://postgres:postgres@localhost:5432/gowit_ad_platform?sslmode=disable"
$env:BACKEND_PORT="8080"
cd backend
go run ./cmd/api
```

**Frontend PowerShell Example:**
```powershell
cd frontend
$env:VITE_API_BASE_URL="http://localhost:8080"
npm install
npm run dev
```

**Frontend Git Bash Example:**
```bash
cd frontend
export VITE_API_BASE_URL="http://localhost:8080"
npm install
npm run dev
```

## Docker Compose
To run the full stack via Docker Compose:
```bash
docker compose up --build
```

**URLs:**
- Frontend: `http://localhost:5173`
- Backend health: `http://localhost:8080/health`
- PostgreSQL host debug port: `localhost:5433`

**Clean DB Reset:**
To completely wipe the database and start fresh:
```bash
docker compose down -v
docker compose up --build
```

## Database Migration Strategy
- Docker Compose uses PostgreSQL init script mounting.
- The script `backend/migrations/001_create_campaigns_table.up.sql` is mounted into `/docker-entrypoint-initdb.d/`.
- The SQL runs only during the first empty postgres volume initialization.
- Use `docker compose down -v` to force clean re-initialization.
- *Note: This is an initialization strategy for local development, not a production-grade migration runner.*

## Testing and Validation

**Backend:**
- `go test ./...`
- `go test -v ./tests`
- CRUD integration test passed.
- Stats integration test passed.
- Soft delete integration test passed.
- Single impression lifecycle test passed.
- Concurrency integration test passed.

*Concurrency Evidence:*
- Campaign budget = 10
- 100 concurrent attempts
- Accepted true count = 10
- Final `remaining_budget` = 0
- Final `total_impressions` = 10
- Final `spent_budget` = 10
- Final status = paused

**Frontend:**
- `npm install` passed.
- `npm run build` passed.
- TypeScript compile passed.
- Vite build passed.
- Runtime smoke test passed against backend.

**Docker:**
- `docker compose config` passed.
- `docker compose up --build` passed.
- Postgres became healthy.
- Migration init mount worked.
- Backend `/health` returned OK.
- Frontend opened at `http://localhost:5173`.
- Docker runtime smoke test passed.

## Load Testing Scope
- No executable k6 script is included in the final delivered implementation.
- HTTP-level k6 load validation was intentionally left out of final scope.
- k6 validation requires separate load-test scenario design, calibration, thresholds, execution control, and result interpretation.
- The existing Go concurrency test successfully validates repository/service-level atomic budget protection.
- k6 would validate HTTP-level behavior under external concurrent traffic.
- `load-tests/README.md` documents the deferred k6 validation approach.

**Future k6 target:**
- `POST /impression/{id}`
- `budget` = 10
- Many concurrent HTTP requests
- Accepted impressions must not exceed 10
- `remaining_budget` must not go below 0
- Final status should become `paused`

## Known Limitations / Out of Scope
- No authentication
- No user management
- No admin panel
- No payment/billing
- No advanced analytics dashboard
- No executable k6 script
- No production deployment validation
- No multi-instance validation
- No CI/CD
- No Kubernetes
- No production-grade migration runner

## AI-Assisted Workflow Transparency
- The project was developed incrementally with AI assistance.
- Workflow and session records are kept in `AI_WORKFLOW.md` and `ai_session/`.
- The records rigorously separate implemented, validated, pending, and intentionally skipped items.
- Human review corrected AI-generated issues during the process, including compile and runtime issues.

## Project Structure
```text
backend/
frontend/
docs/
ai_session/
load-tests/
docker-compose.yml
AI_WORKFLOW.md
```
