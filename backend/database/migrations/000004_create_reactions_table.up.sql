PRAGMA foreign_keys = ON;


CREATE TABLE IF NOT EXISTS reactions(
  reactionID TEXT PRIMARY KEY, 
  entityType TEXT NOT NULL DEFAULT 'post' CHECK (entityType IN ('post', 'comment')), -- post or comment ?
  entityID  TEXT NOT NULL, -- comment id / post id
  reaction INTEGER NOT NULL DEFAULT 0 CHECK (reaction IN (0,1)),  -- like 1 / remove like 0 
  userID TEXT NOT NULL, 
  FOREIGN KEY (userID) REFERENCES users(userID) ON DELETE CASCADE,
  UNIQUE(userID, entityType, entityID)
);
