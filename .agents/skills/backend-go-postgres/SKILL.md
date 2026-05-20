# Backend (Go + PostgreSQL) Skill Instructions

This document guides the backend implementation for the Gowit Mini Ad Platform. Implementation has not started yet.

## 1. Backend Architecture
- Use a simple layered Go backend.
- Preferred structure: `handler` -> `service` -> `repository`.
- Keep the architecture narrow and explainable.
- Avoid CQRS, event sourcing, microservices, and unnecessary abstractions.

## 2. Go Discipline
- Use idiomatic Go.
- Keep functions small and readable.
- Use `context.Context` for database operations.
- Return explicit errors. Do not swallow errors.
- Avoid global mutable state for business logic.

## 3. PostgreSQL Discipline
- PostgreSQL is the source of truth.
- Prefer explicit SQL for critical data integrity operations.
- Do not hide race-condition-sensitive budget decrement logic behind unclear ORM behavior.

## 4. Campaign Domain Awareness
- Campaign CRUD is required.
- Campaign fields include `title`, `budget`, `currency`, `start date`, `end date`, and `status`.
- Status values are `active`, `paused`, `completed`.
- Campaigns must use soft delete.
- Stats must include total impressions, spent budget, and remaining budget.
- The final database design should distinguish `initial_budget` and `remaining_budget`, but detailed DB design will be documented later in `docs/database-design.md`.

## 5. Critical Impression Logic
- `POST /impression/:id` deducts 1 unit from campaign budget for each successful impression.
- Budget must never go negative.
- When budget reaches zero, campaign must auto-pause.
- Only successful budget deductions should increase impression_count.

## 6. Race Condition Safety
- Never implement unsafe read-then-update budget decrement.
- Do not use an in-memory Go mutex as the final multi-instance-safe solution.
- Prefer PostgreSQL atomic conditional update for budget decrement, impression count increment, and auto-pause transition.
- PostgreSQL transaction with row-level locking can be considered only if clearly justified.

## 7. Validation and Error Handling
- Validate request input.
- Reject invalid status values.
- Reject invalid budget values.
- Validate date ranges.
- Return clear API errors.
- Do not leak raw internal DB errors directly.

## 8. Testing Expectations
- Campaign CRUD tests are expected.
- Soft delete behavior should be tested.
- Stats calculation should be tested.
- Impression decrement should be tested.
- **Critical Concurrency Test Target:**
  - initial budget = 10
  - 100 concurrent impression attempts
  - remaining_budget = 0
  - impression_count = 10
  - status = paused
  - budget never negative

*Note: This is a skill instruction document. Backend implementation has not started, and no tests have been created or executed yet.*
