# Planned Database Design

**Status:** *Pre-implementation planning phase. No migration SQL has been created yet.*

## Overview
PostgreSQL is the planned database. The design focuses strictly on the campaigns table, with no impression event table or analytics table planned.

## Primary Table: `campaigns`
- `id`: UUID primary key
- `title`: String
- `currency`: VARCHAR(3)
- `initial_budget`: INTEGER
- `remaining_budget`: INTEGER
- `impression_count`: INTEGER
- `status`: String (active, paused, completed)
- `start_date`: TIMESTAMPTZ
- `end_date`: TIMESTAMPTZ
- `created_at`: TIMESTAMPTZ
- `updated_at`: TIMESTAMPTZ
- `deleted_at`: TIMESTAMPTZ NULL

## Budget Model
- The API/domain can use the general word "budget" for creation input.
- The database strictly distinguishes `initial_budget` and `remaining_budget`.
- **On Create:** `initial_budget = budget`, `remaining_budget = budget`, `impression_count = 0`.
- **Spent Budget:** `spent_budget` is derived, not stored (`spent_budget = initial_budget - remaining_budget`).
- **Race Condition Safety:** `remaining_budget` is used directly for safe atomic decrement logic.

## Planned Constraints
- `initial_budget > 0`
- `remaining_budget >= 0`
- `impression_count >= 0`
- `status` IN ('active', 'paused', 'completed')
- `end_date >= start_date`

## Soft Delete
- Governed by `deleted_at TIMESTAMPTZ NULL`.
- `DELETE` endpoint sets `deleted_at`.
- Soft deleted records are completely excluded from list/detail/stats/impression endpoint behaviors.

## Race-Condition Support
- The atomic update depends strictly on: `remaining_budget > 0`, `status = 'active'`, and `deleted_at IS NULL`.
- The update should decrement `remaining_budget`, increment `impression_count`, and auto-pause in one safe DB operation.
