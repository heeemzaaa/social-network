PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS group_requests (
    requestID TEXT NOT NULL UNIQUE,
    receiverID TEXT NOT NULL,
    senderID TEXT NOT NULL,
	groupID TEXT NOT NULL,
	typeRequest TEXT NOT NULL,

    FOREIGN KEY (receiverID) REFERENCES users(userID),
    FOREIGN KEY (senderID) REFERENCES users(userID)
);