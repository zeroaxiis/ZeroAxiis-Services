# ZeroAxiis Backend Documentation

Welcome to the official development documentation for the **ZeroAxiis Backend**.

This documentation is the primary reference for designing, developing, maintaining, and extending the backend. Every architectural decision, API contract, database model, and development guideline is documented here before implementation.

---

# Project Overview

The ZeroAxiis Backend powers the official ZeroAxiis website, a Software Consulting & Open Source Development company.

The backend serves as the single source of truth for all dynamic content displayed on the website.

Current responsibilities include:

- Authentication
- Team Management
- Client Projects
- Open Source Projects
- Creative Videos
- Blog Management
- Testimonials
- Image Management
- Session Management
- API Response Caching

---

# Technology Stack

| Category | Technology |
|----------|------------|
| Language | Go |
| Framework | Gin |
| Database | MongoDB |
| Cache | Redis |
| Authentication | JWT + bcrypt |
| Media Storage | Cloudinary |
| Containerization | Docker |
| Version Control | Git |

---

# Project Structure

```text
zeroaxiis-backend/
│
├── cmd/
│   └── api/
│
├── docs/
│   ├── README.md
│   ├── architecture.md
│   ├── api.md
│   ├── models.md
│   ├── development.md
│   └── decisions.md
│
├── internal/
│   ├── config/
│   ├── database/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── routes/
│   ├── services/
│   └── utils/
│
├── .dockerignore
├── .env
├── .env.sample
├── .gitignore
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

---

# Documentation

| File | Purpose |
|------|---------|
| architecture.md | Overall backend architecture, request flow, authentication flow, caching strategy, and infrastructure. |
| models.md | MongoDB collections, document structure, field descriptions, and relationships. |
| api.md | Complete API specification including endpoints, requests, responses, authentication, and caching. |
| development.md | Development workflow, folder structure, coding conventions, project rules, and implementation guidelines. |
| decisions.md | Architectural decisions and the reasoning behind every major technical choice. |

---

# Documentation Philosophy

Before writing code, every major component should be documented.

The documentation should answer:

- Why does this feature exist?
- How does it work?
- Which layer is responsible?
- Which database collections are involved?
- Is Redis used?
- Does authentication apply?
- What is the expected API response?

This ensures that implementation follows a well-defined design rather than making decisions during development.

---

# Development Order

The backend will be developed in the following order:

1. Project Foundation
2. Authentication
3. Team Module
4. Projects Module
5. Open Source Module
6. Creative Module
7. Blog Module
8. Testimonials Module
9. Cloudinary Integration
10. GitHub Integration
11. YouTube Integration
12. Production Deployment

---

# Project Philosophy

The ZeroAxiis Backend follows these core principles:

- The backend is the single source of truth.
- Public users have read-only access.
- Only authenticated administrators can modify content.
- Redis is used only for session management and API response caching.
- MongoDB stores persistent business data.
- Cloudinary stores media assets.
- External services are accessed only through the backend.
- Business logic belongs in the service layer.
- Database operations belong in the repository layer.
- Handlers should only process HTTP requests and responses.

---

# Keeping Documentation Updated

Whenever a new feature is introduced:

- Update the relevant documentation before or alongside implementation.
- Keep API specifications synchronized with the code.
- Record important architectural decisions in `decisions.md`.
- Keep models and request/response examples up to date.

The documentation should always reflect the current state of the project.