# go-jwt

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)

A REST API built with [Go Fiber](https://gofiber.io/) demonstrating JWT authentication and protected routes backed by PostgreSQL.

---

## Prerequisites

- [Go](https://go.dev/dl/) 1.21+
- [Docker](https://www.docker.com/) (for the database)
- A `.env` file in the project root

**.env example**
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=root
DB_NAME=go_jwt
JWT_SECRET=your_secret_here
```

---

## Run

**1. Start the database**
```bash
docker compose up -d
```

**2. Install dependencies**
```bash
go mod tidy
```

**3. Start the server**
```bash
go run main.go
```

Server listens on `http://localhost:3000`.

---

## Endpoints

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | `/api/auth/register` | — | Register a new user |
| POST | `/api/auth/login` | — | Login, returns JWT |
| GET | `/api/books` | — | List all books |
| GET | `/api/books/:id` | — | Get a book by ID |
| POST | `/api/books` | JWT | Create a book |
| PUT | `/api/books/:id` | JWT | Update a book |
| DELETE | `/api/books/:id` | JWT | Delete a book |

Pass the token as a `Bearer` token in the `Authorization` header for protected routes.

---

<sub>Learning project inspired by [this Medium article](https://medium.com/code-beyond/go-fiber-jwt-auth-eab51a7e2129). Not intended for production use.</sub>
