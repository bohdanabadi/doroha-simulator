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
	r.RunTLS(":443", "/etc/letsencrypt/live/api.bohdanabadi.com/fullchain.pem", "/etc/letsencrypt/live/api.bohdanabadi.com/privkey.pem")

}
