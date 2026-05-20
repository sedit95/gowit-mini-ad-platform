# AI Workflow

## Tools Used
- Google Antigravity (Current AI coding environment)

## Initial Prompting
The project started as an AI-native Gowit take-home case study. The first session began with strict rules: no initial file creation, no backend/frontend code generation, no dependency installation, no Docker/migration/k6 generation, and a clear directive to wait for the scaffolding prompt. The scaffolding prompt then instructed the creation of a placeholder repository structure without any application code.

## AI Decisions Accepted
- The initial placeholder folder structure was created.
- `AI_CONTEXT.md` was initialized and refined with project boundaries, mandatory stack, race condition risks, testing expectations, and human control rules.
- The `AI_WORKFLOW.md` skeleton, `.agents` structure, `docs` structure, and `ai_session` structure were accepted after review.

## AI Output Reviewed and Corrected
**Issue:** During the initial scaffolding, the AI agent incorrectly created a nested `gowit-mini-ad-platform` directory inside the existing repository root workspace.
**Detection & Correction:** The human developer detected this incorrect nested folder structure before implementation started. A correction was requested, and the AI agent moved the generated structure into the correct repository root and removed the nested folder.
**Verification:** The human developer verified the fix using PowerShell. `Get-ChildItem -Name` returned the expected root-level structure, and `Test-Path .\gowit-mini-ad-platform` successfully returned `False`.

## AI Suggestions Rejected
No implementation-level AI suggestions have been rejected yet because backend/frontend implementation has not started.

## Race Condition Review
The race condition risk (concurrent `/impression/:id` requests potentially causing negative budgets) has been identified early and recorded in `AI_CONTEXT.md`. The final implementation choice has not been created yet, though PostgreSQL-level atomic conditional update is the preferred direction to be validated in later phases.

## Human Corrections
- **Nested Folder Correction:** The human developer caught a repository structure issue where scaffolding was placed in an extra nested folder. The structure was flattened, and the fix was verified via PowerShell.
- **Service Import Correction:** During Campaign Service implementation, the AI-generated `service.go` file contained an unused `time` import. The human developer detected this through compile validation using `go test ./...`. The fix was limited to removing the unused `time` import from `backend/internal/campaign/service.go`. Business logic, validation logic, and RecordImpression logic were not changed. After the correction, `go test ./...` succeeded for the existing backend packages (output showed `[no test files]`, meaning successful compilation). Human review caught and fixed an AI-generated unused import compile issue before continuing.

## Testing Evidence
No tests have been executed yet. A specific Basic Go concurrency test and k6 load testing are planned as future validation.

## Final Ownership Notes
The human developer is strictly controlling the scope, actively reviewing AI output, and explicitly preventing unapproved implementation or feature expansion beyond the case study requirements.

## Agent Governance Refinement
During the scaffolding phase, the human developer directed the refinement of the agent governance files (AGENTS.md, .agents/workflows/, and .agents/skills/). These documents now explicitly define agent roles, strict workflow constraints, and skill expectations (especially regarding race condition safety). **No implementation code has been generated at this stage.**

## Master Prompt Preparation
A final master prompt was prepared before implementation started, defining the AI agent as a controlled, phase-based engineering assistant. The prompt strictly requires explicit human approval before any implementation, prevents scope expansion, and prioritizes race condition safety for `POST /impression/:id`. It explicitly rejects unsafe read-then-update logic and in-memory Go mutexes as final multi-instance-safe solutions, preferring instead a PostgreSQL-level atomic conditional update. It also mandates honest updates to `AI_WORKFLOW.md` and `ai_session` files. The prompt was delivered, the AI agent acknowledged the governance rules, and confirmed it would not implement anything without explicit approval. No backend/frontend implementation was generated during this phase.

