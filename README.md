# GO Bookstore REST API

A hands-on learning project for building RESTful APIs with **Go**. This bookstore API is built from scratch using **Gorilla Mux** for routing and **GORM** with **MySQL** as the data layer, covering the full lifecycle of an HTTP request from route registration to database persistence.

---

## Table of Contents

- [Overview](#overview)
- [Learning Goals](#learning-goals)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Database](#database)
- [Response Format](#response-format)

---

## Overview

This project is a hands-on exploration of building REST APIs in Go without a heavy framework. It implements a simple bookstore service with endpoints for creating and listing books, structured following real-world conventions: a dedicated entry point, separated routing/controller/model layers, a centralised database configuration package, and a shared response type.

The goal is to understand how all the pieces of a Go web service fit together — from registering routes with Gorilla Mux, to decoding JSON request bodies, to persisting data through GORM's ORM layer.

---

## Learning Goals

This project was built to practise and solidify the following concepts:

- **Go module system** — understanding `go.mod`, dependency management with `go mod tidy/download`
- **HTTP routing** — using [Gorilla Mux](https://github.com/gorilla/mux) to define method-specific routes
- **Controller pattern** — separating HTTP handler logic from business/data logic
- **ORM with GORM** — defining models, running `AutoMigrate`, and performing CRUD operations
- **MySQL integration** — connecting Go to a relational database via a DSN string
- **JSON encoding/decoding** — reading request bodies with `json.NewDecoder` and writing responses with `json.NewEncoder`
- **Package organisation** — structuring a Go project across multiple packages (`config`, `models`, `controller`, `routes`, `type`)
- **Consistent API responses** — defining a shared `Response` struct as a typed JSON envelope

---

## Tech Stack

| Technology | Purpose |
|---|---|
| [Go 1.20](https://go.dev/) | Core language |
| [Gorilla Mux](https://github.com/gorilla/mux) | HTTP router & dispatcher |
| [GORM](https://gorm.io/) | ORM / database abstraction |
| [MySQL](https://www.mysql.com/) | Relational database |

---

## Project Structure

```
GO-Learn_REST/
├── cmd/
│   └── main.go               # Application entry point
├── pkg/
│   ├── config/
│   │   └── db.go             # Database connection & configuration
│   ├── controller/
│   │   ├── book-controller.go # HTTP handler functions
│   │   └── index.go
│   ├── models/
│   │   └── book-model.go     # Book model & database operations
│   ├── routes/
│   │   ├── book-routes.go    # Route definitions
│   │   └── index.go
│   └── type/
│       └── index.go          # Shared response types
├── go.mod
└── go.sum
```

---

## Prerequisites

- [Go 1.20+](https://go.dev/dl/)
- [MySQL 8.0+](https://dev.mysql.com/downloads/)

---

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/SeanardK/GO-REST_API.git
cd GO-REST_API
```

### 2. Set up the database

Create the MySQL database before running the application:

```sql
CREATE DATABASE `go-bookstore`;
```

> The application connects to `root@127.0.0.1:3306/go-bookstore` by default. Update the DSN in [pkg/config/db.go](pkg/config/db.go) if your credentials differ.

### 3. Install dependencies

```bash
go mod download
```

### 4. Run the server

```bash
go run cmd/main.go
```

The server starts on **`http://localhost:3001`**.

> GORM will automatically migrate the `books` table on startup via `AutoMigrate`.

---

## Database

The application connects to a local MySQL instance using the following default DSN:

```
root@tcp(127.0.0.1:3306)/go-bookstore
```

To use a different user, password, or host, update the `Connect()` function in [pkg/config/db.go](pkg/config/db.go):

```go
gorm.Open(mysql.Open("user:password@tcp(host:port)/go-bookstore"))
```

---

## Response Format

All endpoints return a consistent JSON envelope:

```json
{
  "message": "string",
  "status": 200,
  "data": []
}
```

| Field     | Type      | Description                        |
|-----------|-----------|------------------------------------|
| `message` | `string`  | Human-readable status message      |
| `status`  | `integer` | HTTP-style status code             |
| `data`    | `array`   | Array of book objects              |
