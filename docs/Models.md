# Database Models

## Overview

The ZeroAxiis Backend uses MongoDB to store all persistent business data.

Each model represents a MongoDB collection and defines the structure of the data stored by the application.

These models are intentionally kept simple to match the current requirements of the ZeroAxiis website. Additional fields can be introduced in future versions without changing the overall architecture.

---

# Admin Model

## Purpose

Stores administrator accounts used to access the private admin panel.

Administrator accounts are created manually using the Admin CLI.

Public registration is disabled.

---

## Collection

```text
admins
```

---

## Fields

| Field | Type | Description |
|--------|------|-------------|
| _id | ObjectId | Unique MongoDB document identifier |
| email | String | Administrator email address |
| password | String | bcrypt hashed password |
| role | String | Administrator role |
| last_login_at | Date | Last successful login |
| created_at | Date | Document creation time |
| updated_at | Date | Last update time |

---

## Example Document

```json
{
    "_id": "6864e0a7b4f3b91d9b3d1234",
    "email": "admin@zeroaxiis.com",
    "password": "$2a$10$...",
    "role": "admin",
    "last_login_at": "2026-07-02T10:00:00Z",
    "created_at": "2026-07-01T08:00:00Z",
    "updated_at": "2026-07-02T10:00:00Z"
}
```

---

# Team Member Model

## Purpose

Stores information displayed on the Team page.

---

## Collection

```text
team_members
```

---

## Fields

| Field | Type | Description |
|--------|------|-------------|
| _id | ObjectId | Unique MongoDB document identifier |
| image_url | String | Cloudinary image URL |
| name | String | Team member name |
| role | String | Position in the company |
| description | String | Short introduction |
| created_at | Date | Document creation time |
| updated_at | Date | Last update time |

---

## Example Document

```json
{
    "_id": "6864e0a7b4f3b91d9b3d1235",
    "image_url": "https://res.cloudinary.com/...",
    "name": "John Doe",
    "role": "Backend Developer",
    "description": "Passionate about scalable backend systems.",
    "created_at": "2026-07-01T08:00:00Z",
    "updated_at": "2026-07-01T08:00:00Z"
}
```

---

# Client Project Model

## Purpose

Stores client projects displayed on the Projects page.

---

## Collection

```text
projects
```

---

## Fields

| Field | Type | Description |
|--------|------|-------------|
| _id | ObjectId | Unique MongoDB document identifier |
| title | String | Project title |
| image_url | String | Cloudinary image URL |
| description | String | Project description |
| client_name | String | Client name |
| project_url | String | Live project URL |
| created_at | Date | Document creation time |
| updated_at | Date | Last update time |

---

## Example Document

```json
{
    "_id": "6864e0a7b4f3b91d9b3d1236",
    "title": "ZeroAxiis Website",
    "image_url": "https://res.cloudinary.com/...",
    "description": "Official company website.",
    "client_name": "ZeroAxiis",
    "project_url": "https://zeroaxiis.vercel.app",
    "created_at": "2026-07-01T08:00:00Z",
    "updated_at": "2026-07-01T08:00:00Z"
}
```

---

# Open Source Project Model

## Purpose

Stores GitHub repositories displayed in the Open Source section.

Repository metadata is fetched automatically from the GitHub API.

---

## Collection

```text
opensource_projects
```

---

## Fields

| Field | Type | Description |
|--------|------|-------------|
| _id | ObjectId | Unique MongoDB document identifier |
| github_url | String | Repository URL |
| repository_name | String | Repository name |
| description | String | Repository description |
| primary_language | String | Main programming language |
| license | String | Repository license |
| stars | Number | GitHub stars |
| forks | Number | GitHub forks |
| last_updated | Date | Last update from GitHub |
| created_at | Date | Document creation time |
| updated_at | Date | Last update time |

---

## Example Document

