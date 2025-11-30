package repository

import (
	"context"

	"github.com/dykethecreator/GoApp/pkg/domain"
)

type ChatRepository interface {
	CreateConversation(ctx context.Context, participantIDs []string, isGroup bool, groupName string) (string, error)
	AddParticipant(ctx context.Context, conversationID, userID string) error

	InsertMessage(ctx context.Context, m *domain.ChatMessage) (*domain.ChatMessage, error)
	ListMessages(ctx context.Context, conversationID string, beforeID string, limit int) ([]*domain.ChatMessage, error)

	ListConversations(ctx context.Context, userID string) ([]*domain.Conversation, error)
}
