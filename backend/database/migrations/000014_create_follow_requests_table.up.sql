PRAGMA  foreign_keys=ON;


CREATE TABLE IF NOT EXISTS follow_requests (
    userID TEXT NOT NULL,               
    requestorID TEXT NOT NULL,         
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (userID, requestorID),
    FOREIGN KEY (userID) REFERENCES users(userID) ON DELETE CASCADE,
    FOREIGN KEY (requestorID) REFERENCES users(userID) ON DELETE CASCADE,
    CHECK (userID != requestorID)
);