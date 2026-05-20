# Frontend (React + TypeScript) Skill Instructions

This document guides the frontend implementation for the Gowit Mini Ad Platform. Implementation has not started yet.

## 1. Frontend Stack
- Use Vite + React + TypeScript.
- Do not use Next.js.
- Keep the frontend simple, fast, and case-focused.

## 2. Required Screens
- Campaign list page.
- New campaign creation form.
- Campaign detail page with live statistics.
- Do not add authentication screens, admin panels, user profile screens, or unrelated dashboards.

## 3. Live Statistics
- Use polling for live statistics.
- Do not introduce WebSockets unless explicitly approved.
- Polling should call `GET /stats/:id` at a reasonable interval during the detail page lifecycle.

## 4. TypeScript Discipline
- Use explicit TypeScript types.
- Define campaign-related types clearly.
- Avoid `any` unless absolutely necessary.
- Keep API request and response shapes aligned with the backend API contract.

## 5. API Integration
- Use a small typed API client.
- Keep API calls centralized under `frontend/src/api/`.
- Keep shared domain types under `frontend/src/types/`.
- Do not silently change backend contracts during frontend work.
- If an API mismatch is detected, stop and report it before modifying backend files.

## 6. UI Simplicity
- Prefer simple, readable UI components.
- Use status badges for active, paused, and completed campaigns.
- Include basic loading and error states.
- Avoid advanced charts, complex dashboards, analytics widgets, and unnecessary UI frameworks unless explicitly approved.

## 7. Form and Validation
- The campaign creation form should validate required fields.
- Validate title, budget, currency, start date, end date, and status.
- Frontend validation is for user experience only.
- Backend validation remains authoritative.

## 8. Scope Boundaries
- Do not add auth flows.
- Do not add payment flows.
- Do not add multi-tenant UI.
- Do not add advanced analytics.
- Do not add WebSocket logic unless explicitly approved.
- Do not modify backend files during frontend work without explicit approval.

## 9. Documentation and Honesty
- Do not claim frontend implementation exists until it is actually created.
- Do not claim frontend build or tests passed unless they were actually executed.
- If frontend implementation later changes API assumptions, document that in `AI_WORKFLOW.md` and the relevant `ai_session` file.

*Note: This is a skill instruction document. Frontend implementation has not started, and no tests have been created or executed yet.*
