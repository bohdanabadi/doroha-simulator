package handler

import "github.com/gin-gonic/gin"

func HelloWorld(c *gin.Context) {
	c.String(200, "Badna Nekhra!")
}
