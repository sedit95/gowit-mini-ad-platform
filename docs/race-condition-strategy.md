# Planned Race Condition Strategy

**Status:** *Pre-implementation planning phase. Implementation and tests do not exist yet.*

## The Risk
The `POST /impression/:id` endpoint is the critical path. Many concurrent requests may hit the same campaign. Unprotected, unsafe read-then-update logic could cause the budget to become negative or over-count impressions.

## Forbidden Strategies
- Reading campaign budget in Go.
- Checking `remaining_budget` in application memory.
- Decrementing in memory and saving later without database-level atomicity/locking.
- In-memory Go mutex as the final multi-instance-safe solution.

## Preferred Strategy
- PostgreSQL atomic conditional update.

### Planned SQL Shape Concept
*Note: This is strategy documentation, not the final migration/implementation SQL.*
```sql
UPDATE campaigns
SET remaining_budget = remaining_budget - 1,
    impression_count = impression_count + 1,
    status = CASE WHEN remaining_budget - 1 = 0 THEN 'paused' ELSE status END
WHERE id = given_campaign_id
  AND deleted_at IS NULL
  AND status = 'active'
  AND remaining_budget > 0
RETURNING id, remaining_budget, impression_count, status
```

### Why Safe
- PostgreSQL applies the row update atomically.
- The `remaining_budget > 0` constraint ensures it never drops below zero.
- Only successful atomic updates increment the `impression_count`.
- Status naturally becomes `paused` in the exact moment `remaining_budget` hits 0.

## Accepted Alternative
- A PostgreSQL transaction using row-level locking (e.g., `SELECT ... FOR UPDATE`) is acceptable only if clearly justified. Atomic conditional update remains the preferred direction.

## Failure Handling
If the atomic update returns no row, the service may perform a read to distinguish between:
- `not_found`
- `soft_deleted`
- `campaign_not_active`
- `budget_exhausted`

*This read is strictly for determining the correct response reason, not for performing the decrement.*

## Expected Business Outcomes
- **Successful deduction:** `accepted: true`
- **Budget exhausted:** `accepted: false, reason: budget_exhausted`
- **Not active:** `accepted: false, reason: campaign_not_active`
- **Not found/soft deleted:** `404`

## Critical Validation Target
- `initial_budget = 10`, `remaining_budget = 10`, `status = active`
- 100 concurrent impression attempts
- Final `remaining_budget = 0`, final `impression_count = 10`, final `status = paused`
- Budget never negative
- `accepted: true` count = 10
- `accepted: false` count = 90
