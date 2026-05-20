# Session 005: Frontend Planning

## What was planned
The frontend architecture and requirements were thoroughly planned prior to implementation. The frontend will be a narrow, case-focused application built with Vite + React + TypeScript. Implementation is explicitly deferred to Step 8.

## Screens & Route Plan
- `/` -> **Campaign List Page** (Shows summary list with title, currency, budgets, status, and dates. Includes create, refresh, detail, and delete actions).
- `/campaigns/new` -> **Campaign Create Page** (Form for title, budget, currency, dates, and status. Implements manual validation and navigates to details upon success).
- `/campaigns/:id` -> **Campaign Detail + Stats Page** (Displays full details and stats. Features "Record Impression" and "Refresh Stats" buttons. Allows minor metadata updates but explicitly forbids budget updates. Successful deletions redirect to the list).

## API Mapping
- `getCampaigns` -> `GET /campaigns`
- `createCampaign` -> `POST /campaigns`
- `getCampaign` -> `GET /campaigns/{id}`
- `updateCampaign` -> `PUT /campaigns/{id}`
- `deleteCampaign` -> `DELETE /campaigns/{id}`
- `recordImpression` -> `POST /impression/{id}`
- `getStats` -> `GET /stats/{id}`

## TypeScript Models
- `CampaignStatus`: `'active' | 'paused' | 'completed'`
- `Campaign`, `CreateCampaignRequest` (includes budget), `UpdateCampaignRequest` (excludes budget)
- `StatsResponse`, `ImpressionResponse`, `ApiError`
- Date fields are mapped as strings.
- Note: `accepted=false` in `ImpressionResponse` represents a business outcome, not necessarily an HTTP-level error.

## Polling & Real-time Strategy
- Polling is restricted to the Detail + Stats Page targeting `GET /stats/{id}` every 3000ms.
- Polling halts gracefully on component unmount.
- Recording an impression triggers an immediate explicit stats refresh.
- **Backend remains the absolute source of truth.** No optimistic budget decrements are permitted.
- WebSockets, Server-Sent Events (SSE), and heavy real-time infrastructure are explicitly rejected.

## Error & Loading Strategy
- Backend `ApiError` envelopes will be parsed into `status`, `code`, and `message`.
- Rejected impressions (`accepted=false`) will be presented as non-fatal business warnings.
- Explicit states for Loading, Empty, Not Found, Network Errors, and Form Validation will be built-in.
- No third-party toast libraries or global React error boundaries are planned.

## Dependency Boundaries
**Allowed/Planned:**
- `react`, `react-dom`, `vite`, `typescript`, `@vitejs/plugin-react`, `react-router-dom`.

**Explicitly Rejected:**
- Next.js, Redux, React Query, Axios, Formik, React Hook Form, Zod/Yup, Moment/Dayjs/date-fns, Tailwind CSS, MUI/AntD, Chart libraries, WebSocket clients.

## Pending Validation Context
- Backend service/repository validation has already succeeded against a real PostgreSQL instance (including the 100-concurrent-request budget safety test).
- **Frontend implementation and validation have NOT started.**
- HTTP-level k6 load validation is still pending.
- Docker Compose validation is still pending.
