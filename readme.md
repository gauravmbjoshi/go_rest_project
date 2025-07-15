# 🗓️ Golang - Event Booking API

A robust RESTful API built using Go and the Gin framework to manage events, user registration, and secure access control. Users can explore and book events, with permissions in place for creation, editing, and registration.

---

## 🚀 Features

- Public event browsing
- Secure event creation, editing, and deletion
- User authentication via JWT
- Event registration and cancellation tied to user identity
- Structured routing and middleware using Gin
- SQLite support for lightweight storage

---

## 🛠️ Technologies Used

- **Go (Golang)**
- **Gin** for routing and middleware
- **JWT (JSON Web Token)** for authentication
- **SQLite** for database (with CGO and MinGW support on Windows)
- [You can later integrate GORM, Redis, MongoDB, etc.]

---

## 📚 API Endpoints

### 🎉 Event Management

| Method | Endpoint            | Description                     | Auth Required | Notes                |
|--------|---------------------|---------------------------------|---------------|----------------------|
| GET    | `/events`           | List all available events       | ❌            | Public access        |
| GET    | `/events/{id}`      | Get event details               | ❌            | Public access        |
| POST   | `/events`           | Create a new event              | ✅            | Only logged-in users |
| PUT    | `/events/{id}`      | Update an event                 | ✅            | Only creator allowed |
| DELETE | `/events/{id}`      | Delete an event                 | ✅            | Only creator allowed |

### 🙍‍♂️ User Authentication

| Method | Endpoint   | Description       | Auth Required |
|--------|------------|-------------------|---------------|
| POST   | `/signup`  | Register new user | ❌            |
| POST   | `/login`   | Authenticate user | ❌            | Returns JWT token |

### 📝 Event Registration

| Method | Endpoint                    | Description           | Auth Required |
|--------|-----------------------------|-----------------------|---------------|
| POST   | `/events/{id}/register`     | Register for an event | ✅            |
| DELETE | `/events/{id}/register`     | Cancel registration   | ✅            |

---

## 🔒 Authentication

This project uses **JWT** to secure protected endpoints. Tokens must be sent in the `Authorization` header:


---

## 🗃️ SQLite Setup (Windows)

This project uses the `go-sqlite3` driver, which relies on native C code and requires CGO to be enabled on Windows.

### 📦 Requirements

- **CGO Enabled**

Set this in PowerShell before running or building your app:

```powershell
$env:CGO_ENABLED = "1"
go run .