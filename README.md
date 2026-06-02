<div align="center">
  <h1>🚀 Gin API Skeleton</h1>
  <p>
    <strong>A clean, production-ready REST API template for Go</strong>
  </p>
  <p>
    Built with Gin • PostgreSQL • pgx • Docker-Ready
  </p>

  [![Go Version](https://img.shields.io/badge/Go-1.26%2B-blue?style=flat-square)](https://golang.org)
  [![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)
  [![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen?style=flat-square)](/)
  
</div>

---

## 📋 Overview

**Gin API Skeleton** is a lightweight, opinionated starter template for building production-grade REST APIs in Go. It provides everything you need to get started with a clean notes service, complete with CRUD operations, PostgreSQL persistence, and a scalable folder structure.

Perfect for:
- Learning Go API development
- Starting a new microservice
- Building a notes or task management backend
- Understanding clean architecture in Go

---

## ✨ Features

- 🔌 **RESTful API** — Complete CRUD operations for notes
- 🗄️ **PostgreSQL** — Robust data persistence with pgx
- ⚡ **Fast & Lightweight** — Built on Gin framework
- 🔄 **Hot Reload** — Live server restart with `air`
- 🔐 **Environment Config** — Secure `.env` configuration
- 📊 **Auto Schema** — Database table creation on startup
- 📡 **Health Check** — Service status monitoring
- 📦 **Clean Architecture** — Separation of concerns across layers

---

## 🛠️ Tech Stack

| Component | Technology | Version |
| --- | --- | --- |
| **Language** | Go | 1.26+ |
| **Web Framework** | Gin | v1.12.0 |
| **Database** | PostgreSQL | 12+ |
| **DB Driver** | pgx | v5.9.2 |
| **Config** | godotenv | v1.5.1 |

---

## 📦 Prerequisites

Before you begin, ensure you have installed:

- **Go** 1.26 or later ([Download](https://golang.org/dl/))
- **PostgreSQL** 12+ ([Download](https://www.postgresql.org/download/) or use [Neon](https://neon.tech/))
- **Git** for version control
- Optional: **Air** for development hot reload

Verify your installations:

```bash
go version
psql --version
```

---

## 🚀 Quick Start

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/sasmithx/gin-api-skeleton.git
cd gin-api-skeleton
```

### 2️⃣ Configure Environment

Create a `.env` file in the project root:

```env
DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
PORT=8080
```

**Example with Neon DB:**
```env
DATABASE_URL=postgres://user:password@ep-xxxxx.us-east-1.aws.neon.tech/dbname?sslmode=require
PORT=8080
```

### 3️⃣ Run the Application

#### Development (with hot reload)

```bash
go install github.com/air-verse/air@latest
air
```

#### Production

```bash
go run .\cmd\api
```

#### Build for Distribution

```bash
go build -o ./bin/api.exe ./cmd/api
./bin/api.exe
```

The API will start on `http://localhost:8080`

---

## 📡 API Documentation

### Base URL

```
http://localhost:8080/api/v1
```

### Endpoints

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` | `/health` | Service health check |
| `POST` | `/notes` | Create a new note |
| `GET` | `/notes` | List all notes |
| `GET` | `/notes/:id` | Get a note by ID |
| `PUT` | `/notes/:id` | Update a note |
| `DELETE` | `/notes/:id` | Delete a note |

### Response Models

#### Note Object

```json
{
  "id": "1",
  "title": "My First Note",
  "content": "This is the content of my first note",
  "pinned": false,
  "created_at": "2026-06-02T09:40:08.265341+05:30",
  "updated_at": "2026-06-02T09:40:08.265341+05:30"
}
```

### Examples

#### ✅ Create a Note

```http
POST /api/v1/notes
Content-Type: application/json

{
  "title": "Important Task",
  "content": "Remember to complete this important task",
  "pinned": true
}
```

**Response (201 Created):**
```json
{
  "id": "1",
  "title": "Important Task",
  "content": "Remember to complete this important task",
  "pinned": true,
  "created_at": "2026-06-02T09:40:08.265341+05:30",
  "updated_at": "2026-06-02T09:40:08.265341+05:30"
}
```

#### ✅ List All Notes

```http
GET /api/v1/notes
```

**Response (200 OK):**
```json
[
  {
    "id": "1",
    "title": "Important Task",
    "content": "Remember to complete this important task",
    "pinned": true,
    "created_at": "2026-06-02T09:40:08.265341+05:30",
    "updated_at": "2026-06-02T09:40:08.265341+05:30"
  }
]
```

#### ✅ Get a Note

```http
GET /api/v1/notes/1
```

**Response (200 OK):**
```json
{
  "id": "1",
  "title": "Important Task",
  "content": "Remember to complete this important task",
  "pinned": true,
  "created_at": "2026-06-02T09:40:08.265341+05:30",
  "updated_at": "2026-06-02T09:40:08.265341+05:30"
}
```

#### ✅ Update a Note

```http
PUT /api/v1/notes/1
Content-Type: application/json

{
  "title": "Updated Title",
  "content": "Updated content",
  "pinned": false
}
```

#### ✅ Delete a Note

```http
DELETE /api/v1/notes/1
```

**Response (204 No Content)**

---

## 🗂️ Project Architecture

### Directory Structure

```
.
├── cmd/
│   └── api/
│       └── main.go              # Application entrypoint
├── internal/
│   ├── config/
│   │   └── config.go            # Environment configuration
│   ├── db/
│   │   ├── postgres.go          # Database connection
│   │   └── schema.go            # Schema initialization
│   ├── notes/
│   │   ├── note_model.go        # Data models
│   │   ├── note_handler.go      # HTTP handlers
│   │   ├── note_repo.go         # Repository / SQL logic
│   │   └── note_route.go        # Route registration
│   └── server/
│       └── router.go            # Gin router setup
├── migrations/
│   └── 001_create_notes_table.sql
├── .env                         # Environment variables
├── .env.example                 # Example environment file
├── .air.toml                    # Air (live reload) config
├── go.mod                       # Go module definition
├── go.sum                       # Dependency checksums
├── request.http                 # Sample API requests
└── README.md                    # This file
```

### Component Overview

#### `cmd/api/main.go`
**Application Bootstrap**
- Loads configuration
- Initializes database connection
- Creates repository and handler
- Starts HTTP server

#### `internal/config/config.go`
**Configuration Management**
- Reads `.env` file
- Validates required variables
- Provides structured config object

#### `internal/db/`
**Database Layer**
- `postgres.go` — Connection pooling and initialization
- `schema.go` — Auto-creates notes table if missing

#### `internal/notes/`
**Notes Feature Implementation**
- `note_model.go` — Note struct and request types
- `note_handler.go` — HTTP request handlers
- `note_repo.go` — SQL queries and database access
- `note_route.go` — Route registration

#### `internal/server/router.go`
**HTTP Router**
- Configures Gin engine
- Mounts `/api/v1` routes
- Registers health check

---

## 🗄️ Database Schema

```sql
CREATE TABLE IF NOT EXISTS notes (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    pinned BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

**Columns:**
- `id` — Auto-incremented primary key
- `title` — Note title (required)
- `content` — Note content (required)
- `pinned` — Pinned status (optional, defaults to false)
- `created_at` — Creation timestamp
- `updated_at` — Last update timestamp

---

## 🔧 Environment Variables

| Variable | Description | Example |
| --- | --- | --- |
| `DATABASE_URL` | PostgreSQL connection string | `postgres://user:pass@localhost:5432/notes` |
| `PORT` | Server port | `8080` |

---

## 📝 Development

### Testing the API

Use the included `request.http` file with VS Code REST Client:

```bash
# Install REST Client extension in VS Code
# Then open request.http and click "Send Request"
```

### Code Quality

Run tests:
```bash
go test ./...
```

Format code:
```bash
go fmt ./...
```

Lint:
```bash
go vet ./...
```

---

## 🚦 Common Issues

### ❌ Database Connection Error

**Problem:** `DATABASE_URL is required`

**Solution:** 
1. Create a `.env` file in the project root
2. Add your PostgreSQL connection string
3. Restart the server

### ❌ Port Already in Use

**Problem:** `bind: Only one usage of each socket address`

**Solution:** 
1. Change `PORT` in `.env` to an available port
2. Or kill the process using port 8080

### ❌ Table Not Found

**Problem:** `relation "notes" does not exist`

**Solution:** 
The app should auto-create the table. If not, restart the server with a fresh `.env` file.

---

## 📚 Learning Resources

- [Gin Documentation](https://github.com/gin-gonic/gin)
- [pgx Documentation](https://github.com/jackc/pgx)
- [Go Best Practices](https://golang.org/doc/effective_go)
- [RESTful API Design](https://restfulapi.net/)

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

---

## 💡 Tips & Best Practices

- Always validate input in handlers before passing to repository
- Use context for cancellation and timeouts
- Keep database queries in the repository layer
- Use prepared statements to prevent SQL injection
- Test error paths, not just happy paths
- Document your API changes in comments

---

<div align="center">
  <p><strong>Happy Coding! 🎉</strong></p>
  <p>Built with ❤️ by the Go </p>
</div>

