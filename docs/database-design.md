# Planned Database Design

**Status:** *Pre-implementation planning phase. No migration SQL has been created yet.*

## Overview
PostgreSQL is the planned database. The design separates the initial budget concept from the remaining budget to ensure data integrity and track spend accurately.

## Planned `campaigns` Table
- `id`
- `title`
- `currency`
- `initial_budget`
- `remaining_budget`
- `impression_count`
- `status`
- `start_date`
- `end_date`
- `created_at`
- `updated_at`
- `deleted_at`

## Budget Mechanics
- **Terminology:** While the UI and domain logic might just use the word "budget", the database design strictly distinguishes `initial_budget` and `remaining_budget`.
- **Spent Calculation:** `spent_budget` is dynamically determined: `spent_budget = initial_budget - remaining_budget`.
- **Race Condition Protection:** The `remaining_budget` field will be the critical target used for safe, atomic decrement logic.
