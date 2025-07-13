PRAGMA foreign_keys = ON;

CREATE TABLE messages (
  id TEXT PRIMARY KEY,
  sender_id TEXT NOT NULL,
  target_id TEXT NOT NULL, 
  type TEXT NOT NULL CHECK(type IN ('private', 'group')),
  readStatus BOOLEAN NOT NULL DEFAULT 0 CHECK (readStatus IN (0, 1)),
  content TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);