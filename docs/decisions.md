# Decision Log

**Status:** *Pre-implementation planning phase. Implementation has not started.*

## Core Technical Decisions
- Vite + React + TypeScript selected; Next.js explicitly not used.
- PostgreSQL selected as the single source of truth.
- Polling selected for live stats; WebSocket not used unless explicitly approved.
- Redis not used unless explicitly approved.
- Go in-memory mutex rejected as final budget protection.
- PostgreSQL atomic conditional update preferred. Transaction + row-level locking accepted only if justified.

## Backend Implementation Decisions
- A 7-endpoint backend API is enough for MVP.
- Campaign create accepts `budget`; DB stores `initial_budget` and `remaining_budget`.
- `PUT /campaigns/:id` does not update budget in the MVP.
- `DELETE` is a soft delete using `deleted_at`.
- Campaign ID will be a UUID.
- Budget uses INTEGER units because each impression deducts 1 unit.
- `spent_budget` is derived, not stored.
- No impression event table. No analytics table.
- Budget exhausted and campaign_not_active are business outcomes returned with `accepted: false` (200 OK).
- Backend uses a simple `handler/service/repository` structure.
- `chi` preferred for routing.
- `pgxpool` preferred for PostgreSQL.
- GORM/heavy ORM is explicitly not used.

## Scope & Workflow Limits
- Authentication, RBAC, payments, multi-tenancy, and advanced analytics are explicitly out of scope.
- README stays concise; docs hold deeper reasoning.
- `AI_WORKFLOW.md` and `ai_session` logs must remain strictly honest.

## Testing Strategy
- Concurrency testing should run against real PostgreSQL if possible.
- Mock tests are deemed insufficient to prove true database-level race condition safety.
