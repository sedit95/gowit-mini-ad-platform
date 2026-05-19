# AI Context

## 1. Project Goal
- AI-native Mini Campaign Management Platform for Gowit case study.
- The system allows campaign management and impression tracking.

## 2. Mandatory Technology Stack
- Backend: Go
- Frontend: React + TypeScript
- Database: PostgreSQL
- Docker Compose planned
- k6 load testing planned
- Basic Go tests planned

## 3. Core Business Requirements
- Campaign CRUD
- Campaign fields: title, budget, currency, start date, end date, status
- Status values: active, paused, completed
- POST /impression/:id deducts 1 unit from campaign budget
- GET /stats/:id returns total impressions, spent budget, remaining budget
- Soft delete is required

## 4. Critical Race Condition Requirement
- `/impression/:id` may receive many concurrent requests.
- Budget must never go negative under any condition.
- When budget reaches zero, campaign must auto-pause.
- Final implementation should prefer database-level concurrency safety.
- PostgreSQL atomic conditional update is the preferred direction, but implementation is not created yet.

## 5. Scope Boundaries
The following are out of scope:
- Authentication
- Payments
- Multi-tenancy
- Advanced analytics
- Complex dashboards
- Message queues
- Unnecessary microservices

## 6. Architecture Principles
- Keep the project narrow and clean.
- Prefer simple, explainable architecture.
- Avoid overengineering.
- Keep backend, frontend, load tests, documentation, and AI workflow clearly separated.
- Decisions must be documented.

## 7. Documentation Expectations
- README.md should stay short and setup-focused.
- Detailed technical reasoning should go under docs/.
- AI_WORKFLOW.md should be updated throughout the project.
- ai_session/ should include readable AI session summaries.

## 8. Testing Expectations
- Basic Go tests are planned.
- A specific concurrency test is planned: budget=10, 100 concurrent impression attempts, remaining_budget=0, impression_count=10, status=paused, and budget never negative.
- k6 load testing is planned to demonstrate race condition safety.

## 9. Human Control Rules
- The human developer controls scope and decisions.
- AI suggestions must be reviewed, questioned, and corrected when necessary.
- AI must not implement unapproved features.
- AI must not expand scope beyond the case study requirements.
