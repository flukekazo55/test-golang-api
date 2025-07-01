# 🧪 Go + Gin REST API Example

This is a simple REST API built with Go using the [Gin Web Framework](https://github.com/gin-gonic/gin).  
It includes basic CRUD for users and JWT-based authentication.

---

# 🚀 Features

- RESTful API for user management
- JWT Authentication (login + middleware)
- Clean project structure with controllers, middleware, models, and routes
- Response format standard:  
  ```json
  {
    "status": 200,
    "message": "OK",
    "data": {...}
  }

---

# 🧱 Project Structure
![image](https://github.com/user-attachments/assets/44b7dd4f-2ba8-41af-8322-078b77de001d)


---

# 🔧 Setup & Run
## 1. Clone the repo
```
git clone https://github.com/yourname/test-gin-api.git
cd test-gin-api
```
## 2. Install dependencies
```
go mod tidy
```
## 3. Run the server
```
go run main.go
```
```
Server will start at: http://localhost:8080
```

# 🧪 Unit Testing
Unit tests are provided for all user-related routes (/users).
## Run tests:
```
go test ./controllers -v
```
