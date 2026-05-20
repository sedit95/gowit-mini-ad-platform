# Decision Log

**Status:** *Pre-implementation planning phase.*

- **Frontend Stack:** Vite + React + TypeScript was selected. Next.js is explicitly not used.
- **Database:** PostgreSQL was selected as the source of truth.
- **Live Statistics:** Standard HTTP polling was selected for live stats. WebSockets are not used unless explicitly approved.
- **Caching/State:** Redis is not used unless explicitly approved.
- **Concurrency Protection:** In-memory Go mutex is rejected as a final budget protection strategy. PostgreSQL atomic conditional updates are preferred.
- **Scope Limits:** Authentication, RBAC, payments, multi-tenancy, and advanced analytics are explicitly out of scope.
- **Documentation Strategy:** `README.md` stays concise and setup-focused. Deeper technical reasoning is held in the `docs/` folder. `AI_WORKFLOW.md` and `ai_session` logs must remain strictly honest.
