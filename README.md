## 📌 gotask — Secure Task Manager API in Go

A production-ready, modular REST API built with Go, PostgreSQL, and JWT authentication.

---

### ⚙️ Features

- ✅ Clean architecture (handler → service → repository)
- 🔐 JWT-based authentication (register, login, protect routes)
- 🗃️ PostgreSQL persistence
- 📦 CRUD operations on tasks
- 🔍 Pagination, filtering, and sorting
- 🧼 Input validation using `go-playground/validator`
- 🧱 Structured error handling
- 📁 Environment-based config loading

---

### 📁 Project Structure

```
gotask/
├── cmd/
│   └── server/         # app entrypoint
├── internal/
│   ├── auth/           # register, login, jwt
│   └── task/           # task logic
├── pkg/
│   ├── config/         # env loader
│   ├── db/             # database connection
│   ├── response/       # response writers
│   └── validation/     # form validation
└── .env                # local secrets (not committed)
```

---

### 🔧 Requirements

- Go 1.21+
- PostgreSQL
- Docker (optional)

---

### 🚀 Getting Started

#### 1. Clone the repo

```bash
git clone https://github.com/sudarshanmg/gotask.git
cd gotask
```

#### 2. Create `.env`

```env
PORT=8080
URL=postgres://<user>:<password>@localhost:5432/gotaskdb?sslmode=disable
JWT_SECRET=yourSuperSecretKey
JWT_EXPIRY=15m
```

#### 3. Run Postgres (Docker optional)

```bash
docker run -d --name pg \
  -p 5432:5432 \
  -e POSTGRES_USER=su \
  -e POSTGRES_PASSWORD=secret \
  -e POSTGRES_DB=gotaskdb \
  postgres
```

#### 4. Create tables

```sql
-- run in psql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username TEXT UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE tasks (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  description TEXT,
  completed BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);
```

#### 5. Run the server

```bash
go run cmd/server/main.go
```

---

### 🧪 API Endpoints

#### 🔑 Auth

- `POST /auth/register` – Register user
- `POST /auth/login` – Login, returns JWT token

#### 📌 Tasks (requires JWT)

- `GET /tasks` – List tasks (supports `?page=1&limit=10&sort=created_at&order=desc`)
- `POST /tasks` – Create a task
- `GET /tasks/{id}` – Get task by ID
- `PUT /tasks/{id}` – Update task
- `DELETE /tasks/{id}` – Delete task

> 💡 Pass `Authorization: Bearer <token>` in headers for protected routes.

---

### 🧭 Roadmap

- [x] JWT auth
- [x] Filtering & pagination
- [ ] Swagger docs
- [ ] Dockerfile & Compose setup
- [ ] Unit & integration tests
- [ ] CI/CD via GitHub Actions

---

### 👨‍💻 Author

Made with ❤️ by [@sudarshanmg](https://github.com/sudarshanmg)

---
