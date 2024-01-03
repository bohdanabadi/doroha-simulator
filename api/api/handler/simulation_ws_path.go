package handler

import (
	"github.com/bohdanabadi/Traffic-Simulation/api/broadcast"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleSimulationConnection(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {

		log.Println("Failed to upgrade to websocket:", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		broadcast.SendBroadcast(message)
	}
}
