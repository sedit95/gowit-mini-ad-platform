# Docker Compose Phase

## Docker Files Created
- `backend/Dockerfile`
- `backend/.dockerignore`
- `frontend/Dockerfile`
- `frontend/.dockerignore`
- `frontend/nginx.conf`
- `docker-compose.yml`

## Service Architecture
### Backend Docker Details
- Multi-stage Go backend build.
- **Builder:** `golang:1.22-alpine` (compiles `./cmd/api` into the `api` binary).
- **Runtime:** `alpine:3.20` (runs `./api` and exposes `8080`).
- `DATABASE_URL` and `BACKEND_PORT` are not hardcoded in the Dockerfile.

### Frontend Docker Details
- Multi-stage Vite frontend build.
- **Builder:** `node:20-alpine`.
- **Runtime:** `nginx:1.27-alpine`.
- `VITE_API_BASE_URL` is configured as a build-time argument. The default browser-reachable API URL remains `http://localhost:8080`.
- Nginx serves static `dist` output with an SPA fallback (`try_files $uri $uri/ /index.html;`) in `nginx.conf`. No API proxying is performed through nginx.

### Docker Compose Services
- `postgres` (postgres:16-alpine)
- `backend`
- `frontend`

## Port, Environment, and Network Decisions
- **Ports:** `postgres` mapped to `5433:5432`, `backend` mapped to `8080:8080`, and `frontend` mapped to `5173:80`.
- **Backend Network:** `DATABASE_URL` targets the internal compose network (`postgres:5432`), not localhost.
- **Frontend Network:** `VITE_API_BASE_URL` targets `http://localhost:8080`, ensuring the browser on the host machine correctly routes API calls to the backend port.
- The existing backend CORS origin is entirely compatible, and Docker Compose's default network handles internal communication.

## Migration Strategy & Limitations
- No separate migration service was used, and migrations are not run during backend startup.
- Instead, the PostgreSQL init mount strategy is used: `./backend/migrations/001_create_campaigns_table.up.sql` is mounted to `/docker-entrypoint-initdb.d/001_create_campaigns_table.up.sql`.
- **Limitation:** The migration only runs during the very first initialization of an empty postgres volume. A full DB reset requires `docker compose down -v`. This is not a production-grade migration runner.

## Compose Build & Validation
- `docker compose config` succeeded.
- `docker compose down -v` ensured a clean state.
- `docker compose up --build` succeeded across all containers.
- The `postgres` container became healthy and the migration script successfully created the `campaigns` table.
- The `backend` container connected immediately; `/health` returned `200 OK`.
- The `frontend` container ran successfully and was accessible at `http://localhost:5173`.

## Runtime Smoke Test Result
- List page, Create campaign, Detail/Stats page, Record impression, Polling, Update section (without budget field), and Delete all worked perfectly.
- The browser console was completely clear of critical errors, confirming no CORS issues.

## Issue Review
- **No Docker fix was required** following the runtime smoke test.
- No port conflicts, database connection failures, missing tables, CORS blocks, or container crashes were observed.

## Remaining Pending Validations
- This was an isolated Docker Compose local runtime validation.
- It was not a k6 load test. HTTP-level k6 validation is still pending.
- Production deployment validation and multi-instance validation remain pending.
