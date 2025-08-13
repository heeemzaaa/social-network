PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS notifications (
    notifId TEXT NOT NULL UNIQUE,
    senderId TEXT NOT NULL,
    targetId TEXT NOT NULL,
	seen BOOLEAN DEFAULT 0 CHECK (seen IN (0, 1)),
	notifType TEXT NOT NULL,
	notifStatus TEXT NOT NULL,
    content TEXT NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (senderId) REFERENCES users(userID)
);