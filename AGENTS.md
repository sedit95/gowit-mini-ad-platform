# Agent Governance for Gowit Mini Ad Platform

This project is an AI-native Mini Campaign Management Platform for the Gowit case study.

## Core Directives
- **Mandatory Stack**: Go (Backend), React + TypeScript (Frontend), PostgreSQL (Database). Docker Compose and k6 load testing are planned.
- **Scope Limitations**: The scope is strictly limited to the case study requirements. Features like authentication, payments, multi-tenancy, advanced analytics, and unnecessary microservices are strictly forbidden. Do not implement unapproved features.
- **Coding Discipline**: Keep the architecture simple, clean, and narrow. Avoid overengineering.
- **Human Control**: The human developer owns all final decisions. AI agents must request review for significant architectural changes.
- **Race Condition Sensitivity**: The /impression/:id endpoint must be safe from concurrent requests. The campaign budget must never go negative, auto-pausing when it reaches zero.
- **Testing**: Basic Go tests, a specific concurrency test, and k6 load tests are required.
- **Documentation**: README must be setup-focused. Detailed reasoning belongs in docs/. AI workflow must be tracked accurately.
- **Action Control**: Do not generate implementation code without explicit instruction from the human developer.
