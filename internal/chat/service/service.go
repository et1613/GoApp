package service

import (
	"context"

	"github.com/dykethecreator/GoApp/internal/chat/repository"
	"github.com/dykethecreator/GoApp/pkg/domain"
)

type ChatService struct {
	repo repository.ChatRepository
}

func NewChatService(r repository.ChatRepository) *ChatService { return &ChatService{repo: r} }

func (s *ChatService) CreateConversation(ctx context.Context, participantIDs []string, isGroup bool, groupName string) (string, error) {
	return s.repo.CreateConversation(ctx, participantIDs, isGroup, groupName)
}

func (s *ChatService) SendMessage(ctx context.Context, m *domain.ChatMessage) (*domain.ChatMessage, error) {
	return s.repo.InsertMessage(ctx, m)
}

func (s *ChatService) ListMessages(ctx context.Context, conversationID, beforeID string, limit int) ([]*domain.ChatMessage, error) {
	return s.repo.ListMessages(ctx, conversationID, beforeID, limit)
}

func (s *ChatService) ListConversations(ctx context.Context, userID string) ([]*domain.Conversation, error) {
	return s.repo.ListConversations(ctx, userID)
}
