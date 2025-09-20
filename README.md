# Go Login API

A secure, robust, and scalable RESTful API for user authentication and management, built with Go (Golang). This project
demonstrates best practices for building a modern backend service, including JWT-based authentication, password hashing,
and clean architecture.

## 🚀 Features

- **User Registration & Login:** Secure endpoints for creating accounts and logging in.
- **JWT Authentication:** Stateless authentication using JSON Web Tokens for session management.
- **Password Hashing:** Uses `bcrypt` for secure password storage (salted and hashed).
- **Clean Architecture:** Organized in a modular, maintainable structure (handlers, services, repositories).
- **MongoDB:** Persistent data storage using the popular NoSQL Database.
- **Environment Configuration:** Easy configuration management using environment variables.
- **Docker Ready:** Containerized application for easy development and deployment.
- **RESTful Principles:** Follows standard HTTP methods and status codes.

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Web Framework:** [Gin](https://github.com/gin-gonic/gin) (High-performance HTTP web framework)
- **Database:** MongoDB
- **Authentication:** JWT (JSON Web Tokens)
- **Password Hashing:** golang.org/x/crypto/bcrypt
- **Containerization:** Docker & Docker Compose
- **Configuration:** Godotenv

## 📦 Prerequisites

Before you begin, ensure you have the following installed on your machine:

- Go (version 1.19 or higher) - [Download here](https://go.dev/dl/)
- Docker and Docker Compose - [Download here](https://www.docker.com/get-started/)

## ⚙️ Installation & Setup

1. Clone the repository

```bash
git clone https://github.com/leandrowiemesfilho/go-login-api.git
cd go-login-api
```

2. **Set up environment variables**

Copy the example environment file and update the values for your local setup.

```text
MONGO_URI=mongodb://localhost:27017
MONGO_DB=users
MONGO_USERNAME=admin
MONGO_PASSWORD=admin
JWT_SECRET=your_secret_goes_here
JWT_EXPIRATION=36000
```

4. **Run with Docker Compose (Recommended)**

The easiest way to get the database and API running is with Docker Compose.

```bash
docker-compose up
```

This command will:

- Build a Docker image for the Go API.
- Start a PostgreSQL container.
- Run database migrations automatically (if configured).
- Start the API server on the port specified in your .env file (default: 8080).

5. **Run without Docker**

If you have a local MongoDB instance, you can run the application directly.

- Ensure your MongoDB server is running and a database exists.
- Update the `.env` file with your local DB connection string (`DB_HOST=localhost`).
- Start the application:

```bash
go run main.go
```

The API server will be running at `http://localhost:8080`.

## 🚀 API Endpoints

| Method | Endpoint      | Description                           | 	Auth Required |
|:-------|:--------------|:--------------------------------------|:---------------|
| POST   | `/api/signup` | Register a new user                   | 	No            |
| POST   | `/api/login`  | Authenticate a user and receive a JWT | 	No            |
