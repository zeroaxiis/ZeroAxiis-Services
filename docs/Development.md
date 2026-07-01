# Development Guide

## Overview

This document explains how to set up, run, and contribute to the ZeroAxiis Backend during development.

The project is containerized using Docker and follows a simple, handler-centric architecture built with Go and Gin.

---

# Prerequisites

Before running the project, make sure the following tools are installed.

- Go 1.25+
- Docker
- Docker Compose
- Git

---

# Clone Repository

```bash
git clone <repository-url>

cd zeroaxiis-backend
```

---

# Environment Variables

Create a `.env` file in the project root.

Example:

```env
PORT=8080

MONGO_URI=

REDIS_URI=

JWT_SECRET=

CLOUDINARY_API_KEY=

CLOUDINARY_API_SECRET=
```

Never commit the `.env` file to GitHub.

---

# Running the Project

Build and start every service.

```bash
docker compose up --build
```

or

```bash
make run
```

---

# Available Make Commands

## Run

```bash
make run
```

Builds and starts every container.

---

## Start

```bash
make up
```

Starts existing containers.

---

## Stop

```bash
make down
```

Stops every container.

---

## Restart

```bash
make restart
```

Rebuilds and restarts every container.

---

## Logs

```bash
make logs
```

Displays container logs.

---

## Clean

```bash
make clean
```

Stops containers and removes volumes.

---

# Project Structure

```text
zeroaxiis-backend/

├── cmd/
│   ├── api/
│   └── admin/
│
├── config/
├── docs/
├── internal/
│   ├── config/
│   ├── database/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── routes/
│   └── utils/
│
├── Dockerfile
├── docker-compose.yaml
├── Makefile
├── go.mod
└── go.sum
```

---

# Development Workflow

Every new feature should follow the same workflow.

1. Create or update the model if required.
2. Create the handler.
3. Register the route.
4. Test the endpoint.
5. Verify Redis caching (if applicable).
6. Verify MongoDB changes.
7. Update documentation if the API changes.

---

# Administrator Management

Administrator accounts are not created through public API endpoints.

Use the Admin CLI for administrative tasks such as:

- Create Administrator
- Reset Password
- Manage Administrator Accounts

This ensures administrator management remains private and inaccessible from the public internet.

---

# Docker Services

The development environment consists of three services.

| Service | Purpose |
|----------|---------|
| API | Go + Gin Backend |
| MongoDB | Persistent Database |
| Redis | Session Management & API Cache |

---

# Branch Strategy

Recommended workflow:

- `main` → Production-ready code.
- `development` → Active development.
- Feature branches → Individual features or fixes.

Example:

```text
feature/team-api

feature/blog-api

fix/login
```

---

# Development Principles

- Keep handlers small and focused.
- Reuse common logic through the `utils` package.
- Avoid duplicated code.
- Keep MongoDB models simple.
- Never expose private administrator functionality publicly.
- Follow consistent API response formats.
- Keep documentation updated whenever architecture or APIs change.

---

# Summary

The ZeroAxiis Backend is designed to provide a simple and predictable development experience.

Docker provides a consistent environment, Make simplifies common tasks, and the handler-centric architecture keeps the project approachable while remaining scalable for future growth.