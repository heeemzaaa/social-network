PRAGMA foreign_keys = ON;



CREATE TABLE IF NOT EXISTS followers(
    userID TEXT NOT NULL, 
    followerID TEXT NOT NULL, 
    PRIMARY KEY (userID, followerID),
    FOREIGN KEY (userID)  REFERENCES users(userID),
    FOREIGN KEY  (followerID) REFERENCES users(userID),
    UNIQUE (userID, followerID)
);


