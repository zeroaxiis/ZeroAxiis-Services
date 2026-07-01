<div align="center">

# ZeroAxiis Backend

A production-ready backend powering the **ZeroAxiis** website — Built with **Go**, **Gin**, **MongoDB**, **Redis**, **JWT**, **Docker**, and **Cloudinary**.

</div>


## About

ZeroAxiis Backend is the official API responsible for serving dynamic content to the ZeroAxiis website.

The backend follows a layered architecture and is designed with scalability, maintainability, and clean code principles in mind.

The frontend consumes this API to display:

- Team Members
- Client Projects
- Open Source Projects
- Creative Videos
- Blogs
- Testimonials


# Features

### Authentication

- JWT Authentication
- bcrypt Password Hashing
- Redis Session Management
- 15-Minute Inactivity Logout
- Manual Admin Creation (No Public Registration)

### Content Management

- Team Management
- Client Projects
- Open Source Projects
- Creative Videos
- Blogs
- Testimonials

### Storage

- MongoDB
- Redis
- Cloudinary

### Integrations

- GitHub Repository Metadata
- YouTube Video Metadata

### Infrastructure

- Docker
- Docker Compose
- Layered Architecture
- Repository Pattern


# Tech Stack

| Category | Technology |
|-----------|------------|
| Language | Go |
| Framework | Gin |
| Database | MongoDB |
| Cache | Redis |
| Authentication | JWT + bcrypt |
| Image Storage | Cloudinary |
| Containerization | Docker |
| Version Control | Git |


# Project Structure

```text
zeroaxiis-backend/
│
├── cmd/
│   └── api/
│
├── docs/
│
├── internal/
│   ├── config/
│   ├── database/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   ├── routes/
│   ├── services/
│   └── utils/
│
├── .dockerignore
├── .env
├── .gitignore
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
└── README.md
```


# Documentation

Detailed documentation is available inside the **docs/** directory.

| File | Description |
|------|-------------|
| architecture.md | Backend architecture and request flow |
| api.md | API endpoints |
| models.md | MongoDB models |
| development.md | Development guide |
| decisions.md | Architectural decisions |



# Design Principles

This backend follows a few core principles:

- Backend is the single source of truth.
- Public users have read-only access.
- Only authenticated administrators can modify data.
- Redis is used for sessions and caching.
- MongoDB stores persistent business data.
- Cloudinary stores media assets.
- External APIs are handled by the backend.

NOTE: This repository is publicly accessible for transparency, portfolio, educational, and evaluation purposes only not for public use.
