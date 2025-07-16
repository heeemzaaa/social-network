PRAGMA foreign_keys = ON;


CREATE TABLE IF NOT EXISTS posts (
    postID TEXT PRIMARY KEY,
    userID TEXT NOT NULL, 
    content TEXT NOT NULL,
    media TEXT NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    privacy TEXT NOT NULL DEFAULT 'public' CHECK( privacy IN ('public', 'private', 'almost private')),
    FOREIGN KEY (userID) REFERENCES users(userID) ON DELETE CASCADE 
);