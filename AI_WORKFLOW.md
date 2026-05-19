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
