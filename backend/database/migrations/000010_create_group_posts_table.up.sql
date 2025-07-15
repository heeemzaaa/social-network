PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS group_posts (
    postID TEXT PRIMARY KEY,
    groupID TEXT NOT NULL,
    userID TEXT NOT NULL,
    content TEXT NOT NULL, 
    imagePath TEXT,
    createdAt  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (groupID) REFERENCES groups(groupID),
    FOREIGN KEY (userID)  REFERENCES users(userID)
);

