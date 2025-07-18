PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS post_access (
    postID TEXT NOT NULL,
    userID TEXT NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (postID) REFERENCES posts(postID) ON DELETE CASCADE,
    FOREIGN KEY (userID) REFERENCES users(userID) ON DELETE CASCADE,
    PRIMARY KEY (postID, userID)
);