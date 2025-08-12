PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS notifications (
    notificationID TEXT NOT NULL UNIQUE,
    senderID  TEXT NOT NULL,
    targetID   TEXT  NOT NULL, 
	seenStatus INTEGER NOT NULL DEFAULT 0 CHECK(seenStatus IN(0 , 1)),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (senderID) REFERENCES users(userID)
);