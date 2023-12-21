package broadcast

import (
	"log"
	"sync/atomic"
)

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
	//If Hub running and listening to channel so no deadlocks occur, represented via int because atomic limitation
	isRunning int32
}

var H = Hub{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
	isRunning:  0,
}

func (h *Hub) IsHubRunning() bool {
	return atomic.LoadInt32(&h.isRunning) == 1
}
func (h *Hub) setIsRunning(running bool) {
	var val int32
	if running {
		val = 1
	}
	atomic.StoreInt32(&h.isRunning, val)
}

// Run starts the hub to accept various requests.
func (h *Hub) Run() {
	h.setIsRunning(true)
	defer h.setIsRunning(false)
	for {
		select {
		case client := <-h.Register:
			log.Println("Registering client")
			h.Clients[client] = true
		case client := <-h.Unregister:
			log.Println("Unregistering client")
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			if len(h.Clients) == 0 {
				log.Println("No clients to broadcast to")
				continue
			}
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					log.Println("Dropping message, clients is not able to keep up")
				}
			}
		}
	}
}

// SendBroadcast sends messages to the broadcast channel.
func SendBroadcast(message []byte) {
	if H.IsHubRunning() {
		H.Broadcast <- message
	}
}

// RegisterClient new client.
func RegisterClient(client *Client) {
	if H.IsHubRunning() {
		H.Register <- client
	}
}

// RegisterClient new client.
func UnregisterClient(client *Client) {
	if H.IsHubRunning() {
		H.Unregister <- client
	}
}
