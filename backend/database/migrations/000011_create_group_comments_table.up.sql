PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS group_posts_comments (
    commentID  TEXT NOT NULL PRIMARY KEY,
    postID TEXT NOT NULL,
    groupID TEXT NOT NULL,
    userID TEXT NOT NULL,
    content TEXT NOT NULL, 
    imageContent TEXT,
    FOREIGN KEY (groupID) REFERENCES groups(groupID) ON DELETE CASCADE,
    FOREIGN KEY (userID)  REFERENCES users(userID) ON DELETE CASCADE,
    FOREIGN KEY (postID)  REFERENCES group_posts(postID) ON DELETE CASCADE
);



