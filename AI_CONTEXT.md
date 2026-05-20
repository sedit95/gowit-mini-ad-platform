# AI Context

## 1. Project Goal
- AI-native Mini Campaign Management Platform for Gowit case study.
- The system allows campaign management and impression tracking.

## 2. Mandatory Technology Stack
- Backend: Go
- Frontend: Vite + React + TypeScript (Next.js will not be used for this case study)
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
- Campaigns must not be permanently deleted. Use soft delete semantics, such as `deleted_at`.

## 4. Critical Race Condition Requirement
- `/impression/:id` may receive many concurrent requests.
- Budget must never go negative under any condition.
- When budget reaches zero, campaign must auto-pause.
- Unsafe read-then-update logic must be avoided.
- Budget decrement, impression count increment, and auto-pause status transition should be handled atomically.
- The preferred direction is PostgreSQL-level atomic conditional update.
- PostgreSQL transaction with row-level locking may be considered only if clearly justified.
- In-memory Go mutex must not be used as the final budget protection strategy because it does not protect multiple backend instances.
- Implementation has not started yet.

## 5. Scope Boundaries
The following are out of scope:
- Authentication
- Authorization / RBAC
- User management
- Payments
- Multi-tenancy
- Advanced analytics
- Complex dashboards
- Message queues
- Event sourcing
- CQRS
- Unnecessary microservices
- Kubernetes
- Cloud deployment
- CI/CD pipeline
- Redis unless explicitly approved
- WebSocket unless explicitly approved
- Notification systems
- Email/SMS integrations
- Admin panel

## 6. Architecture Principles
- Keep the project narrow and clean.
- Prefer simple, explainable architecture.
- Avoid overengineering.
- Keep backend, frontend, load tests, documentation, and AI workflow clearly separated.
- Decisions must be documented.
- Prefer explicit SQL for race-condition-sensitive database updates.
- Do not hide critical budget decrement logic behind unclear ORM behavior.

## 7. Documentation Expectations
- README.md should stay short and setup-focused.
- Detailed technical reasoning should go under docs/.
- AI_WORKFLOW.md should be updated throughout the project.
- ai_session/ should include readable AI session summaries.

## 8. Testing Expectations
- Basic Go tests are planned.
- A specific concurrency test is planned: budget=10, 100 concurrent impression attempts, remaining_budget=0, impression_count=10, status=paused, and budget never negative.
- Only successful budget deductions should increase impression_count.
- Requests beyond available budget must not create negative remaining_budget.
- k6 load testing is planned to demonstrate HTTP-level race condition safety.
- Tests have not been implemented or executed yet.

## 9. Human Control Rules
- The human developer controls scope and decisions.
- AI must provide a plan before modifying implementation files.
- Implementation files must not be modified without explicit human approval.
- AI suggestions must be reviewed, questioned, and corrected when necessary.
- AI must not implement unapproved features.
- AI must not expand scope beyond the case study requirements.
