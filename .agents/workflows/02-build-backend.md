# Build Backend Workflow

This workflow guides the construction of the backend:
- **Language/Framework**: Go standard library or lightweight router.
- **Features**: Campaign CRUD, impression deduction endpoint (POST /impression/:id), stats endpoint (GET /stats/:id).
- **Data rules**: Soft deletion for campaigns, strict validation of fields.
- **Persistence**: PostgreSQL integration.
- **Critical Focus**: Race condition safety for the impression endpoint.

*Note: This is workflow guidance only. Do not write implementation code until instructed.*
