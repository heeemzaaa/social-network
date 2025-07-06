PRAGMA foreign_keys = ON;

CREATE TABLE messages (
  id TEXT PRIMARY KEY,
  sender_id TEXT NOT NULL,
  sender_name TEXT NOT NULL,
  target_id TEXT NOT NULL, 
  type TEXT NOT NULL CHECK(type IN ('private', 'group')),
  content TEXT NOT NULL,
  created_at TEXT NOT NULL
);