PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS notifications (
    notifID TEXT NOT NULL UNIQUE,
    recieverID TEXT NOT NULL,
    senderID TEXT NOT NULL,
	seen TEXT NOT NULL,
	notifType TEXT NOT NULL,
	notifState TEXT NOT NULL, 
	content TEXT,
    groupID TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (reciever_Id) REFERENCES users(userID),
    FOREIGN KEY (sender_Id) REFERENCES users(userID)
);