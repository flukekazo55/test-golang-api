package main

import (
	"github.com/gin-gonic/gin"

	// route import
	"test-golang-api/routes"
)

func main() {
	r := gin.Default()

	// register API routes
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
