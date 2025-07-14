PRAGMA foreign_keys = ON;



CREATE TABLE IF NOT EXISTS followers(
    userID TEXT NOT NULL, 
    followerID TEXT NOT NULL, 
    FOREIGN KEY (userID)  REFERENCES users(userID),
    FOREIGN KEY  (followerID) REFERENCES users(userID),
    PRIMARY KEY (userID, followerID)
    CHECK (userID != followerID)
);