## AI Context Refinement
`AI_CONTEXT.md` was refined before implementation to clearly align with the case study constraints. The frontend stack was clarified as Vite + React + TypeScript, explicitly excluding Next.js. Scope boundaries were expanded to clearly define forbidden technologies. The critical race condition rules were strengthened: unsafe read-then-update logic is explicitly forbidden, PostgreSQL-level atomic conditional updates are documented as the preferred direction, and in-memory Go mutexes are explicitly rejected as the final budget protection strategy because they do not protect multiple backend instances. Human approval rules and testing expectations were also tightened. No backend, frontend, Docker, migration, k6, or test implementation was generated, and no tests were executed during this phase.

## Documentation and Skill Refinement
Documentation and agent skill files were refined before implementation. The goal was to strengthen future agent guidance and reviewer-facing documentation. The work covered `README.md`, AI context, backend/frontend/race-condition/Docker/k6/documentation skills, and `docs/` planning files. The `initial_budget` / `remaining_budget` distinction was captured in `docs/database-design.md` for future implementation. No implementation or tests were produced during this phase.

## Backend Planning Phase
Backend planning was completed before implementation. The API contract was refined around a narrow 7-endpoint MVP:
- `GET /campaigns`
- `POST /campaigns`
- `GET /campaigns/:id`
- `PUT /campaigns/:id`
- `DELETE /campaigns/:id`
- `POST /impression/:id`
- `GET /stats/:id`

Campaign create accepts a general budget field, while persistence maps it to `initial_budget` and `remaining_budget`. `PUT /campaigns/:id` will not update budget in MVP. `DELETE /campaigns/:id` is soft delete using `deleted_at`. Campaign ID was planned as UUID. Budget uses INTEGER units because each successful impression deducts 1 unit. `spent_budget` is derived as `initial_budget` - `remaining_budget` and is not stored.

PostgreSQL atomic conditional update remains the preferred strategy for `POST /impression/:id`. Unsafe read-then-update logic remains forbidden. In-memory Go mutex remains rejected as the final multi-instance-safe solution.

`chi` was selected as the preferred router direction. `pgxpool` was selected as the preferred PostgreSQL connection direction. GORM/heavy ORM remains rejected.

Backend tests were planned, with `impression_concurrency_test.go` as the most critical test. The concurrency test target remains:
- initial budget = 10
- 100 concurrent impression attempts
- remaining_budget = 0
- impression_count = 10
- status = paused
- budget never negative

No backend code was generated. No migration SQL was created. No tests were created or executed. No Docker, frontend, or k6 files were generated.

## Backend Implementation Phase
Backend implementation was created incrementally. Files were generated in controlled phases. The backend currently compiles with `go test ./...`. The compile check does not mean functional tests exist. No Go test files have been created yet. No race condition/concurrency test has been executed yet. No k6 test has been executed yet. No Docker setup exists yet. Migration SQL exists but has not been executed. Manual API smoke testing is planned for the validation phase.

## Backend Validation Phase
- **Setup & Environment:** PostgreSQL was installed locally on Windows. A Turkish locale issue causing `initdb` failure was resolved by using the `C` locale. `psql` was executed via full path. `gowit_ad_platform` and `gowit_ad_platform_test` databases were created.
- **Migration & Runtime Validation:** Migrations applied successfully with pgcrypto; schemas and constraints matched plans perfectly. Backend started successfully with `DATABASE_URL`, config validation worked properly, and `GET /health` returned `200 OK`.
- **Manual Smoke Test:** All 7 endpoints (`POST /campaigns`, `GET /campaigns`, `GET /campaigns/:id`, `GET /stats/:id`, `POST /impression/:id`, `DELETE /campaigns/:id`) were validated. Sequential budget exhaustion succeeded (status updated to `paused`), extra impressions returned `budget_exhausted` cleanly, and soft deletes accurately mapped to 404s.
- **Integration & Concurrency Testing:** `go test ./...` and `go test -v ./tests` passed against `TEST_DATABASE_URL`. The critical concurrency test proved the PostgreSQL atomic conditional update works: 100 concurrent attempts on a budget of 10 resulted in exactly 10 accepted impressions, 90 rejected impressions, final remaining budget of 0, and campaign paused without negative budgets.
- **Pending Validation:** This is strictly backend service/repository-level validation. HTTP-level k6 load testing, Docker Compose validation, multi-instance production validation, and frontend validation are still pending.

