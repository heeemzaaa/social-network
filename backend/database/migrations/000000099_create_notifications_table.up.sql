PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS notifications (
    notif_id TEXT NOT NULL UNIQUE,
    reciever_Id TEXT NOT NULL,
    sender_Id TEXT NOT NULL,
	seen TEXT NOT NULL,
	notif_type TEXT NOT NULL,
	notif_state TEXT NOT NULL, 
	content TEXT,

    FOREIGN KEY (reciever_Id) REFERENCES users(userID),
    FOREIGN KEY (sender_Id) REFERENCES users(userID),
);