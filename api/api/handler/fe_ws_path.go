package handler

import (
	"github.com/bohdanabadi/Traffic-Simulation/api/broadcast"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleFrontendConnection(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to websocket:", err)
		return
	}
	client := &broadcast.Client{Conn: conn, Send: make(chan []byte, 256)}
	broadcast.RegisterClient(client)

	client.ListenWrite()
}
