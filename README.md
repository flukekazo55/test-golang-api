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
![Screenshot 2025-07-01 152435](https://github.com/user-attachments/assets/0cb7be71-8ad3-42be-bf0f-af615e19773f)


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
3. Run the server
```
go run main.go
```
```
Server will start at: http://localhost:8080
```
