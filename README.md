# Go Blog System

A full-stack blog platform built with Go (Gin) and React, featuring a public-facing blog and a protected admin dashboard.

This project demonstrates clean architecture, RESTful API design, and JWT-based authentication — designed as a portfolio piece for backend/full-stack engineering roles.

## Key Design Decisions

- **Login-only admin system**: No public registration by design. Admin accounts are pre-created, simulating a real-world CMS where user management is handled internally.
- **Separation of concerns**: Public visitors browse content; authenticated admins manage it.
- **Draft/Published workflow**: Posts can be saved as drafts before publishing.

## Features

- Public blog browsing (list & detail views)
- JWT-based admin authentication
- Admin dashboard for CRUD operations on posts
- Post status management (draft / published)
- Role-based route protection (public vs admin)
- RESTful API with consistent error handling
- Vite proxy for seamless frontend-backend development

## Tech Stack

### Backend
| Technology | Purpose |
|------------|---------|
| Go | Application language |
| Gin | HTTP web framework |
| GORM | ORM for PostgreSQL |
| PostgreSQL | Primary database |
| JWT | Stateless authentication |
| bcrypt | Password hashing |

### Frontend
| Technology | Purpose |
|------------|---------|
| React 18 | UI framework |
| TypeScript | Type safety |
| React Router v6 | Client-side routing |
| Fetch API | HTTP client |
| Vite | Build tool & dev server |

### Infrastructure
| Technology | Purpose |
|------------|---------|
| Docker | Containerisation |
| docker-compose | Local environment orchestration |

## Architecture Overview

```
┌─────────────────┐      ┌─────────────────┐      ┌─────────────────┐
│     React       │ ───► │    Gin API      │ ───► │   PostgreSQL    │
│    Frontend     │      │    Backend      │      │    Database     │
└─────────────────┘      └─────────────────┘      └─────────────────┘
        │                        │
        │                        ├── Public routes (no auth)
        │                        └── Protected routes (JWT required)
        │
        ├── Public pages (/, /posts/:id)
        └── Admin pages (/admin/*) → AuthGuard → redirect if no token
```

**Backend layers:**
- `handler` — HTTP request handling and response formatting
- `repository` — Database access (GORM)
- `middleware` — JWT validation, request context injection
- `model` — Domain entities
- `utils` — Password hashing, JWT generation/parsing

**Frontend structure:**
- Route guards redirect unauthenticated users from `/admin/*` to `/login`
- API client auto-attaches JWT token from localStorage
- Centralised error handling for API responses

## API Design

### Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/posts` | List published posts |
| GET | `/api/posts/:id` | Get post by ID |
| POST | `/auth/login` | Admin login |

### Protected Endpoints (JWT Required)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/posts` | List all posts (incl. drafts) |
| POST | `/api/posts` | Create new post |
| PUT | `/api/posts/:id` | Update post (author only) |
| DELETE | `/api/posts/:id` | Delete post (author only) |
| GET | `/api/me` | Get current user info |

## Authentication Design

**Why no registration?**

This system intentionally omits public user registration:

1. **Simulates real-world CMS** — Most content management systems don't allow public sign-ups. Admins are provisioned internally.
2. **Reduced attack surface** — No registration means no spam accounts, no email verification complexity, and simpler security scope.
3. **Focus on core functionality** — For a portfolio project, this keeps the auth flow clean while still demonstrating JWT-based session management.

Admin accounts are created directly in the database or via a seed script.

## Getting Started

### Prerequisites

- Go 1.22+
- Node.js 18+
- Docker & Docker Compose

### Quick Start (Docker)

The fastest way to run the complete system:

```bash
# 1. Clone and configure
cp .env.example .env

# 2. Start all services (PostgreSQL + Backend)
docker compose up -d

# 3. Start frontend (separate terminal)
cd web && npm install && npm run dev
```

**Default admin credentials** (auto-created on first startup):
- Email: `admin@example.com`
- Password: `admin123`

Open `http://localhost:3000` and login to start using the system.

> ⚠️ **Note**: Admin seeding is enabled by default in Docker for demo purposes. Disable `SEED_ADMIN` in production.

---

### Manual Setup (Development)

#### 1. Start Database

```bash
docker compose up postgres -d
```

#### 2. Configure Environment

```bash
cp .env.example .env
```

Edit `.env` for local development:

```env
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=blog_user
DB_PASSWORD=blog_pass
DB_NAME=blog
JWT_SECRET=your-secret-key
SEED_ADMIN=true
```

#### 3. Run Backend

```bash
go run cmd/server/main.go
```

Backend runs at: `http://localhost:8080`

#### 4. Run Frontend

```bash
cd web
npm install
npm run dev
```

Frontend runs at: `http://localhost:3000`

---

### Admin User Setup

**Option A: Automatic Seeding (Recommended for Demo)**

When `SEED_ADMIN=true`, the application automatically creates an admin user on startup:

| Variable | Default | Description |
|----------|---------|-------------|
| `SEED_ADMIN` | `false` | Enable/disable auto-seeding |
| `SEED_ADMIN_USER` | `admin` | Username |
| `SEED_ADMIN_EMAIL` | `admin@example.com` | Email (login) |
| `SEED_ADMIN_PASS` | `admin123` | Password |

The seed is **idempotent** — running multiple times will not create duplicate users.

**Option B: Manual Creation**

Connect to PostgreSQL and insert a user (password is bcrypt-hashed):

```sql
INSERT INTO users (username, email, password_hash, role, created_at, updated_at)
VALUES ('admin', 'admin@example.com', '$2a$10$...', 'admin', NOW(), NOW());
```

## Project Structure

```
go-blog/
├── cmd/server/          # Application entrypoint
├── internal/
│   ├── config/          # Environment configuration
│   ├── database/        # Database connection & migrations
│   ├── handler/         # HTTP handlers
│   ├── middleware/      # Auth middleware
│   ├── model/           # Domain models
│   ├── repository/      # Data access layer
│   └── utils/           # JWT, password utilities
├── web/                 # React frontend
│   ├── src/
│   │   ├── api/         # API client
│   │   ├── components/  # Shared components
│   │   ├── pages/       # Route components
│   │   └── utils/       # Auth helpers
├── docker/              # Dockerfile
└── docker-compose.yml
```

## Future Improvements

- [ ] Pagination for post listings
- [ ] Rich text editor (Markdown or WYSIWYG)
- [ ] Role-based permissions (admin vs editor)
- [ ] Image upload support
- [ ] CI/CD pipeline (GitHub Actions)
- [ ] Unit and integration tests
- [ ] Redis caching for public endpoints

## Author

Built as a portfolio project for backend engineering internship applications.
