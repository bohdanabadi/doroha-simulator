package broadcast

import (
	"github.com/gorilla/websocket"
	"io"
	"log"
	"time"
)

// Client is a middleman which reads and writes messages to the WebSocket connection.
type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

// ListenWrite only starts the WritePump for broadcasting messages to this client.
func (c *Client) ListenWrite() {
	go c.WritePump()
}

// WritePump pumps messages from the hub to the websocket connection.
func (c *Client) WritePump() {
	maxRetries := 5
	retryInterval := time.Duration(500)

	defer func() {
		c.Conn.Close()
		UnregisterClient(c)
	}()
	for {
		select {
		case message := <-c.Send:

			var err error
			var w io.WriteCloser
			for retry := 0; retry < maxRetries; retry++ {
				w, err = c.Conn.NextWriter(websocket.TextMessage)
				if err == nil {
					break
				}
				time.Sleep(retryInterval)
			}
			if err != nil {
				log.Printf("Failed to get writer: %v", err)
				return
			}
			w.Write(message)
			if err := w.Close(); err != nil {
				log.Printf("Failed to close writer: %v", err)
				return
			}
		}
	}
}

// Close cleanly closes the WebSocket connection.
func (c *Client) Close() {
	// Close the send channel which will cause WritePump to close the connection.
	UnregisterClient(c)
	close(c.Send)
}
