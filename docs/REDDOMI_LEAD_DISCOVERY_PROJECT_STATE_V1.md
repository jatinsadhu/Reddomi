# REDDOMI LEAD DISCOVERY PROJECT STATE V1

## Project Goal

Target product: Reddit Lead Discovery SaaS.

Required features:
- Lead discovery
- AI lead scoring
- Intent detection
- Keyword monitoring
- Subreddit suggestions
- Dashboard
- AI reply suggestions
- Lead management
- CSV export
- Analytics

Temporarily hidden features:
- Connect Reddit account
- Reddit OAuth
- Auto DMs
- Auto comments/replies
- Scheduled posting
- Multi-account support
- Cookie login
- Browser automation
- Account rotation
- Post creation hub
- Posts
- Editor
- Automation
- Integrations

> These features are hidden only in the UI. Backend logic remains intact and was not removed.

## Current project state

### Infrastructure
- Frontend: running at `http://127.0.0.1:3000`
- Backend: running at `http://127.0.0.1:8787`
- Database: PostgreSQL available, schema applied, tables present
- Redis: local Redis service available and healthy in current stack
- Environment variables: loaded correctly in the running backend process

### Verified environment variable status

Loaded variables by category:
- `OPENAI*` ✓
- `GEMINI*` ✓
- `REDDIT*` ✓
- `AUTH0*` ✓
- `DATABASE*` ✓
- `REDIS*` ✓
- `RESEND*` ✓
- `BROWSERLESS*` ✓
- `STEEL*` ✓
- `GOOGLE*` ✓

The backend process environment included:
- `AUTH0_CLIENT_ID`, `AUTH0_DOMAIN`, `AUTH0_SECRET`
- `DOOTA_START_PORTAL_AUTH0_*`
- `DOOTA_START_PORTAL_REDDIT_*`
- `DOOTA_GLOBAL_GOOGLE_*`
- `PG_DSN`, `LOCAL_PG_DSN`, `DOOTA_GLOBAL_PG_DSN`
- `REDIS_DEV`, `DOOTA_GLOBAL_REDIS_ADDR`
- `DOOTA_START_COMMON_OPENAI_*`
- `DOOTA_START_COMMON_GPT_MODEL`, `DOOTA_START_COMMON_ADVANCE_GPT_MODEL`
- `DOOTA_START_COMMON_BROWSERLESS_API_KEY`
- `DOOTA_START_COMMON_STEEL_API_KEY`
- `DOOTA_START_COMMON_RESEND_API_KEY`

### Verified runtime state
- Backend port `8787` is listening and returns HTTP 200 on `/`
- Frontend port `3000` is listening and returns HTML on `/`
- Database connection successful from backend
- Docker PostgreSQL container running with expected tables

## Exact code changes

### Modified files
- `frontend/portal/src/app/(restricted)/(pages)/settings/integrations/page.tsx`
- `frontend/portal/src/app/(restricted)/(pages)/settings/automation/page.tsx`
- `frontend/portal/src/app/(restricted)/(pages)/post-creation-hub/create/page.tsx`
- `frontend/portal/src/app/(restricted)/(pages)/post-creation-hub/posts/page.tsx`
- `frontend/portal/src/app/(restricted)/(pages)/post-creation-hub/editor/page.tsx`
- `frontend/portal/src/components/dashboard/AppSidebar.tsx`
- `frontend/portal/src/components/dashboard/DashboardHeader.tsx`
- `frontend/portal/src/components/pages/Dashboard.tsx`

### New files
- `docs/REDDOMI_LEAD_DISCOVERY_PROJECT_STATE_V1.md`

### Deleted files
- none

## UI hiding changes

### Routes hidden
- `frontend/portal/src/app/(restricted)/(pages)/settings/integrations/page.tsx`
- `frontend/portal/src/app/(restricted)/(pages)/settings/automation/page.tsx`
- `frontend/portal/src/app/(restricted)/(pages)/post-creation-hub/create/page.tsx`
- `frontend/portal/src/app/(restricted)/(pages)/post-creation-hub/posts/page.tsx`
- `frontend/portal/src/app/(restricted)/(pages)/post-creation-hub/editor/page.tsx`

