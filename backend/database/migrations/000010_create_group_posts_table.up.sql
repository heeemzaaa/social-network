PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS group_posts (
    postID TEXT NOT NULL,
    groupID TEXT NOT NULL,
    userID TEXT NOT NULL,
    content TEXT NOT NULL, 
    FOREIGN KEY (groupID) REFERENCES groups(groupID),
    FOREIGN KEY (userID)  REFERENCES users(userID),
    PRIMARY KEY (postID, groupID, userID)
);

