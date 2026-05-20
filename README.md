# Gowit Mini Ad Platform

## Overview
This is an AI-native Mini Campaign Management Platform for the Gowit take-home case study. It will allow a retailer to manage ad campaigns and track impressions. The project focuses on correctness, race condition safety, AI workflow transparency, and clean scoped delivery.

## Current Status
**The repository is currently in the pre-implementation planning phase.**
Repository scaffolding, AI context, AI workflow tracking, and agent governance files have been prepared. Backend, frontend, Docker Compose, database migrations, k6 scripts, and Go tests have not been implemented yet.

## Tech Stack
- **Backend:** Go
- **Frontend:** Vite + React + TypeScript
- **Database:** PostgreSQL
- **Planned:** Docker Compose
- **Planned:** k6 load testing
- **Planned:** Basic Go tests

## Planned Scope
- Campaign CRUD
- Campaign fields: title, budget, currency, start date, end date, status
- Status values: active, paused, completed
- POST `/impression/:id` budget decrement
- Budget must never go negative
- Auto-pause when budget reaches zero
- GET `/stats/:id`
- Soft delete
- Campaign list
- New campaign creation form
- Campaign detail page with polling-based live statistics

## Out of Scope
- Authentication
- Authorization / RBAC
- User management
- Payments
- Multi-tenancy
- Advanced analytics
- Complex dashboards
- Message queues
- Event sourcing
- CQRS
- Microservices
- Kubernetes
- Cloud deployment
- Redis unless explicitly approved
- WebSockets unless explicitly approved

## Repository Structure
- `backend/`: Go application, tests, and database migrations (Planned)
- `frontend/`: Vite + React frontend (Planned)
- `load-tests/`: k6 load test scripts (Planned)
- `ai_session/`: Readable summaries of relevant AI interactions
- `.agents/`: Agent governance, skills, and workflow guidance
- `docs/`: Deeper technical reasoning and architecture decisions
- `scripts/`: Helper scripts (Planned)
- `AI_CONTEXT.md`: Project context and constraints given to the AI agent
- `AI_WORKFLOW.md`: Tracks AI usage, accepted decisions, corrections, and current progress
- `AGENTS.md`: Root-level agent governance document

## AI-Native Workflow
This repository includes AI workflow documentation.
- `AI_CONTEXT.md` describes the project context and constraints given to the AI agent.
- `AI_WORKFLOW.md` tracks AI usage, accepted decisions, corrections, and current progress.
- `ai_session/` contains readable summaries of relevant AI interactions.
- `AGENTS.md` and `.agents/` define agent governance, skills, and workflow guidance.

The AI agent is explicitly not allowed to implement without human approval. The project has already recorded one real AI correction: a nested project folder issue detected by the human developer and corrected before implementation.

## Documentation
- [Architecture](docs/architecture.md)
- [API Contract](docs/api-contract.md)
- [Database Design](docs/database-design.md)
- [Race Condition Strategy](docs/race-condition-strategy.md)
- [Decisions](docs/decisions.md)

Detailed technical reasoning will be added progressively as planning and implementation phases are approved.

## Running the Project
**Application implementation has not started yet.**
Running instructions will be added after backend, frontend, and Docker Compose phases are implemented.

## Testing and Load Testing
- Tests have not been implemented or executed yet.
- A Basic Go concurrency test is planned.
- k6 HTTP-level load testing is planned.
- The critical validation target is:
  - initial budget = 10
  - 100 concurrent impression attempts
  - remaining_budget = 0
  - impression_count = 10
  - status = paused
  - budget never negative

## Notes for Reviewers
The repository currently demonstrates the AI-native planning, governance, and documentation setup. Implementation will proceed in controlled phases. No backend/frontend application code has been generated yet.
