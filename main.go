package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hi!, Keep your Gin up")
	})

	// Start api
	r.Run(":8081")
}
