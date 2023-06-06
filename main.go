package main

import (
	"github.com/bohdanabadi/Traffic-Simulation/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", handler.HelloWorld)

	// Start server
	r.Run(":8081")
}