```json
{
    "_id": "6864e0a7b4f3b91d9b3d1237",
    "github_url": "https://github.com/ZeroAxiis/backend",
    "repository_name": "backend",
    "description": "Official backend of ZeroAxiis.",
    "primary_language": "Go",
    "license": "MIT",
    "stars": 54,
    "forks": 12,
    "last_updated": "2026-07-01T20:30:00Z",
    "created_at": "2026-07-01T08:00:00Z",
    "updated_at": "2026-07-01T08:00:00Z"
}
```

---

# Creative Video Model

## Purpose

Stores videos displayed on the Creative page.

Video metadata is fetched automatically from the YouTube Data API.

---

## Collection

```text
creative_videos
```

---

## Fields

| Field | Type | Description |
|--------|------|-------------|
| _id | ObjectId | Unique MongoDB document identifier |
| youtube_url | String | Original YouTube video URL |
| video_id | String | YouTube video ID |
| title | String | Video title |
| description | String | Video description fetched from YouTube |
| thumbnail_url | String | Original YouTube thumbnail URL |
| summary | String | Administrator-written summary |
| created_at | Date | Document creation time |
| updated_at | Date | Last update time |

---

## Example Document

```json
{
    "_id": "6864e0a7b4f3b91d9b3d1238",
    "youtube_url": "https://youtu.be/abc123",
    "video_id": "abc123",
    "title": "Building ZeroAxiis",
    "description": "Official project walkthrough.",
    "thumbnail_url": "https://i.ytimg.com/...",
    "summary": "Overview of the ZeroAxiis backend architecture.",
    "created_at": "2026-07-01T08:00:00Z",
    "updated_at": "2026-07-01T08:00:00Z"
}
```

---

# Blog Model

## Purpose

Stores blog posts displayed on the Blog page.

---

## Collection

```text
blogs
```

---

## Fields

| Field | Type | Description |
|--------|------|-------------|
| _id | ObjectId | Unique MongoDB document identifier |
| title | String | Blog title |
| image_url | String | Cloudinary image URL |
| author | String | Blog author |
| content | String | Blog content |
| created_at | Date | Document creation time |
| updated_at | Date | Last update time |

---

## Example Document

```json
{
    "_id": "6864e0a7b4f3b91d9b3d1239",
    "title": "Why We Chose Go",
    "image_url": "https://res.cloudinary.com/...",
    "author": "ZeroAxiis",
    "content": "Complete blog content...",
    "created_at": "2026-07-01T08:00:00Z",
    "updated_at": "2026-07-01T08:00:00Z"
}
```

---

# Testimonial Model

## Purpose

Stores client testimonials displayed on the website.

---

## Collection

```text
testimonials
```

---

## Fields

| Field | Type | Description |
|--------|------|-------------|
| _id | ObjectId | Unique MongoDB document identifier |
| client_name | String | Client name |
| company | String | Client company |
| comment | String | Testimonial message |
| created_at | Date | Document creation time |
| updated_at | Date | Last update time |

---

## Example Document

```json
{
    "_id": "6864e0a7b4f3b91d9b3d1240",
    "client_name": "John Doe",
    "company": "ABC Technologies",
    "comment": "Working with ZeroAxiis was an amazing experience.",
    "created_at": "2026-07-01T08:00:00Z",
    "updated_at": "2026-07-01T08:00:00Z"
}
```

---

# Model Relationships

The current application keeps all models independent.

Relationships are managed at the application level rather than through MongoDB references.

Examples:

- Team Members are independent documents.
- Projects are independent documents.
- Open Source Projects are independent documents.
- Creative Videos are independent documents.
- Blogs are independent documents.
- Testimonials are independent documents.

This approach keeps the database simple, flexible, and easy to maintain.

---

# Summary

The ZeroAxiis Backend currently consists of seven MongoDB collections.

- Admins
- Team Members
- Client Projects
- Open Source Projects
- Creative Videos
- Blogs
- Testimonials

These models represent the current requirements of the ZeroAxiis platform and provide a simple foundation that can be extended as the project grows.