package main

import (
	"github.com/gin-gonic/gin"

	// route import
	"test-golang-api/routes"

	// database import
	"test-golang-api/database"
)

func main() {
	// connect to MongoDB
	database.ConnectDB()

	r := gin.Default()

	// register API routes
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
