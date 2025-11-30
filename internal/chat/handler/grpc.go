package handler

import (
	"context"

	"github.com/dykethecreator/GoApp/internal/auth/middleware"
	"github.com/dykethecreator/GoApp/internal/chat/service"
	"github.com/dykethecreator/GoApp/internal/realtime"
	"github.com/dykethecreator/GoApp/pkg/domain"
	"github.com/dykethecreator/GoApp/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChatHandler struct {
	proto.UnimplementedChatServiceServer
	svc *service.ChatService
}

func NewChatHandler(s *service.ChatService) *ChatHandler { return &ChatHandler{svc: s} }

func (h *ChatHandler) Register(s *grpc.Server) { proto.RegisterChatServiceServer(s, h) }

func (h *ChatHandler) CreateConversation(ctx context.Context, req *proto.CreateConversationRequest) (*proto.CreateConversationResponse, error) {
	id, err := h.svc.CreateConversation(ctx, req.ParticipantIds, req.IsGroup, req.GroupName)
	if err != nil {
		return nil, err
	}
	return &proto.CreateConversationResponse{
		Conversation: &proto.Conversation{
			Id:             id,
			ParticipantIds: req.ParticipantIds,
			IsGroup:        req.IsGroup,
			GroupName:      req.GroupName,
		},
	}, nil
}

func (h *ChatHandler) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*proto.SendMessageResponse, error) {
	convID, _ := uuid.Parse(req.ConversationId)

	// Extract sender_id from interceptor context
	senderID := uuid.Nil
	if userIDStr, ok := middleware.UserIDFromContext(ctx); ok {
		senderID, _ = uuid.Parse(userIDStr)
	}

	msg := &domain.ChatMessage{
		ConversationID: convID,
		SenderID:       senderID,
		Content:        req.Content,
	}
	if req.MediaUrl != "" {
		msg.MediaURL = &req.MediaUrl
	}
	if req.MediaType != "" {
		msg.MediaType = &req.MediaType
	}
	if req.ClientMessageId != "" {
		msg.ClientMessageID = &req.ClientMessageId
	}

	m, err := h.svc.SendMessage(ctx, msg)
	if err != nil {
		return nil, err
	}

	// Broadcast to realtime service (best-effort, non-blocking)
	go func() {
		hub := realtime.GetGlobalHub()
		if hub == nil {
			println("[Chat] ❌ Hub is nil!")
			return
		}
		println("[Chat] 📡 Broadcasting message to hub...")

		// Get conversation participants to broadcast
		// Use background context to avoid cancelled context issues
		bgCtx := context.Background()
		conversations, convErr := h.svc.ListConversations(bgCtx, senderID.String())
		if convErr != nil {
			println("[Chat] ❌ Failed to get conversations:", convErr.Error())
			return
		}

		for _, conv := range conversations {
			if conv.ID == convID {
				participantIDs := make([]string, len(conv.ParticipantIDs))
				for i, pid := range conv.ParticipantIDs {
					participantIDs[i] = pid.String()
				}

				println("[Chat] 📨 Broadcasting to", len(participantIDs), "participants")

				// Broadcast new message event
				hub.BroadcastMessage(req.ConversationId, participantIDs, &proto.NewMessage{
					MessageId:      m.ID.String(),
					ConversationId: m.ConversationID.String(),
					SenderId:       m.SenderID.String(),
					Content:        m.Content,
					MediaUrl:       safeStringPtr(m.MediaURL),
					MediaType:      safeStringPtr(m.MediaType),
					CreatedAt:      m.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
				})
				break
			}
		}
	}()

	return &proto.SendMessageResponse{Message: &proto.Message{
		Id:             m.ID.String(),
		ConversationId: m.ConversationID.String(),
		SenderId:       m.SenderID.String(),
		Content:        m.Content,
	}}, nil
}

func (h *ChatHandler) ListMessages(ctx context.Context, req *proto.ListMessagesRequest) (*proto.ListMessagesResponse, error) {
	items, err := h.svc.ListMessages(ctx, req.ConversationId, req.BeforeMessageId, int(req.Limit))
	if err != nil {
		return nil, err
	}
	out := make([]*proto.Message, 0, len(items))
	for _, m := range items {
		out = append(out, &proto.Message{Id: m.ID.String(), ConversationId: m.ConversationID.String(), Content: m.Content})
	}
	return &proto.ListMessagesResponse{Messages: out}, nil
}

func (h *ChatHandler) GetConversations(ctx context.Context, req *proto.GetConversationsRequest) (*proto.GetConversationsResponse, error) {
	// Extract user_id from interceptor context
	userIDStr, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "user_id not found in context")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user_id format")
	}

	// Get conversations for this user
	conversations, err := h.svc.ListConversations(ctx, userID.String())
	if err != nil {
		return nil, err
	}

	// Convert to proto
	out := make([]*proto.Conversation, 0, len(conversations))
	for _, conv := range conversations {
		participantIDs := make([]string, len(conv.ParticipantIDs))
		for i, pid := range conv.ParticipantIDs {
			participantIDs[i] = pid.String()
		}

		groupName := ""
		if conv.GroupName != nil {
			groupName = *conv.GroupName
		}

		out = append(out, &proto.Conversation{
			Id:             conv.ID.String(),
			ParticipantIds: participantIDs,
			IsGroup:        conv.IsGroup,
			GroupName:      groupName,
			CreatedAt:      conv.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &proto.GetConversationsResponse{Conversations: out}, nil
}

// safeStringPtr returns empty string if pointer is nil
func safeStringPtr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
