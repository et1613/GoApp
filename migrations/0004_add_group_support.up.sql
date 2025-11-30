-- Add group conversation support
ALTER TABLE conversations 
ADD COLUMN is_group BOOLEAN NOT NULL DEFAULT FALSE,
ADD COLUMN group_name TEXT;

-- Index for finding group conversations
CREATE INDEX IF NOT EXISTS idx_conversations_is_group ON conversations(is_group);
