package handler

import (
	"io"
	"log"

	"github.com/dykethecreator/GoApp/internal/auth/middleware"
	"github.com/dykethecreator/GoApp/internal/realtime"
	"github.com/dykethecreator/GoApp/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RealtimeHandler struct {
	proto.UnimplementedRealtimeServiceServer
	hub *realtime.Hub
}

func NewRealtimeHandler(hub *realtime.Hub) *RealtimeHandler {
	return &RealtimeHandler{hub: hub}
}

func (h *RealtimeHandler) Register(s *grpc.Server) {
	proto.RegisterRealtimeServiceServer(s, h)
}

// Connect establishes a bidirectional stream for real-time events
func (h *RealtimeHandler) Connect(stream proto.RealtimeService_ConnectServer) error {
	ctx := stream.Context()

	// Extract user_id from auth interceptor
	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "user_id not found in context")
	}

	log.Printf("[Realtime] User %s connecting...", userID)

	// Register client with hub
	client := h.hub.RegisterClient(userID, stream)
	defer h.hub.UnregisterClient(client)

	// Start write pump in goroutine (sends server events to client)
	go client.WritePump()

	// Read pump (receive client events)
	for {
		clientEvent, err := stream.Recv()
		if err == io.EOF {
			log.Printf("[Realtime] User %s disconnected (EOF)", userID)
			return nil
		}
		if err != nil {
			log.Printf("[Realtime] User %s receive error: %v", userID, err)
			return err
		}

		// Handle client events
		switch e := clientEvent.Event.(type) {
		case *proto.ClientEvent_Ping:
			// Respond with pong via client.send channel
			pong := &proto.ServerEvent{
				Event: &proto.ServerEvent_Pong{
					Pong: &proto.Pong{Timestamp: e.Ping.Timestamp},
				},
			}
			select {
			case client.Send <- pong:
			default:
				log.Printf("[Realtime] User %s: pong buffer full", userID)
			}

		case *proto.ClientEvent_Typing:
			// Broadcast typing indicator to conversation participants
			// TODO: fetch participant IDs from conversation
			log.Printf("[Realtime] User %s typing in conversation %s", userID, e.Typing.ConversationId)
			// h.hub.BroadcastTyping(e.Typing.ConversationId, userID, participantIDs, e.Typing.IsTyping)

		case *proto.ClientEvent_ReadReceipt:
			// Mark message as read
			log.Printf("[Realtime] User %s read message %s", userID, e.ReadReceipt.MessageId)
			// TODO: update message_status table
		}
	}
}
