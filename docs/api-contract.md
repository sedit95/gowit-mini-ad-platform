# Planned API Contract

**Status:** *Pre-implementation planning phase. The API is not yet implemented.*

## Planned Endpoints
- `GET /campaigns`: List campaigns
- `POST /campaigns`: Create a new campaign
- `GET /campaigns/:id`: Retrieve a specific campaign
- `PUT /campaigns/:id`: Update a specific campaign
- `DELETE /campaigns/:id`: Soft delete a specific campaign
- `POST /impression/:id`: Record an impression, decrementing the campaign budget
- `GET /stats/:id`: Retrieve campaign statistics

## Core Concepts
- **Campaign Fields:** `title`, `currency`, `start date`, `end date`, `status`.
- **Status Values:** `active`, `paused`, `completed`.
- **Stats Fields:** `total impressions`, `spent budget`, `remaining budget`.
- **Soft Delete:** Campaigns are never permanently deleted; they use soft delete semantics.

## Validation & Responses
- Expected request validation (e.g., date ranges, valid status, valid budget).
- Standardized error response expectations (e.g., clear API errors without leaking raw DB errors).
