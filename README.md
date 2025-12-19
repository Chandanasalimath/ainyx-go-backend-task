# Ainyx Go Backend Task

## Overview

This project is a Go-based RESTful backend service built as part of the **Ainyx Go Backend Development Task**. It demonstrates clean backend architecture using Go Fiber, SQLC, and a layered design (handler–service–repository).

The application manages **users with name and date of birth (DOB)** and dynamically calculates the user's **age** at runtime using Go’s `time` package.

---

## Tech Stack

* **Go**
* **GoFiber** – HTTP web framework
* **PostgreSQL** (schema-based, DB connection kept minimal for task scope)
* **SQLC** – type-safe database access
* **go-playground/validator** – request validation
* **Zap Logger** – structured logging (optional integration)

---

## Project Structure

```
cmd/server/main.go     # Application entry point
config/               # Configuration files (if extended)
db/                   # DB schema & migrations
sqlc.yaml             # SQLC configuration
internal/
  handler/             # HTTP handlers
  service/             # Business logic
  repository/          # SQLC-generated DB access
  routes/              # Route registration
```

---

## API Endpoints

### Create User

**POST** `/users`

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

### Get User by ID

**GET** `/users/{id}`

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}
```

### Update User

**PUT** `/users/{id}`

### Delete User

**DELETE** `/users/{id}`

### List Users

**GET** `/users`

---

## Age Calculation

Age is calculated dynamically in the service layer using Go’s `time` package, ensuring accuracy without storing redundant data in the database.

---

## Running the Application

```bash
go mod tidy
go run ./cmd/server
```

The server starts on:

```
http://localhost:3000
```

---

## Notes

* Database connection is intentionally minimal (`nil`) as the focus of this task is **architecture, SQLC usage, and API design**.
* The project demonstrates scalable backend structure suitable for production extensions.

---

## Author

**Chandana Salimath**

GitHub: [https://github.com/Chandanasalimath](https://github.com/Chandanasalimath)
# Ainyx Go Backend Task 
