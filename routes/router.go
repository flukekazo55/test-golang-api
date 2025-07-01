package routes

import (
	"github.com/gin-gonic/gin"

	// controllers import
	"test-golang-api/controllers"

	// middleware import
	"test-golang-api/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)

	auth := r.Group("/", middleware.AuthMiddleware()) // 🔒 Protected routes
	{
		// Users Routes
		auth.GET("/users", controllers.GetUsers)
		auth.GET("/users/:id", controllers.GetUserByID)
		auth.POST("/users", controllers.PostUser)
		auth.PUT("/users/:id", controllers.UpdateUser)
		auth.DELETE("/users/:id", controllers.DeleteUser)
	}
}
