package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourproject/internal/handlers"
)

func main() {
	r := gin.Default()

	r.GET("/", handlers.HelloWorld)

	// Start server
	r.Run(":8080")
}
