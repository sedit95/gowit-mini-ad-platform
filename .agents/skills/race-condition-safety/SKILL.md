# Race Condition Safety Skill Instructions

When reviewing or implementing the impression deduction logic:
- **Rule #1**: Never use a "read-then-update" pattern in application code without protection.
- **Rule #2**: Avoid in-memory mutexes as the final solution (they fail in multi-instance setups).
- **Preferred Solution**: Rely on PostgreSQL atomic conditional updates (e.g., UPDATE ... WHERE budget > 0) or transaction-level locking (SELECT ... FOR UPDATE).
- **Proof Requirement**: The solution must be proven via Go concurrency tests and/k6 load tests.

*Note: Do not write SQL implementation yet.*
