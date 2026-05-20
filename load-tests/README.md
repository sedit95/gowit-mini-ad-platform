# Load Tests

**Note on Scope:** No executable k6 scripts are included in the final delivered implementation. HTTP-level k6 load validation was intentionally left out of the final scope. 

Designing and executing professional k6 validation requires separate load-test scenario design, target calibration, strict performance thresholds, and careful result interpretation. 

## Existing Validation

While HTTP-level k6 testing is pending, the core architectural logic has already been validated:
- **Go Integration Tests:** The project includes robust Go integration tests.
- **Backend Service/Repository Concurrency Test:** A dedicated Go concurrency test (`impression_concurrency_test.go`) directly tests the PostgreSQL atomic budget deduction logic.
- **Real PostgreSQL Validation:** The concurrency tests run against a real, isolated PostgreSQL database, confirming data integrity under high parallel load.
- **Docker Compose Validation:** A full containerized runtime smoke test has successfully demonstrated that the components integrate cleanly.

It is important to clearly distinguish these validation layers:
- The **existing Go concurrency test** validates the repository and service-level atomic budget protection logic directly against the database engine.
- A **k6 load test** would validate the HTTP-level behavior, networking limits, and API gateway routing under heavy external concurrent traffic.

## Recommended Future Validation

HTTP-level k6 testing is highly recommended as future validation work to confirm production readiness.

The intended future k6 scenario target is:
- Target Endpoint: `POST /impression/{id}`
- Campaign Setup: Initial `budget` = 10
- Execution: Generate many highly concurrent HTTP POST requests.
- Pass Criteria:
  - The total number of accepted impressions (HTTP 200) must exactly equal, and not exceed, 10.
  - The `remaining_budget` in the database must never drop below 0.
  - The final status of the campaign must correctly transition to `paused`.

*Note: No k6 test was executed during this delivery phase. Production-scale performance validation is not complete.*