## Frontend Planning Phase
Frontend planning has been completed. The frontend stack will be Vite + React + TypeScript, explicitly rejecting Next.js. The application will feature a Campaign List Page, a Campaign Create Page, and a Campaign Detail + Stats Page with a 3000ms polling strategy for live stats. No WebSockets or optimistic budget decrements will be used; the backend remains the source of truth. TypeScript models and API mappings have been established. Allowed dependencies include React Router, while Redux, Axios, Formik, and Tailwind have been rejected to keep the scope narrow. Frontend implementation, HTTP-level k6 validation, and Docker validation are explicitly pending and have not started.

## Frontend Implementation Phase
Frontend implementation was completed incrementally following the plan (Campaign List, Create, Detail+Stats pages, and 3000ms stats polling).

During the runtime smoke test (Frontend at `http://localhost:5173`, Backend at `http://localhost:8080`), browser CORS enforcement blocked frontend API requests, failing `POST /campaigns` during preflight.

**CORS Fix performed:**
- Added a minimal local CORS middleware in `backend/internal/http/cors.go`.
- Registered the middleware in `backend/cmd/api/main.go` before routes.
- Allowed origin: `http://localhost:5173`, methods: GET, POST, PUT, DELETE, OPTIONS, headers: Content-Type.
- `OPTIONS` preflight returns HTTP 204 No Content.
- No external CORS dependency was added, and no campaign business logic or frontend code was changed.

**Validation after fix:**
- `go test ./...` passed on the backend.
- Both backend and frontend ran successfully.
- List page, Campaign Create, Detail/Stats page, Record Impression, Update Campaign (without budget fields), and Delete Campaign actions all worked properly.
- Stats polling successfully executed without issues.
- Browser console remained clear of critical errors.

**Important Boundaries:**
- This was strictly frontend runtime smoke validation, not k6 load testing.
- HTTP-level k6 validation is still pending.
- Docker Compose validation is still pending.
- The CORS fix was limited specifically to the local development frontend origin. Full system production validation is not complete.

## Docker Compose Phase
The Docker Compose phase successfully containerized the application for local runtime validation.

**Docker files created:** `backend/Dockerfile`, `backend/.dockerignore`, `frontend/Dockerfile`, `frontend/.dockerignore`, `frontend/nginx.conf`, and `docker-compose.yml`.

**Architecture:** 
- `postgres`: PostgreSQL 16 Alpine image. Maps port `5433:5432`.
- `backend`: Multi-stage Go build (`golang:1.22-alpine` builder to `alpine:3.20` runtime). Connects to PostgreSQL using `postgres:5432`. Maps port `8080:8080`.
- `frontend`: Multi-stage Vite build (`node:20-alpine` builder to `nginx:1.27-alpine` runtime). Nginx handles SPA fallback. Maps port `5173:80`.

**Environment & Network:**
- Backend `DATABASE_URL` references the Docker service name (`postgres:5432`).
- Frontend `VITE_API_BASE_URL` is baked in at build-time using the default `http://localhost:8080`. The browser resolves this correctly on the host machine.
- Existing local backend CORS configuration correctly accommodates the browser's origin (`http://localhost:5173`).

**Migration Strategy:**
- Initial migration uses the standard `postgres` image entrypoint by mapping the `001_create_campaigns_table.up.sql` file into `/docker-entrypoint-initdb.d/`.
- No separate migration service or application-level startup migration was used. Resetting the DB requires `docker compose down -v`.

