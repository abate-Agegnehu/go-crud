# Go CRUD API

A simple RESTful CRUD API for managing users, built with Go and PostgreSQL. This project is designed for learning purposes to understand the fundamentals of building REST APIs with Go.

## Project Structure

```
go-crud/
├── cmd/main/main.go              # Application entry point
├── internal/
│   ├── database/db.go            # Database connection
│   ├── handler/user_handler.go   # HTTP request handlers
│   ├── models/user.go            # Data models and request structs
│   ├── repository/user_repository.go  # Database queries
│   └── service/user_service.go   # Business logic
├── migrations/
│   └── 001_create_users_table.sql  # Database migration
├── pkg/utils/response.go         # JSON response helpers
├── .env                          # Environment variables
├── go.mod
└── go.sum
```

## Architecture

The project follows a **layered architecture** pattern:

- **Handler** - Handles HTTP requests and responses
- **Service** - Contains business logic
- **Repository** - Manages database operations
- **Models** - Defines data structures

## Prerequisites

- Go 1.21 or higher
- PostgreSQL

## Setup

1. Clone the repository and navigate to the project directory.

2. Create a `.env` file in the root directory with your PostgreSQL credentials:

   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=crud_db
   ```

3. Create the database in PostgreSQL:

   ```sql
   CREATE DATABASE crud_db;
   ```

4. Run the migration to create the users table:

   ```bash
   psql -U <username> -d crud_db -f migrations/001_create_users_table.sql
   ```

5. Run the application:

   ```bash
   go run cmd/main/main.go
   ```

   The server will start on `http://localhost:8080`.

## API Endpoints

| Method | Endpoint              | Description       |
|--------|-----------------------|-------------------|
| POST   | /api/users/create     | Create a new user |
| GET    | /api/users/all        | Get all users     |
| GET    | /api/users/get?id=1   | Get user by ID    |
| PUT    | /api/users/update?id=1| Update a user     |
| DELETE | /api/users/delete?id=1| Delete a user     |

## Request / Response Examples

### Create User

**Request:**

```bash
curl -X POST http://localhost:8080/api/users/create \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john@example.com",
    "phone": "1234567890"
  }'
```

**Response (201 Created):**

```json
{
  "id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "email": "john@example.com",
  "phone": "1234567890",
  "created_at": "2026-08-18T10:00:00Z",
  "updated_at": "2026-08-18T10:00:00Z"
}
```

### Get All Users

```bash
curl http://localhost:8080/api/users/all
```

### Get User by ID

```bash
curl http://localhost:8080/api/users/get?id=1
```

### Update User

```bash
curl -X PUT "http://localhost:8080/api/users/update?id=1" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Jane",
    "last_name": "Doe"
  }'
```

### Delete User

```bash
curl -X DELETE "http://localhost:8080/api/users/delete?id=1"
```

## Dependencies

- [github.com/lib/pq](https://github.com/lib/pq) - PostgreSQL driver
- [github.com/joho/godotenv](https://github.com/joho/godotenv) - Environment variable loader
