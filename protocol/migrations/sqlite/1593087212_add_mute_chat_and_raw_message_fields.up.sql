ALTER TABLE chats ADD COLUMN muted BOOLEAN DEFAULT FALSE;
ALTER TABLE raw_messages ADD COLUMN skip_encryption BOOLEAN DEFAULT FALSE;
ALTER TABLE raw_messages ADD COLUMN send_push_notification BOOLEAN DEFAULT FALSE;
