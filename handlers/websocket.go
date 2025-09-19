package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// hub maintains active websocket clients and broadcasts messages to them.
type hub struct {
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	clients    map[*websocket.Conn]bool
	broadcast  chan []byte
}

var globalHub *hub

func newHub() *hub {
	return &hub{
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan []byte, 16),
	}
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c] = true
		case c := <-h.unregister:
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				c.Close()
			}
		case msg := <-h.broadcast:
			for c := range h.clients {
				// best-effort write; if it fails, remove the client
				if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
					log.Println("ws write error, removing client:", err)
					delete(h.clients, c)
					c.Close()
				}
			}
		}
	}
}

// BroadcastJSON marshals v to JSON and broadcasts to all connected clients.
func BroadcastJSON(v interface{}) {
	if globalHub == nil {
		return
	}
	b, err := json.Marshal(v)
	if err != nil {
		log.Println("BroadcastJSON marshal error:", err)
		return
	}
	select {
	case globalHub.broadcast <- b:
	default:
		// drop message if buffer full
		log.Println("Broadcast channel full; dropping message")
	}
}

// WebSocketHandler upgrades the connection and registers it with the hub.
func WebSocketHandler(c *gin.Context) {
	if globalHub == nil {
		globalHub = newHub()
		go globalHub.run()
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}

	globalHub.register <- conn

	// keep the connection alive by reading; when read fails, unregister
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			globalHub.unregister <- conn
			break
		}
	}
}
