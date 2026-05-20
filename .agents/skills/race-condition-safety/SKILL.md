# Race Condition Safety Skill Instructions

This document guides the race condition prevention strategy for the Gowit Mini Ad Platform. Implementation has not started yet.

## 1. Core Safety Rule
- The `POST /impression/:id` endpoint is the most critical endpoint in this case study.
- Budget must never go negative under any condition.
- Race condition safety is more important than UI polish or optional features.

## 2. Forbidden Unsafe Pattern
Explicitly forbid unsafe read-then-update logic:
- read campaign budget
- check remaining budget in Go
- decrement in application memory
- save later

*This approach is unsafe under concurrent requests unless protected by database-level atomicity or locking.*

## 3. Preferred Strategy
- Prefer PostgreSQL-level atomic conditional update.
- Budget decrement, impression_count increment, and auto-pause status transition should happen in one atomic database operation.
- The update should only succeed when:
  - campaign exists
  - campaign is not soft-deleted
  - campaign status is active
  - remaining_budget > 0

## 4. Accepted Alternative
- PostgreSQL transaction with row-level locking may be considered only if clearly justified.
- If transaction + row lock is proposed, the reasoning must explain why it is safe.

## 5. Rejected Final Strategy
- Do not use an in-memory Go mutex as the final budget protection strategy.
- A Go mutex protects only a single backend process and does not remain safe with multiple backend instances.

## 6. Impression Counting Rule
- Only successful budget deductions should increase impression_count.
- Requests beyond available budget must not create negative remaining_budget.
- Requests beyond available budget must not be counted as successful impressions.

## 7. Auto-Pause Rule
- When remaining_budget reaches zero, campaign status must become paused.
- Auto-pause should be part of the same safe database operation or protected transaction.

## 8. Testing Proof
The final implementation must be proven with:
- Basic Go concurrency tests
- HTTP-level k6 load tests when available

**Critical Test Target:**
- initial budget = 10
- campaign status = active
- 100 concurrent impression attempts
- remaining_budget = 0
- impression_count = 10
- status = paused
- budget never negative

## 9. Documentation and Workflow
- Race condition reasoning must be documented in `AI_WORKFLOW.md` and `docs/race-condition-strategy.md` during the appropriate phase.
- Do not claim race condition safety is solved until implementation and tests actually exist.
- Do not claim tests passed unless they were actually executed.

*Note: This is a skill instruction document. Backend implementation has not started, and no tests have been created or executed yet.*
