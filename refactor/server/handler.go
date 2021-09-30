package server

import "github.com/sirupsen/logrus"

// Handler maintains the set of active clients and broadcasts messages to the
// clients.
type Handler struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Handler {
	return &Handler{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Handler) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				logrus.WithField("prefix", "run").Infof("connection closed")
			}
		case <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- []byte("acknowledge"):
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
