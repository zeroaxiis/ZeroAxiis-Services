# ZeroAxiis Backend Architecture

## Overview

The ZeroAxiis Backend is the official API responsible for powering the ZeroAxiis website.

It serves as the single source of truth for all dynamic content displayed on the website while managing authentication, session management, API caching, media uploads, and external service integrations.

The backend follows a simple layered architecture that keeps the project easy to understand, maintain, and extend while avoiding unnecessary complexity.

---

# High-Level Architecture

```text
                         Client (Frontend)
                                │
                                ▼
                        HTTP Request
                                │
                                ▼
                            Routes
                                │
                                ▼
                          Middleware
                                │
                                ▼
                            Handler
        ┌──────────────┬──────────────┬──────────────┐
        ▼              ▼              ▼
    MongoDB         Redis          Utils
                                      │
              ┌──────────────┬────────┼────────┬──────────────┐
              ▼              ▼        ▼        ▼              ▼
            JWT          bcrypt   Cloudinary  GitHub    YouTube Data API
```

---

# Architecture Philosophy

The backend is designed around a few core principles.

- Backend is the single source of truth.
- Frontend is responsible only for presentation.
- Business logic is handled directly inside handlers.
- Database operations are performed directly from handlers.
- Shared reusable functionality belongs inside the `utils` package.
- Redis is used only for Session Management and API Response Caching.
- MongoDB stores only permanent business data.
- Cloudinary stores only ZeroAxiis-owned media.
- External services are accessed only through utility functions.
- Every package has a single responsibility.

---

# Package Responsibilities

## Routes

Routes define the public API exposed by the backend.

Responsibilities:

- Register endpoints.
- Apply middleware.
- Forward requests to handlers.

Routes should not contain business logic.

---

## Middleware

Middleware executes before requests reach handlers.

Responsibilities:

- Authentication
- Authorization
- Logging
- Request validation (basic)
- Rate limiting (future)

Middleware should never contain business logic.

---

## Handlers

Handlers receive HTTP requests.

Responsibilities:

- Read request parameters.
- Parse request body.
- Execute application logic.
- Read and write MongoDB data.
- Read and write Redis data.
- Call utility functions whenever reusable functionality or external services are required.
- Return HTTP responses.

Handlers should remain organized and should never duplicate reusable logic that belongs inside the `utils` package.

---

## Database

The database package manages database connections.

Responsibilities:

- MongoDB Connection
- Redis Connection

The database package should never contain business logic.

---

## Models

Models define the structure of MongoDB documents.

Responsibilities:

- Represent MongoDB collections.
- Keep document structures consistent.
- Share common data structures across the application.

Models should never contain business logic.

---

## Utils

The utils package contains reusable helper functions shared across the backend.

Responsibilities:

- JWT generation and validation.
- Password hashing using bcrypt.
- Cloudinary helper functions.
- GitHub API helper functions.
- YouTube Data API helper functions.
- Response helpers.
- Common utility functions.

Utility functions should remain reusable and independent from business logic.

---

# Database Architecture

MongoDB stores all persistent business data.

Collections include:

- Admins
- Team Members
- Client Projects
- Open Source Projects
- Creative Videos
- Blogs
- Testimonials

MongoDB stores:

- Application Data
- Business Metadata
- Cloudinary URLs
- GitHub Repository Metadata
- YouTube Video Metadata

MongoDB never stores:

- Active Sessions
- API Cache
- Temporary Authentication State

---

# Redis Architecture

Redis has two responsibilities.

## Session Management

Stores active administrator sessions.

Each session has a 15-minute inactivity timeout.

Every authenticated request refreshes the session TTL.

If a session expires, Redis automatically removes it.

The next authenticated request returns **401 Unauthorized**, requiring the administrator to log in again.

---

## Response Caching

Frequently accessed API responses are cached.

Examples:

```text
team:list

projects:list

opensource:list

creative:list

blog:list

testimonials:list

home
```

