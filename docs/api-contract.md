# Planned API Contract

**Status:** *Pre-implementation planning phase. The API is not yet implemented.*

## Planned Endpoints
The endpoint set is intentionally narrow. No authentication, user, or analytics dashboard endpoints will be added.

- `GET /campaigns`: List campaigns
- `POST /campaigns`: Create a new campaign
- `GET /campaigns/:id`: Retrieve a specific campaign
- `PUT /campaigns/:id`: Update a specific campaign
- `DELETE /campaigns/:id`: Soft delete a specific campaign
- `POST /impression/:id`: Record an impression, decrementing the campaign budget
- `GET /stats/:id`: Retrieve campaign statistics

## Campaign Create (`POST /campaigns`)
- **Accepts:** `title`, `budget`, `currency`, `start_date`, `end_date`, `status` (optional, default active).
- **Validation:** `title` is required, `budget > 0`, `currency` required (preferably 3 chars), `start_date` and `end_date` required, `end_date >= start_date`, `status` must be active, paused, or completed.
- **Persistence Mapping:** The API accepts the general field `budget`. The database maps this to `initial_budget = budget` and `remaining_budget = budget`. `impression_count` starts at 0.

## Campaign Update (`PUT /campaigns/:id`)
- Must not update budget in MVP.
- `budget`, `initial_budget`, `remaining_budget`, `impression_count`, `deleted_at`, and `created_at` are not updatable.
- If `budget` is sent in the request, return `validation_error`.

## Campaign Delete (`DELETE /campaigns/:id`)
- This is a soft delete. Hard deletes are not used.
- Soft deleted campaigns behave like `not_found` for detail, stats, and impression endpoints.

## Impression Endpoint (`POST /impression/:id`)
- Has no request body.
- **Successful deduction:**
  - `accepted: true`
  - `remaining_budget` decreases by 1
  - `impression_count` increases by 1
  - `status` becomes paused if `remaining_budget` reaches 0
- **Budget Exhausted:**
  - `accepted: false`, `reason: budget_exhausted`
  - Returns `200 OK` (expected business outcome).
- **Campaign Not Active:**
  - `accepted: false`, `reason: campaign_not_active`
  - Returns `200 OK` (expected business outcome).
- **Not Found or Soft Deleted:** Returns `404 not_found`.

## Stats Endpoint (`GET /stats/:id`)
- **Returns:** `campaign_id`, `title`, `currency`, `total_impressions`, `initial_budget`, `spent_budget`, `remaining_budget`, `status`.
- **Logic:** `total_impressions = impression_count` and `spent_budget = initial_budget - remaining_budget`.
- Available for active, paused, and completed campaigns.
- Soft deleted campaigns return `404 not_found`.

## Error Response Format & Status Decisions
Responses will not leak raw DB errors to clients. Errors use this simple format:
```json
{
  "error": {
    "code": "validation_error",
    "message": "Message here"
  }
}
```
- `validation_error`: 400
- `invalid_status`: 400
- invalid UUID / campaign id: 400
- `not_found`: 404
- `internal_error`: 500
- `budget_exhausted` and `campaign_not_active`: 200 OK with `accepted: false`
