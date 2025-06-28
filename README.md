# 📝 Golang Blog Service – GoFr Summer of Code Assignment 4

This is a simple blog post REST API built using **Golang**, implementing clean architecture and best practices such as:

- ✅ Three-layer architecture (Handler, Service, Store)
- ✅ Dependency Injection
- ✅ Design Patterns (Factory)
- ✅ SOLID Principles
- ✅ Interfaces for abstraction and testability

---

## 📌 Features

- Create a blog post (`POST /posts`)
- Get all blog posts (`GET /posts`)
- Get a blog post by ID (`GET /posts/{id}`)

---

## 📂 Project Structure (Three-Layer Architecture)

```bash
├── handler/         # Handles HTTP requests/responses
│   └── post_handler.go
├── service/         # Business logic
│   └── post_service.go
├── store/           # Data access layer (MySQL)
│   └── mysql_store.go
├── models/          # Data models
│   └── post.go
├── interfaces/      # Interface definitions for Store abstraction
├── main.go          # Initializes server, routes, DI
├── go.mod / go.sum  # Go module files
