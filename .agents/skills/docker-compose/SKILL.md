# Docker Compose Skill Instructions

This document guides the Docker Compose orchestration for the Gowit Mini Ad Platform. Implementation has not started yet.

## 1. Docker Compose Goal
- The goal is to run the full local system with a single command.
- Target command will be: `docker compose up --build`
- Do not claim Docker Compose works until it is actually implemented and tested.

## 2. Planned Services
Docker Compose should eventually include:
- PostgreSQL
- Go backend
- Vite + React + TypeScript frontend

## 3. Environment Variables
The Docker setup should align with `.env.example` and future app configuration.
Expected variables may include:
- `DATABASE_URL`
- `BACKEND_PORT`
- `FRONTEND_PORT`
- `VITE_API_BASE_URL`
- `POSTGRES_DB`
- `POSTGRES_USER`
- `POSTGRES_PASSWORD`

## 4. Service Dependency Awareness
- Backend depends on PostgreSQL.
- Frontend depends on backend API availability.
- Use simple dependency rules.
- Add health checks only if they improve reliability without overcomplicating the setup.

## 5. Scope Boundaries
Do not introduce:
- Kubernetes
- Cloud deployment
- CI/CD pipelines
- Monitoring stacks
- Production orchestration
- Nginx reverse proxy unless explicitly approved
- Redis unless explicitly approved

## 6. File Boundaries for Future Docker Phase
When the Docker phase is explicitly approved, expected files may include:
- `docker-compose.yml`
- `backend/Dockerfile`
- `frontend/Dockerfile`
- `.env.example` updates if needed
- `README.md` updates after commands are verified
- `AI_WORKFLOW.md` and relevant `ai_session` updates

## 7. Documentation and Honesty
- Do not add run commands to `README.md` until Docker Compose exists and has been validated.
- Do not claim `docker compose up --build` works unless it was actually executed successfully.
- If Docker is not tested, clearly say it was not tested.

*Note: This is a skill instruction document. Docker implementation has not started, and no Docker files have been created or tested yet.*
