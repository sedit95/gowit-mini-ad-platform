# k6 Load Testing Skill Instructions

When working on load tests:
- Focus on concurrent impression testing against the /impression/:id endpoint.
- Aggressively attempt to force the campaign budget below zero.
- Validate that the expected result (HTTP 400/409, or ignored deduction) occurs and the budget remains non-negative.
- Document the k6 test results.

*Note: Do not create the k6 script yet.*
