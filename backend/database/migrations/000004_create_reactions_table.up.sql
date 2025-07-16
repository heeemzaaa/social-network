PRAGMA foreign_keys = ON;


CREATE TABLE IF NOT EXISTS reactions(
  reactionID TEXT PRIMARY KEY, 
  entityType TEXT NOT NULL DEFAULT 'post' CHECK (entityType IN ('post', 'comment')),
  entityID  TEXT NOT NULL,
  reaction INTEGER NOT NULL DEFAULT 0 CHECK (reaction IN (0,1)), 
  userID TEXT NOT NULL,
  FOREIGN KEY (userID) REFERENCES users(userID) ON DELETE CASCADE,
  UNIQUE(userID, entityType, entityID)
);