**Validation Results:**
- `docker compose config` and `docker compose up --build` succeeded perfectly.
- Database initialized automatically, and the backend service connected without issue (`/health` OK).
- The frontend was accessible at `http://localhost:5173`.
- Smoke testing confirmed List, Create, Detail, Impression, Update, Polling, and Delete features functioned exactly as they did natively. No CORS errors, connection failures, or container crashes occurred.

**Honesty Boundaries:**
- This was local Docker Compose validation, not a scalable production deployment.
- HTTP-level k6 load testing is still pending.
- Multi-instance testing is still pending.

## Final Project Status

### Completed Implementation
- Backend implementation completed.
- Backend PostgreSQL migration exists and was applied locally and through Docker Compose init mount.
- Backend validation completed with Go integration tests.
- Frontend planning completed.
- Frontend implementation completed with Vite + React + TypeScript.
- Docker Compose implementation completed with postgres, backend, and frontend services.
- Final README was updated.

### Backend Validation Evidence
- `go test ./...` passed.
- `go test -v ./tests` passed.
- CRUD integration test passed.
- Stats integration test passed.
- Soft delete integration test passed.
- Single impression lifecycle test passed.
- Concurrency integration test passed.
- Concurrency scenario:
  - budget=10
  - 100 concurrent attempts
  - accepted=true count=10
  - final remaining_budget=0
  - final total_impressions=10
  - final spent_budget=10
  - final status=paused

### Frontend Validation Evidence
- `npm install` passed.
- `npm run build` passed.
- TypeScript compile passed.
- Vite build passed.
- Frontend runtime smoke test passed against the backend.
- List/create/detail/stats/impression/polling/update/delete flows were manually verified.
- Budget update field was absent from the frontend update section.

### Docker Validation Evidence
- `backend/Dockerfile` was created and built.
- `frontend/Dockerfile` was created and built.
- `docker-compose.yml` was created.
- `docker compose config` passed.
- `docker compose up --build` passed.
- postgres/backend/frontend containers ran successfully.
- PostgreSQL init mount migration worked.
- `campaigns` table was created.
- backend `/health` returned OK.
- frontend opened at `http://localhost:5173`.
- Docker runtime smoke test passed.
- No Docker fix was required after smoke testing.

### CORS & Runtime Integration Fix
- Frontend local origin `http://localhost:5173` initially triggered CORS errors against backend `http://localhost:8080`.
- A minimal backend CORS middleware was added earlier in the project.
- Allowed origin remains `http://localhost:5173`.
- This fix allowed local and Docker frontend runtime integration.
- No wildcard CORS policy was introduced.

### Docker & Migration Boundaries
- Docker Compose uses PostgreSQL init mount for migration.
- No separate migration service was used.
- Backend startup does not run migrations.
- PostgreSQL init scripts run only during first empty volume initialization.
- Use `docker compose down -v` for a clean DB reset.
- This is not a production-grade migration runner.

### k6 / Load Testing Scope
- No executable k6 script is included in the final delivered implementation.
- HTTP-level k6 load validation was intentionally left out of final delivered scope.
- `load-tests/README.md` documents the deferred k6 validation approach.
- Existing Go concurrency tests validate repository/service-level atomic budget protection.
- k6 would be future work for HTTP-level external concurrent traffic validation.
- Do not claim k6 was run.
- Do not claim load testing passed.

### Known Out-of-Scope Items
- No authentication.
- No user management.
- No admin panel.
- No payment/billing.
- No advanced analytics dashboard.
- No production deployment validation.
- No multi-instance validation.
- No CI/CD.
- No Kubernetes.
- No production-grade migration runner.

### Final Honesty Boundary
- The project is validated as a local case-study implementation with backend tests, frontend build/runtime smoke tests, and Docker Compose runtime smoke tests.
- Do not claim full production readiness.
- Do not claim full production performance validation.
