# Auth Service

This is a microservice for handling user authentication. It is built with Go and follows Clean Architecture principles.

## Tech Stack

- **Golang** (Go 1.21+)
- **PostgreSQL**
- **pgx** (PostgreSQL driver)
- **bcrypt** (for password hashing)
- **JWT** (JSON Web Token)
- **golang-migrate** (for DB migrations)
- **net/http**

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/yourusername/auth-service.git
cd auth-service
```

2. Setup PostgreSQL
You can use Docker Compose or install it locally. Make sure it runs on:

Host: localhost

Port: 5432

Database: auth_service

3. Apply Migrations
Install golang-migrate:
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

Then run:
```bash
migrate -path migrations -database "postgres://postgres:reza@localhost:5432/auth_service?sslmode=disable" up
```

4. Run the Service
```bash
go run cmd/main.go
```
Server will start on http://localhost:8080.

### API Endpoints
POST /api/register
Registers a new user.

Request Body:

json
```bash
{
  "name": "Ali",
  "email": "ali@example.com",
  "password": "123456"
}
```
POST /api/login
Authenticates a user.

Returns:

access_token (JWT)

refresh_token (in HTTP-only cookie)

POST /api/refresh-token
Returns new access & refresh tokens if the refresh token cookie is valid.

