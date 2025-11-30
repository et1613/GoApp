package realtime

import (
	"log"
	"sync"
	"time"

	"github.com/dykethecreator/GoApp/proto"
)

// Hub manages active client connections and broadcasts messages
type Hub struct {
	clients    map[string]*Client // userID -> Client
	broadcast  chan *proto.ServerEvent
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

// Client represents a connected user with their stream
type Client struct {
	UserID string
	Stream proto.RealtimeService_ConnectServer
	Send   chan *proto.ServerEvent
	Hub    *Hub
}

// NewHub creates a new Hub instance
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		broadcast:  make(chan *proto.ServerEvent, 256),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Run starts the hub's main loop
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.UserID] = client
			h.mu.Unlock()
			log.Printf("[Hub] User %s connected (total: %d)", client.UserID, len(h.clients))

			// Broadcast presence update
			h.BroadcastPresence(client.UserID, "online")

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client.UserID]; ok {
				delete(h.clients, client.UserID)
				close(client.Send)
			}
			h.mu.Unlock()
			log.Printf("[Hub] User %s disconnected (total: %d)", client.UserID, len(h.clients))

			// Broadcast presence update
			h.BroadcastPresence(client.UserID, "offline")

		case event := <-h.broadcast:
			h.mu.RLock()
			for _, client := range h.clients {
				select {
				case client.Send <- event:
				default:
					// Client send buffer full, skip
					log.Printf("[Hub] Warning: Client %s send buffer full, dropping message", client.UserID)
				}
			}
			h.mu.RUnlock()
		}
	}
}

// RegisterClient adds a new client connection
func (h *Hub) RegisterClient(userID string, stream proto.RealtimeService_ConnectServer) *Client {
	client := &Client{
		UserID: userID,
		Stream: stream,
		Send:   make(chan *proto.ServerEvent, 256),
		Hub:    h,
	}
	h.register <- client
	return client
}

// UnregisterClient removes a client connection
func (h *Hub) UnregisterClient(client *Client) {
	h.unregister <- client
}

// BroadcastMessage sends a new message event to all clients in a conversation
func (h *Hub) BroadcastMessage(conversationID string, participantIDs []string, msg *proto.NewMessage) {
	event := &proto.ServerEvent{
		Event: &proto.ServerEvent_NewMessage{NewMessage: msg},
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	log.Printf("[Hub] Broadcasting message %s to conversation %s (%d participants)", msg.MessageId[:8], conversationID[:8], len(participantIDs))

	sent := 0
	for _, uid := range participantIDs {
		if client, ok := h.clients[uid]; ok {
			select {
			case client.Send <- event:
				sent++
				log.Printf("[Hub] ✅ Sent to client %s", uid[:8])
			default:
				log.Printf("[Hub] ⚠️  Client %s send buffer full for message %s", uid[:8], msg.MessageId[:8])
			}
		} else {
			log.Printf("[Hub] ⚠️  Client %s not connected", uid[:8])
		}
	}
	log.Printf("[Hub] Message sent to %d/%d connected clients", sent, len(participantIDs))
}

// BroadcastTyping sends typing indicator to conversation participants
func (h *Hub) BroadcastTyping(conversationID, userID string, participantIDs []string, isTyping bool) {
	event := &proto.ServerEvent{
		Event: &proto.ServerEvent_Typing{
			Typing: &proto.TypingIndicator{
				ConversationId: conversationID,
				UserId:         userID,
				IsTyping:       isTyping,
			},
		},
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, uid := range participantIDs {
		if uid == userID {
			continue // Don't send typing to the typer
		}
		if client, ok := h.clients[uid]; ok {
			select {
			case client.Send <- event:
			default:
			}
		}
	}
}

// BroadcastPresence sends online/offline status to all connected clients
func (h *Hub) BroadcastPresence(userID, status string) {
	event := &proto.ServerEvent{
		Event: &proto.ServerEvent_Presence{
			Presence: &proto.PresenceUpdate{
				UserId:   userID,
				Status:   status,
				LastSeen: time.Now().Format(time.RFC3339),
			},
		},
	}
	h.broadcast <- event
}

// WritePump sends queued messages to the client stream
func (c *Client) WritePump() {
	for event := range c.Send {
		if err := c.Stream.Send(event); err != nil {
			log.Printf("[Client %s] Send error: %v", c.UserID, err)
			return
		}
	}
}
