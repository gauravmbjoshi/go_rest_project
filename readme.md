# 🗓️ Event Booking API

A robust RESTful API built using Go and the Gin framework to manage events, user registration, and secure access control. The system enables users to explore and book events with appropriate permissions in place for actions like event creation, editing, and registration.

---

## 🚀 Features

- Public event browsing
- Secure event creation, editing, and deletion
- User authentication via JWT
- Event registration and cancellation tied to user identity
- Structured routing and middleware with Gin

---

## 🛠️ Technologies Used

- **Go (Golang)**
- **Gin** for HTTP routing and middleware
- **JWT (JSON Web Token)** for authentication
- [Add database and other libraries: e.g., GORM, Redis, MongoDB]

---

## 📚 API Endpoints

### 🎉 Event Management

| Method | Endpoint              | Description                        | Auth Required  | Notes                |
|--------|-----------------------|------------------------------------|----------------|----------------------|
| GET    | `/events`             | List all available events          | ❌             | Public access        |
| GET    | `/events/{id}`        | Get details of a specific event    | ❌             | Public access        |
| POST   | `/events`             | Create a new event                 | ✅             | Only logged-in users |
| PUT    | `/events/{id}`        | Update an existing event           | ✅             | Only creator allowed |
| DELETE | `/events/{id}`        | Delete an event                    | ✅             | Only creator allowed |

### 🙍‍♂️ User Authentication

| Method | Endpoint   | Description       | Auth Required |
|--------|------------|-------------------|----------------|
| POST   | `/signup`  | Register new user | ❌             |
| POST   | `/login`   | Authenticate user | ❌             | Returns JWT token |

### 📝 Event Registration

| Method | Endpoint                        | Description              | Auth Required  |
|--------|---------------------------------|--------------------------|----------------|
| POST   | `/events/{id}/register`         | Register for an event    | ✅             |
| DELETE | `/events/{id}/register`         | Cancel registration      | ✅             |

---

## 🔒 Authentication

This project uses **JWT** to secure endpoints. Gin middleware checks for tokens in the `Authorization` header:

```http
Authorization: Bearer <your_token_here>