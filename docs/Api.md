# API Documentation

## Overview

The ZeroAxiis Backend exposes a RESTful API used by the ZeroAxiis website and private administrator panel.

Public users have read-only access to website content.

Only authenticated administrators can create, update, or delete content.

All API endpoints return JSON responses using a consistent response structure.

---

# Base URL

```text
/api/v1
```

---

# Authentication

The backend uses JWT Authentication with Redis-backed sessions.

Protected endpoints require:

- Valid JWT
- Active Redis Session

If either validation fails, the API returns:

```http
401 Unauthorized
```

---

# Authentication APIs

## Login

### Endpoint

```http
POST /api/v1/auth/login
```

### Purpose

Authenticates an administrator and creates a Redis session.

### Request

```json
{
    "email": "admin@zeroaxiis.com",
    "password": "password"
}
```

### Success Response

```json
{
    "success": true,
    "message": "Login successful.",
    "data": {
        "access_token": "<JWT_TOKEN>",
        "expires_in": 900
    }
}
```

### Error Response

```json
{
    "success": false,
    "message": "Invalid email or password."
}
```

---

## Logout

### Endpoint

```http
POST /api/v1/auth/logout
```

### Purpose

Removes the administrator's active Redis session.

### Success Response

```json
{
    "success": true,
    "message": "Logged out successfully."
}
```

---

## Change Password

### Endpoint

```http
PATCH /api/v1/auth/change-password
```

### Purpose

Allows an authenticated administrator to change their password.

### Request

```json
{
    "current_password": "old_password",
    "new_password": "new_password"
}
```

### Success Response

```json
{
    "success": true,
    "message": "Password updated successfully."
}
```

---

# Team APIs

## Get Team Members

```http
GET /api/v1/team
```

Returns every team member.

---

## Create Team Member

```http
POST /api/v1/team
```

Creates a new team member.

---

## Update Team Member

```http
PATCH /api/v1/team/:id
```

Updates an existing team member.

---

## Delete Team Member

```http
DELETE /api/v1/team/:id
```

Deletes a team member.

---

# Client Project APIs

## Get Projects

```http
GET /api/v1/projects
```

Returns all client projects.

---

## Create Project

```http
POST /api/v1/projects
```

Creates a new project.

---

## Update Project

```http
PATCH /api/v1/projects/:id
```

Updates an existing project.

---

## Delete Project

```http
DELETE /api/v1/projects/:id
```

Deletes a project.

---

# Open Source APIs

## Get Open Source Projects

```http
GET /api/v1/opensource
```

Returns all open source repositories.

---

## Create Open Source Project

```http
POST /api/v1/opensource
```

Administrator submits a GitHub repository URL.

The backend automatically fetches repository metadata from the GitHub API.

---

## Update Open Source Project

```http
PATCH /api/v1/opensource/:id
```

Updates an existing repository.

---

## Delete Open Source Project

```http
DELETE /api/v1/opensource/:id
```

Deletes a repository.

---

# Creative APIs

## Get Creative Videos

```http
GET /api/v1/creative
```

Returns every creative video.

---

## Create Creative Video

```http
POST /api/v1/creative
```

Administrator submits a YouTube URL.

The backend automatically retrieves video metadata from the YouTube Data API.

---

## Update Creative Video

```http
PATCH /api/v1/creative/:id
```

Updates a creative video.

---

## Delete Creative Video

```http
DELETE /api/v1/creative/:id
```

Deletes a creative video.

---

# Blog APIs

## Get Blogs

```http
GET /api/v1/blogs
```

Returns all blog posts.

---

## Create Blog

```http
POST /api/v1/blogs
```

Creates a new blog.

---

## Update Blog

```http
PATCH /api/v1/blogs/:id
```

Updates a blog.

---

## Delete Blog

```http
DELETE /api/v1/blogs/:id
```

Deletes a blog.

---

# Testimonial APIs

## Get Testimonials

```http
GET /api/v1/testimonials
```

Returns every testimonial.

---

## Create Testimonial

```http
POST /api/v1/testimonials
```

Creates a testimonial.

---

## Update Testimonial

```http
PATCH /api/v1/testimonials/:id
```

Updates a testimonial.

---

## Delete Testimonial

```http
DELETE /api/v1/testimonials/:id
```

Deletes a testimonial.

---

# Response Format

Every API returns a consistent JSON structure.

## Success Response

```json
{
    "success": true,
    "message": "Operation completed successfully.",
    "data": {}
}
```

---

## Error Response

```json
{
    "success": false,
    "message": "Something went wrong."
}
```

---

# Authentication Flow

```text
Login

↓

JWT Generated

↓

Redis Session Created

↓

Protected Request

↓

JWT Validation

↓

Redis Session Validation

↓

Authorized

↓

Redis TTL Refreshed
```

If the administrator remains inactive for more than 15 minutes:

```text
Redis Session Expires

↓

Next Request

↓

401 Unauthorized

↓

Login Required
```

---

# API Versioning

All endpoints are prefixed with:

```text
/api/v1
```

Future breaking changes will be introduced through new API versions instead of modifying existing endpoints.

Example:

```text
/api/v2
```

---

# Summary

The ZeroAxiis Backend exposes a RESTful API focused on simplicity, consistency, and maintainability.

Public users have read-only access to content.

Authenticated administrators can manage website content through protected endpoints.

Every API follows a consistent response structure, uses JWT authentication with Redis-backed sessions, and follows RESTful conventions for creating, updating, retrieving, and deleting resources. 