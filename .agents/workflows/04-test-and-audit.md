# Test and Audit Workflow

This workflow ensures project reliability:
- **Unit/Integration Tests**: Basic Go tests to verify CRUD logic.
- **Concurrency Testing**: Specific Go test simulating 100 concurrent impression attempts to prove the budget never goes negative.
- **Load Testing**: k6 load test validation.
- **Manual API Verification**: Validating no negative budget occurrences manually if necessary.
- **Audit**: Review documentation to ensure it matches the final behavior.

*Note: This is workflow guidance only. Do not write actual test code until instructed.*
