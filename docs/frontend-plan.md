# Frontend Plan: Gowit Mini Ad Platform

## 1. Status
This document is for frontend planning only. Frontend implementation has not started yet and will happen in a later phase. Backend validation has already been completed at the service/repository level with real PostgreSQL integration tests. However, HTTP-level k6 load validation and full Docker Compose validation are still pending.

## 2. Frontend Scope
The frontend will be built using Vite + React + TypeScript. **Next.js is explicitly rejected.** The frontend will be kept narrow and strictly focused on the case study requirements.

**Required screens:**
- Campaign List Page
- Campaign Create Page/Form
- Campaign Detail + Stats Page

**Out of scope:**
- Authentication / Login / Register
- Admin panel
- Payments
- Advanced analytics dashboard / Charts
- WebSockets
- Complex dashboards
- Multi-tenancy
- User management

## 3. Routes
- `/` -> Campaign List Page
- `/campaigns/new` -> Campaign Create Page
- `/campaigns/:id` -> Campaign Detail + Stats Page

## 4. Backend API Mapping
The frontend API client will map to the backend endpoints as follows:
- `getCampaigns` -> `GET /campaigns`
- `createCampaign` -> `POST /campaigns`
- `getCampaign` -> `GET /campaigns/{id}`
- `updateCampaign` -> `PUT /campaigns/{id}`
- `deleteCampaign` -> `DELETE /campaigns/{id}`
- `recordImpression` -> `POST /impression/{id}`
- `getStats` -> `GET /stats/{id}`

## 5. TypeScript Models
The planned frontend types mapping to the API contracts:

- `CampaignStatus` = `'active' | 'paused' | 'completed'`
- `Campaign`
- `CreateCampaignRequest` (Includes `budget`)
- `UpdateCampaignRequest` (Must NOT include `budget`)
- `StatsResponse`
- `ImpressionResponse` (Note: `accepted=false` is a business outcome, not necessarily an HTTP error)
- `ApiError`

*Note: Date fields will be treated as strings on the frontend client boundary.*

## 6. Campaign List Page Plan
- Displays: title, currency, initial budget, remaining budget, impression count, status, start/end dates.
- Actions: Create Campaign, Refresh, View Details, Delete.
- Deletions will require confirmation and subsequently refresh the list.
- Standard loading, empty, and error states will be handled.
- Complexity such as pagination, advanced filtering, or charts is excluded.

## 7. Campaign Create Page Plan
- Fields: `title`, `budget`, `currency`, `start_date`, `end_date`, `status`.
- `status` defaults to `active`.
- **Frontend Validation:**
  - `title`: required
  - `budget`: integer > 0
  - `currency`: required, preferably 3 characters
  - `start_date` / `end_date`: required
  - `end_date` >= `start_date`
  - `status`: must be valid
- Form state may manage the budget as a string, but the API payload will send `budget` as a number.
- `datetime-local` input values will be converted to ISO strings before submission.
- On success, navigation directs the user to the Campaign Detail page.
- No heavy form library is planned.

## 8. Campaign Detail + Stats Page Plan
- Integrates multiple API calls:
  - `GET /campaigns/{id}`
  - `GET /stats/{id}`
  - `POST /impression/{id}`
  - `DELETE /campaigns/{id}`
  - `PUT /campaigns/{id}` (for a small metadata update section)
- Displays campaign details alongside real-time statistics.
- Includes a "Record Impression" button.
- Includes a "Refresh Stats" button.
- Handles and displays business messages for rejected impressions (`accepted=false`) such as `budget_exhausted` and `campaign_not_active`.
- Contains a small update section for `title`, `currency`, dates, and `status`. **Budget updates are strictly forbidden.**
- Deletions navigate back to the Campaign List after success.
- Excludes charts or advanced analytics.

## 9. Polling Strategy
- Polling is restricted solely to the detail/stats page.
- Target: Poll `GET /stats/{id}`.
- Interval: 3000 ms.
- Behavior: Stops polling upon component unmount.
- Triggers an immediate stats refresh directly after a manual "Record Impression".
- **No WebSocket.**
- **No optimistic budget decrement.** The backend remains the absolute source of truth.

## 10. Error and Loading State Plan
- `ApiError` structure: `status`, `code`, `message`.
- Mappings:
  - `validation_error` and `invalid_status` -> Displayed as form/page errors.
  - `not_found` -> Renders a specific not found state.
  - `internal_error` / network error -> Renders a generic error.
- Implements an empty state for the List page.
- Implements a submit loading state for the Create page.
- Implements initial loading and background polling states for the Detail page.
- Handled gracefully: `accepted=false` from `recordImpression` shows a business warning/message, rather than crashing the page as an unhandled error.
- **No toast libraries or global error boundaries are planned.**

## 11. Planned File Structure
```text
frontend/
  package.json
  package-lock.json
  index.html
  vite.config.ts
  tsconfig.json
  tsconfig.node.json
  src/
    main.tsx
    App.tsx
    api/
      client.ts
      campaigns.ts
    types/
      campaign.ts
    pages/
      CampaignListPage.tsx
      CampaignCreatePage.tsx
      CampaignDetailPage.tsx
    components/
      StatusBadge.tsx
      LoadingState.tsx
      ErrorMessage.tsx
    utils/
      date.ts
```

## 12. Dependency Boundaries
**Allowed/Planned:**
- `react`, `react-dom`
- `vite`
- `typescript`
- `@vitejs/plugin-react`
- `react-router-dom`

**Explicitly Avoid:**
- Redux, React Query
- Axios
- Formik, React Hook Form
- Zod, Yup
- Moment, Dayjs, date-fns
- Tailwind CSS
- MUI, AntD
- Chart libraries, WebSocket libraries

## 13. Implementation Boundaries
- Implementation will happen in step 8.
- The frontend must not modify backend contracts or assume different backend behaviors.
- If an API mismatch is identified, it must be reported rather than silently altering backend code.
- No frontend code is generated during this planning phase.
