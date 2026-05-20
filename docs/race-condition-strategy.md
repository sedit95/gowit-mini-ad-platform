# Planned Race Condition Strategy

**Status:** *Pre-implementation planning phase. Implementation and tests do not exist yet.*

## The Risk
The `POST /impression/:id` endpoint is highly susceptible to race conditions due to concurrent impression requests. Unprotected concurrent decrements could cause the campaign budget to become negative.

## Forbidden Strategies
- **Unsafe Read-Then-Update:** Reading the budget, checking it in Go, decrementing in memory, and saving back is strictly forbidden.
- **In-Memory Go Mutex:** A standard Go mutex is rejected as the final multi-instance-safe strategy because it only protects a single backend process.

## Planned Strategy
- **Preferred Direction:** PostgreSQL-level atomic conditional update. The budget decrement, impression count increment, and auto-pause status transition must happen in a single, atomic database operation.
- **Accepted Alternative:** A PostgreSQL transaction using row-level locking (e.g., `SELECT ... FOR UPDATE`) is acceptable only if clearly justified.

## Business Logic Rules
- A successful deduction increments `impression_count`.
- Requests beyond the available budget must not be counted as successful impressions.
- The campaign `status` must auto-pause immediately when `remaining_budget` reaches zero.

## Planned Validation
- Basic Go concurrency tests.
- k6 HTTP-level load tests.
