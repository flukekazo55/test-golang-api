# 🧪 Go + Gin REST API Example

This is a simple REST API built with Go using the [Gin Web Framework](https://github.com/gin-gonic/gin).  
It includes basic CRUD for users and JWT-based authentication.

---

## 🚀 Features

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
