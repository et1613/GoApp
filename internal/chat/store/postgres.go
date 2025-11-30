package store

import (
	"context"
	"database/sql"

	"github.com/dykethecreator/GoApp/internal/chat/repository"
	"github.com/dykethecreator/GoApp/pkg/domain"
	"github.com/google/uuid"
)

type ChatStore struct{ db *sql.DB }

func NewChatStore(db *sql.DB) repository.ChatRepository { return &ChatStore{db: db} }

func (s *ChatStore) CreateConversation(ctx context.Context, participantIDs []string, isGroup bool, groupName string) (string, error) {
	id := uuid.New()
	var grpName *string
	if groupName != "" {
		grpName = &groupName
	}
	if _, err := s.db.ExecContext(ctx,
		`INSERT INTO conversations(id, is_group, group_name, created_at) VALUES($1, $2, $3, NOW())`,
		id, isGroup, grpName); err != nil {
		return "", err
	}
	for _, uid := range participantIDs {
		if _, err := s.db.ExecContext(ctx, `INSERT INTO conversation_participants(conversation_id, user_id, joined_at) VALUES($1,$2,NOW())`, id, uid); err != nil {
			return "", err
		}
	}
	return id.String(), nil
}

func (s *ChatStore) AddParticipant(ctx context.Context, conversationID, userID string) error {
	_, err := s.db.ExecContext(ctx, `INSERT INTO conversation_participants(conversation_id, user_id, joined_at) VALUES($1,$2,NOW()) ON CONFLICT DO NOTHING`, conversationID, userID)
	return err
}

func (s *ChatStore) InsertMessage(ctx context.Context, m *domain.ChatMessage) (*domain.ChatMessage, error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	err := s.db.QueryRowContext(ctx, `
		INSERT INTO messages(id, conversation_id, sender_id, content, media_url, media_type, client_message_id, created_at) 
		VALUES($1,$2,$3,$4,$5,$6,$7,NOW()) 
		RETURNING created_at
	`, m.ID, m.ConversationID, m.SenderID, m.Content, m.MediaURL, m.MediaType, m.ClientMessageID).Scan(&m.CreatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (s *ChatStore) ListMessages(ctx context.Context, conversationID string, beforeID string, limit int) ([]*domain.ChatMessage, error) {
	if limit <= 0 {
		limit = 50
	}
	var rows *sql.Rows
	var err error
	if beforeID != "" {
		rows, err = s.db.QueryContext(ctx, `SELECT id, conversation_id, sender_id, content, media_url, media_type, client_message_id, created_at FROM messages WHERE conversation_id=$1 AND created_at < (SELECT created_at FROM messages WHERE id=$2) ORDER BY created_at DESC LIMIT $3`, conversationID, beforeID, limit)
	} else {
		rows, err = s.db.QueryContext(ctx, `SELECT id, conversation_id, sender_id, content, media_url, media_type, client_message_id, created_at FROM messages WHERE conversation_id=$1 ORDER BY created_at DESC LIMIT $2`, conversationID, limit)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []*domain.ChatMessage{}
	for rows.Next() {
		var m domain.ChatMessage
		var mediaURL, mediaType, clientID sql.NullString
		if err := rows.Scan(&m.ID, &m.ConversationID, &m.SenderID, &m.Content, &mediaURL, &mediaType, &clientID, new(sql.NullTime)); err != nil {
			return nil, err
		}
		if mediaURL.Valid {
			s := mediaURL.String
			m.MediaURL = &s
		}
		if mediaType.Valid {
			s := mediaType.String
			m.MediaType = &s
		}
		if clientID.Valid {
			s := clientID.String
			m.ClientMessageID = &s
		}
		out = append(out, &m)
	}
	return out, nil
}

func (s *ChatStore) ListConversations(ctx context.Context, userID string) ([]*domain.Conversation, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT c.id, c.is_group, c.group_name, c.created_at 
		FROM conversations c 
		JOIN conversation_participants p ON p.conversation_id = c.id 
		WHERE p.user_id = $1 
		ORDER BY c.created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []*domain.Conversation{}
	for rows.Next() {
		var conv domain.Conversation
		var groupName sql.NullString
		if err := rows.Scan(&conv.ID, &conv.IsGroup, &groupName, &conv.CreatedAt); err != nil {
			return nil, err
		}
		if groupName.Valid {
			s := groupName.String
			conv.GroupName = &s
		}

		// Fetch participant IDs for this conversation
		participantRows, pErr := s.db.QueryContext(ctx, `
			SELECT user_id 
			FROM conversation_participants 
			WHERE conversation_id = $1
		`, conv.ID)
		if pErr != nil {
			return nil, pErr
		}

		conv.ParticipantIDs = []uuid.UUID{}
		for participantRows.Next() {
			var pid uuid.UUID
			if err := participantRows.Scan(&pid); err != nil {
				participantRows.Close()
				return nil, err
			}
			conv.ParticipantIDs = append(conv.ParticipantIDs, pid)
		}
		participantRows.Close()

		out = append(out, &conv)
	}
	return out, nil
}
