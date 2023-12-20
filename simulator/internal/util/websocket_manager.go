package util

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type WebSocketManager struct {
	URL            string
	Connection     *websocket.Conn
	ReconnectDelay time.Duration
	MaxRetries     int
}

func NewWebSocketManager(url string, reconnectDelay time.Duration, maxRetries int) *WebSocketManager {
	return &WebSocketManager{
		URL:            url,
		ReconnectDelay: reconnectDelay,
		MaxRetries:     maxRetries,
	}
}
func (wm *WebSocketManager) Connect() error {
	var err error
	for {
		wm.Connection, _, err = websocket.DefaultDialer.Dial(wm.URL, nil)
		if err != nil {
			log.Printf("Error connecting to WebSocket: %v. Retrying in 15 seconds...", err)
			time.Sleep(15 * time.Second)
			continue
		}
		return nil
	}
}

func (wm *WebSocketManager) Send(data []byte) error {
	if wm.Connection == nil {
		if !wm.HandleReconnection() {
			return fmt.Errorf("could not reconnect")
		}
	}

	if err := wm.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
		wm.Connection = nil // Set connection to nil to trigger reconnection on next send attempt
		return fmt.Errorf("could not send message: %w", err)
	}

	return nil // Message sent successfully
}

func (wm *WebSocketManager) HandleReconnection() bool {
	for retries := 0; retries < wm.MaxRetries; retries++ {
		if err := wm.Connect(); err == nil {
			return true
		}
		time.Sleep(wm.ReconnectDelay)
	}
	fmt.Println("Max retries reached. Could not reconnect.")
	return false
}
func (wm *WebSocketManager) Close() error {
	return wm.Connection.Close()
}
