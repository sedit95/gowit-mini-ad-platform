# AI Workflow

## Tools Used
- Google Antigravity (Current AI coding environment)

## Initial Prompting
The project started as an AI-native Gowit take-home case study. The first session began with strict rules: no initial file creation, no backend/frontend code generation, no dependency installation, no Docker/migration/k6 generation, and a clear directive to wait for the scaffolding prompt. The scaffolding prompt then instructed the creation of a placeholder repository structure without any application code.

## AI Decisions Accepted
- The initial placeholder folder structure was created.
- `AI_CONTEXT.md` was initialized and refined with project boundaries, mandatory stack, race condition risks, testing expectations, and human control rules.
- The `AI_WORKFLOW.md` skeleton, `.agents` structure, `docs` structure, and `ai_session` structure were accepted after review.

## AI Output Reviewed and Corrected
**Issue:** During the initial scaffolding, the AI agent incorrectly created a nested `gowit-mini-ad-platform` directory inside the existing repository root workspace.
**Detection & Correction:** The human developer detected this incorrect nested folder structure before implementation started. A correction was requested, and the AI agent moved the generated structure into the correct repository root and removed the nested folder.
**Verification:** The human developer verified the fix using PowerShell. `Get-ChildItem -Name` returned the expected root-level structure, and `Test-Path .\gowit-mini-ad-platform` successfully returned `False`.

## AI Suggestions Rejected
No implementation-level AI suggestions have been rejected yet because backend/frontend implementation has not started.

## Race Condition Review
The race condition risk (concurrent `/impression/:id` requests potentially causing negative budgets) has been identified early and recorded in `AI_CONTEXT.md`. The final implementation choice has not been created yet, though PostgreSQL-level atomic conditional update is the preferred direction to be validated in later phases.

## Human Corrections
- **Nested Folder Correction:** The human developer caught a repository structure issue where scaffolding was placed in an extra nested folder. The structure was flattened, and the fix was verified via PowerShell.

## Testing Evidence
No tests have been executed yet. A specific Basic Go concurrency test and k6 load testing are planned as future validation.

## Final Ownership Notes
The human developer is strictly controlling the scope, actively reviewing AI output, and explicitly preventing unapproved implementation or feature expansion beyond the case study requirements.

## Agent Governance Refinement
During the scaffolding phase, the human developer directed the refinement of the agent governance files (AGENTS.md, .agents/workflows/, and .agents/skills/). These documents now explicitly define agent roles, strict workflow constraints, and skill expectations (especially regarding race condition safety). **No implementation code has been generated at this stage.**

## Master Prompt Preparation
A final master prompt was prepared before implementation started, defining the AI agent as a controlled, phase-based engineering assistant. The prompt strictly requires explicit human approval before any implementation, prevents scope expansion, and prioritizes race condition safety for `POST /impression/:id`. It explicitly rejects unsafe read-then-update logic and in-memory Go mutexes as final multi-instance-safe solutions, preferring instead a PostgreSQL-level atomic conditional update. It also mandates honest updates to `AI_WORKFLOW.md` and `ai_session` files. The prompt was delivered, the AI agent acknowledged the governance rules, and confirmed it would not implement anything without explicit approval. No backend/frontend implementation was generated during this phase.

## AI Context Refinement
`AI_CONTEXT.md` was refined before implementation to clearly align with the case study constraints. The frontend stack was clarified as Vite + React + TypeScript, explicitly excluding Next.js. Scope boundaries were expanded to clearly define forbidden technologies. The critical race condition rules were strengthened: unsafe read-then-update logic is explicitly forbidden, PostgreSQL-level atomic conditional updates are documented as the preferred direction, and in-memory Go mutexes are explicitly rejected as the final budget protection strategy because they do not protect multiple backend instances. Human approval rules and testing expectations were also tightened. No backend, frontend, Docker, migration, k6, or test implementation was generated, and no tests were executed during this phase.
