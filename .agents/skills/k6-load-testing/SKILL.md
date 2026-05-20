# k6 Load Testing Skill Instructions

This document guides the k6 load testing orchestration for the Gowit Mini Ad Platform. Implementation has not started yet.

## 1. Purpose
- k6 is planned to validate race condition safety at HTTP level.
- The goal is not performance benchmarking for vanity metrics.
- The goal is to prove that concurrent `POST /impression/:id` requests cannot make campaign budget negative.

## 2. Critical Scenario
The future k6 test should validate:
- initial campaign budget = 10
- campaign status = active
- 100 concurrent impression attempts
- remaining_budget = 0
- impression_count = 10
- status = paused
- budget never negative
- requests beyond available budget are not counted as successful impressions

## 3. Expected HTTP Flow
The k6 script should eventually:
- create a campaign or use a known test campaign ID
- send concurrent `POST /impression/:id` requests
- call `GET /stats/:id` after the load phase
- verify remaining_budget is never negative
- verify impression_count does not exceed the initial budget
- verify status becomes paused when budget reaches zero

## 4. Required Assertions
The test should check:
- HTTP responses are handled explicitly
- failed or rejected impression attempts are expected after budget is exhausted
- stats endpoint confirms final budget and impression values
- no negative remaining_budget appears

## 5. Backend Failure Handling
If k6 reveals a race condition bug:
- stop
- report the observed failure
- show the stats result
- explain the likely cause
- ask for approval before modifying backend code
- Do not directly modify backend implementation during k6 testing phase without explicit approval.

## 6. Scope Boundaries
Do not introduce:
- external load testing services
- cloud test infrastructure
- complex dashboards
- monitoring stacks
- performance tuning unrelated to the case requirements

## 7. Documentation and Honesty
- Do not claim k6 validation passed unless the script was actually executed.
- Do not claim race condition safety is proven until implementation and test evidence exist.
- Future k6 results should be documented in `README.md`, `AI_WORKFLOW.md`, and the relevant `ai_session` file only after execution.
- If k6 was not run, clearly state that it was not run.

*Note: This is a skill instruction document. k6 load testing implementation has not started, and no tests have been created or executed yet.*