Whenever content changes, the corresponding cache is invalidated.

Redis never stores permanent business data.

---

# Authentication Flow

```text
Administrator Login
        │
        ▼
Verify Credentials
        │
        ▼
MongoDB
        │
        ▼
bcrypt
(Password Verification)
        │
        ▼
Generate JWT
        │
        ▼
Create Redis Session
        │
        ▼
Return Access Token
```

For every protected request:

```text
Client Request
      │
      ▼
Verify JWT
      │
      ▼
Check Redis Session
      │
      ▼
Refresh Session TTL
      │
      ▼
Authorized
```

Only `last_login_at` is stored in MongoDB.

Current session activity is managed entirely by Redis.

---

# Image Upload Flow

```text
Administrator
      │
      ▼
Backend
      │
      ▼
Cloudinary
      │
      ▼
Image URL
      │
      ▼
MongoDB
```

Only image URLs are stored inside MongoDB.

Cloudinary stores only media uploaded and owned by ZeroAxiis.

Examples include:

- Team Profile Pictures
- Project Thumbnails
- Blog Cover Images

External assets are never uploaded to Cloudinary.

---
# GitHub Integration

The backend communicates directly with the GitHub API through utility functions.

When an administrator submits a repository URL, the backend automatically retrieves:

- Repository Name
- Description
- Primary Language
- License
- Stars
- Fork Count
- Last Updated

The backend stores only repository metadata inside MongoDB.

Repository owner information is not stored because every repository belongs to the ZeroAxiis GitHub organization.

GitHub images are never uploaded to Cloudinary.

The frontend never communicates directly with GitHub.

---

# YouTube Integration

The backend communicates directly with the YouTube Data API through utility functions.

When an administrator submits a YouTube video URL, the backend automatically retrieves:

- Video ID
- Title
- Description
- Thumbnail URL

The administrator only provides:

- YouTube URL
- Summary

The original YouTube thumbnail URL is stored directly inside MongoDB.

The thumbnail is never uploaded to Cloudinary.

The frontend never communicates directly with the YouTube Data API.

---

# Session Management

The backend uses Redis-based sessions.

Session lifecycle:

```text
Login

↓

Redis Session Created

↓

15 Minute TTL

↓

User Activity

↓

TTL Refreshed

↓

No Activity

↓

Redis Removes Session

↓

Next Protected Request

↓

401 Unauthorized
```

Only `last_login_at` is stored in MongoDB.

Current session activity is managed entirely by Redis.

---

# Content Flow

```text
Administrator

↓

Routes

↓

Middleware

↓

Handler

↓

MongoDB

↓

Delete Related Redis Cache

↓

Next Request

↓

Fresh Data Cached Again
```

Whenever data is created, updated or deleted, the corresponding Redis cache is removed.

The next GET request retrieves fresh data from MongoDB and stores it back into Redis for future requests.

---

# Error Handling

Every API should return a consistent response structure.

Example:

```json
{
    "success": true,
    "message": "Team member created successfully.",
    "data": {}
}
```

Errors should follow the same response format.

Example:

```json
{
    "success": false,
    "message": "Team member not found.",
    "error": {}
}
```

---

# Future Architecture

Planned improvements:

- Role-Based Access Control (RBAC)
- Audit Logs
- Search
- Analytics
- Background Jobs
- Email Notifications
- Metrics & Monitoring

---

# Summary

The ZeroAxiis Backend follows a simple handler-centric architecture where every package has a single responsibility.

Handlers manage application logic and communicate directly with MongoDB and Redis while using utility functions for reusable functionality and external service integrations.

MongoDB stores permanent business data.

Redis manages sessions and API response caching.

Cloudinary stores only ZeroAxiis-owned media.

GitHub and YouTube remain the source of truth for their own assets.

This separation keeps the codebase maintainable, scalable, and easy to extend as the project grows.