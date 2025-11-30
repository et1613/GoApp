-- Remove group conversation support
DROP INDEX IF EXISTS idx_conversations_is_group;

ALTER TABLE conversations 
DROP COLUMN IF EXISTS group_name,
DROP COLUMN IF EXISTS is_group;
