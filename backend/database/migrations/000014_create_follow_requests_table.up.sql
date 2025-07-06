PRAGMA  foreign_keys=ON;


CREATE TABLE follow_requests (
userID TEXT PRIMARY KEY,
requestorID TEXT NOT NULL,
sent_at TEXT NOT NULL
);