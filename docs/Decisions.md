# Architecture Decisions

## Overview

This document records the important architectural decisions made during the development of the ZeroAxiis Backend.

Each decision includes the reasoning behind it so future development remains consistent and new contributors can understand why certain approaches were chosen.

---

# Decision 001

## Title

Private Administrator System

## Decision

Public administrator registration is disabled.

Administrator accounts are created manually using the Admin CLI.

## Reason

The ZeroAxiis admin panel is private and intended only for trusted administrators.

Exposing a public registration endpoint would create an unnecessary security risk.

---

# Decision 002

## Title

Handler-Centric Architecture

## Decision

Business logic and database operations are implemented directly inside handlers.

## Reason

The project intentionally avoids unnecessary abstraction.

Keeping logic inside handlers makes the codebase easier to understand while learning Go and is sufficient for the current project size.

---

# Decision 003

## Title

No Service Layer

## Decision

The project does not include a Service layer.

## Reason

The current application is not large enough to justify another abstraction layer.

If the project grows significantly in the future, a Service layer can be introduced without affecting the public API.

---

# Decision 004

## Title

No Repository Layer

## Decision

Database operations communicate directly with MongoDB from handlers.

## Reason

Introducing a Repository layer at this stage would add complexity without providing meaningful benefits.

---

# Decision 005

## Title

Redis Session Management

## Decision

Redis is responsible for administrator sessions.

## Reason

Redis provides automatic expiration through TTL and allows inactive sessions to be removed without additional cleanup logic.

---

# Decision 006

## Title

Redis Response Caching

## Decision

Frequently requested API responses are cached inside Redis.

## Reason

Caching reduces unnecessary MongoDB queries while improving response times.

Whenever content changes, the related cache is invalidated.

---

# Decision 007

## Title

MongoDB Stores Only Persistent Data

## Decision

MongoDB stores only permanent business data.

## Reason

Temporary application state belongs in Redis.

Separating permanent and temporary data keeps responsibilities clear.

---

# Decision 008

## Title

Cloudinary Media Storage

## Decision

Only ZeroAxiis-owned media is uploaded to Cloudinary.

## Reason

Cloudinary should be used only for assets owned by the organization.

Third-party media should remain hosted by its original provider.

---

# Decision 009

## Title

GitHub Repository Metadata

## Decision

GitHub repositories are added by providing only the repository URL.

The backend retrieves repository metadata automatically.

## Reason

This eliminates duplicate manual data entry and ensures repository information remains consistent.

---

# Decision 010

## Title

YouTube Video Metadata

## Decision

Creative videos are added using a YouTube URL.

The backend retrieves video metadata automatically.

## Reason

The administrator only needs to provide a summary.

Video information remains synchronized with YouTube.

---

# Decision 011

## Title

Original YouTube Thumbnail

## Decision

The original YouTube thumbnail URL is stored directly in MongoDB.

## Reason

Downloading and re-uploading thumbnails to Cloudinary would duplicate assets unnecessarily.

---

# Decision 012

## Title

RESTful API Design

## Decision

The backend follows RESTful endpoint conventions.

Examples:

```text
GET

POST

PATCH

DELETE
```

## Reason

REST conventions improve consistency and make the API easier to understand.

---

# Decision 013

## Title

API Versioning

## Decision

All endpoints are prefixed with:

```text
/api/v1
```

## Reason

Versioning allows future breaking changes without affecting existing clients.

---

# Decision 014

## Title

JWT Authentication

## Decision

JWT is used for administrator authentication.

## Reason

JWT provides stateless authentication while Redis manages active sessions and inactivity.

---

# Decision 015

## Title

Administrator Session Timeout

## Decision

Administrator sessions expire after 15 minutes of inactivity.

Every authenticated request refreshes the Redis TTL.

## Reason

This protects the admin panel from unauthorized access when an administrator leaves a session unattended.

---

# Decision 016

## Title

Utility Package

## Decision

Reusable functionality is placed inside the `utils` package.

Examples include:

- JWT
- bcrypt
- Cloudinary
- GitHub API
- YouTube Data API
- Response Helpers

## Reason

Keeping reusable logic together reduces code duplication while avoiding unnecessary architectural complexity.

---

# Decision 017

## Title

Admin CLI

## Decision

Administrator management is performed through a dedicated CLI tool.

The CLI is responsible for:

- Creating administrators
- Resetting passwords
- Future administrator maintenance tasks

## Reason

Administrator management should never be exposed through public HTTP endpoints.

The CLI is executed manually by trusted developers and does not run as part of the web server.

---

# Summary

These decisions define the current architecture of the ZeroAxiis Backend.

Future architectural changes should update this document to maintain consistency and provide historical context for development decisions.