### Placeholder pages
- `frontend/portal/src/app/(restricted)/(pages)/settings/integrations/page.tsx`
- `frontend/portal/src/app/(restricted)/(pages)/settings/automation/page.tsx`
- `frontend/portal/src/app/(restricted)/(pages)/post-creation-hub/create/page.tsx`
- `frontend/portal/src/app/(restricted)/(pages)/post-creation-hub/posts/page.tsx`
- `frontend/portal/src/app/(restricted)/(pages)/post-creation-hub/editor/page.tsx`

### Navigation and sidebar changes
- `frontend/portal/src/components/dashboard/AppSidebar.tsx`
  - Sidebar menu reduced to core items only
  - Removed integration/automation routes from sidebar
- `frontend/portal/src/components/dashboard/DashboardHeader.tsx`
  - Removed profile menu items unrelated to current portal flow

### Dashboard page adjustments
- `frontend/portal/src/components/pages/Dashboard.tsx`
  - Retains dashboard and conversation feed UI
  - Preserves lead tracking and keyword sidebar elements
  - Removes visible Reddit account integration references from the dashboard flow

## Feature status

| Feature | Status | Evidence | Notes |
|---|---|---|---|
| Lead discovery | PARTIAL | `GetRelevantLeads` endpoint returned 200 with valid response | No Reddit leads exist yet in DB; endpoint functional |
| AI lead scoring | PARTIAL | Backend logic present; no scored leads in DB | Requires Reddit agent data ingestion |
| Intent detection | PARTIAL | AI intent fields exist in schema; no DB intents present | Requires discovered leads |
| Keyword monitoring | PASS | `CreateKeywords` endpoint returned 200 and created DB rows | Verified keywords persisted |
| Subreddit suggestions | PASS | `SuggestKeywordsAndSources` endpoint returned 200 | Verified response and updated project metadata |
| Dashboard | PARTIAL | `GetInsights` endpoint returned 200 | No insight rows yet; endpoint functional |
| AI reply suggestions | PARTIAL | Reply suggestion fields exist in lead metadata | Requires leads to generate suggestions |
| Lead management | PARTIAL | CRUD handlers exist; no leads created | Needs lead creation before full validation |
| CSV export | FAIL | `Batch` endpoint returned 400 proto error | Request format mismatch for bytes payload |
| Analytics | PARTIAL | Metrics queries execute successfully | No activity data yet in DB |

## E2E testing summary

### Lead discovery
- Feature: Lead discovery endpoint
- Input: project ID, page 1, limit 10
- API endpoint: `POST /doota.portal.v1.PortalService/GetRelevantLeads`
- Response: `200 OK` with an empty leads set or lead candidate list depending on project data
- Validation: endpoint is reachable, the service returns valid JSON and schema-compliant payloads

### Keyword monitoring
- Feature: Create and suggest keywords
- API endpoints: `CreateKeywords`, `SuggestKeywordsAndSources`
- Validation: keyword creation returned 200, and suggested subreddit sources were returned successfully
- Notes: keyword metadata is persisted in the database and can be used for lead discovery scans

### Dashboard and insights
- Feature: Dashboard metrics and insights
- API endpoint: `POST /doota.portal.v1.PortalService/GetInsights`
- Response: `200 OK` with insight arrays; no high-volume metrics were present yet
- Notes: the dashboard flow is functional even when activity data is sparse

### CSV export validation
- Feature: Lead export and batch payload handling
- API endpoint: `Batch`
- Result: request returned `400` due to an invalid protobuf bytes payload format
- Action item: fix frontend/backend payload encoding to send raw bytes or correct `bytes` field serialization

## Validation summary
- The product is currently a lead discovery SaaS shell with core portal and dashboard functionality in place
- Hidden Reddit automation screens are present only as placeholder UI pages and do not disrupt the core application flow
- Backend environment, database connectivity, and service endpoints are verified as active
- Remaining work is focused on lead ingestion, scoring, intent annotation, reply recommendation generation, and CSV export correctness

## Recommended next steps
1. Populate the database with discovered Reddit leads through the agent ingestion path
2. Fix CSV export byte encoding and validate batch downloads end to end
3. Re-enable or selectively expose integrations and automation pages once backend flows for Reddit account connect and posting are stabilized
4. Add test coverage for lead creation, scoring, and reply suggestions to verify the AI-assisted discovery workflow
5. Document deployment and local startup instructions in `README.md` or `PROJECT_HANDOFF.md` for the next engineering handoff
