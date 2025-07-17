# 🗓️ Golang - Event Booking API

A secure and scalable RESTful API built with Go and Gin to manage events, user registration, and access control. Users can browse, book, create, and manage events, with robust permission handling.

---

## 🚀 Features

- 🔍 Public event browsing
- ✍️ Authenticated event creation, editing, and deletion
- 🔒 JWT-based user authentication
- 📌 Event registration and cancellation tied to user identity
- ⚙️ Gin-powered routing and middleware
- 🗃️ Lightweight storage using SQLite

---

## 🛠️ Technologies

- **Go (Golang)** — main language
- **Gin** — for routing and middleware
- **JWT** — authentication mechanism
- **SQLite** — database engine (requires CGO & MinGW on Windows)

> _Optional future enhancements: GORM, Redis, MongoDB_

---

## 📚 API Endpoints

### 🎉 Event Management

| Method | Endpoint            | Description               | Auth Required | Access Control         |
|--------|---------------------|---------------------------|----------------|------------------------|
| GET    | `/events`           | View all events           | ❌             | Public                 |
| GET    | `/events/{id}`      | View event details        | ❌             | Public                 |
| POST   | `/events`           | Create event              | ✅             | Only logged-in users   |
| PUT    | `/events/{id}`      | Update event              | ✅             | Only creator allowed   |
| DELETE | `/events/{id}`      | Delete event              | ✅             | Only creator allowed   |

### 🙍‍♂️ User Authentication

| Method | Endpoint   | Description         | Auth Required |
|--------|------------|---------------------|---------------|
| POST   | `/signup`  | Register new user   | ❌            |
| POST   | `/login`   | Authenticate user   | ❌            |

> Returns a **JWT token** for authenticated access.

### 📝 Event Registration

| Method | Endpoint                    | Description           | Auth Required |
|--------|-----------------------------|-----------------------|---------------|
| POST   | `/events/{id}/register`     | Register for event    | ✅            |
| DELETE | `/events/{id}/register`     | Cancel registration   | ✅            |

---

## 🔒 JWT Authentication

Include your JWT token in the request header for protected routes:

```http
Authorization: Bearer <your_token_here>