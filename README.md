# Go Blog Backend

A production-style blog backend built with Go, designed for learning and internship preparation.

## Tech Stack

- Go (Gin)
- PostgreSQL
- Redis
- GORM
- Docker & Docker Compose
- Viper (Configuration)

## Project Structure

```aiignore
go-blog/
├── cmd/server # Application entry
├── internal # Core business logic
├── pkg # Shared utilities
├── docker-compose.yml
```


## Getting Started

### Prerequisites

- Go 1.22+
- Docker & Docker Compose

### Run with Docker

```bash
docker compose up -d
```

### Run Application
```bash
go run cmd/server/main.go
```

### Visit:
```aiignore
GET http://localhost:8080/health
```