PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS notifications (
    notifId TEXT NOT NULL UNIQUE,
    recieverId TEXT NOT NULL,
    senderId TEXT NOT NULL,
    senderFullName TEXT NOT NULL,
	seen TEXT NOT NULL,
	notifType TEXT NOT NULL,
	notifStatus TEXT NOT NULL,     
    groupId TEXT,
    groupName TEXT,
    eventId TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (recieverId) REFERENCES users(userID),
    FOREIGN KEY (senderId) REFERENCES users(userID)
);