# Documentation and AI Workflow Skill Instructions

This document guides the documentation and AI workflow management for the Gowit Mini Ad Platform. Implementation has not started yet.

## 1. Documentation Purpose
- Documentation is part of the evaluation, not an afterthought.
- Keep documentation reviewer-friendly, honest, and useful.
- Avoid bloated documentation that hides the actual decisions.

## 2. README.md Rules
- `README.md` should stay concise and setup/status-focused.
- `README.md` is the reviewer entry point.
- Do not add run commands until they actually work.
- Do not claim backend, frontend, Docker, k6, migrations, or tests exist until they actually exist.
- Link or refer to `docs/` for deeper reasoning.

## 3. docs/ Rules
Use `docs/` for deeper technical reasoning:
- `docs/architecture.md` for system architecture decisions.
- `docs/api-contract.md` for API contract decisions.
- `docs/database-design.md` for database modeling decisions.
- `docs/race-condition-strategy.md` for concurrency and budget safety reasoning.
- `docs/decisions.md` for accepted/rejected decisions and trade-offs.

**Important Future DB Documentation Note:**
- `docs/database-design.md` must later clarify the distinction between `initial_budget` and `remaining_budget`.
- This distinction will support accurate spent budget calculation, remaining budget tracking, and race-condition-safe decrement logic.
- Do not claim this DB design is implemented until migrations and code actually exist.

## 4. AI_WORKFLOW.md Rules
- `AI_WORKFLOW.md` must be honest, chronological, and evidence-based.
- Record real AI usage, accepted decisions, rejected suggestions, and human corrections.
- Do not invent AI mistakes just to look good.
- Do not invent future events.
- Do not claim implementation progress that has not happened.
- Do not claim race condition safety is solved before implementation and validation exist.
- Do not claim tests passed unless they were actually executed.

## 5. ai_session/ Rules
`ai_session` files should contain readable summaries of meaningful AI interactions.
For relevant phases, capture:
- what was requested from AI
- what AI produced
- what the human developer accepted
- what the human developer rejected or corrected
- files changed
- whether tests or validations were run
- whether implementation code was generated

Do not dump irrelevant noisy logs unless specifically needed.

## 6. Human Correction Documentation
When AI produces an incorrect, incomplete, risky, or out-of-scope output:
- record what happened
- record how the human developer detected it
- record what correction was requested
- record the final verified outcome

**Known Example Already Recorded:**
- AI created a nested `gowit-mini-ad-platform` folder during scaffolding.
- Human developer detected it.
- AI corrected the structure.
- PowerShell verification confirmed the fix.

## 7. Test and Validation Evidence
- If tests were not run, clearly state tests were not run.
- If Docker was not tested, clearly state Docker was not tested.
- If k6 was not executed, clearly state k6 was not executed.
- Record command outputs or summaries only after actual execution.
- Never write fake test evidence.

## 8. Scope and Honesty
- Do not hide unfinished work.
- Do not overstate completeness.
- Do not make the project look more implemented than it is.
- Clear limitations are better than fake completeness.

*Note: This is a skill instruction document. Documentation finalization for the whole project is not complete, implementation has not started, and no tests have been created or executed yet.*
