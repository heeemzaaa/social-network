PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS groups (
    groupID TEXT PRIMARY KEY,
    groupCreatorID TEXT NOT NULL,
    title  TEXT NOT NULL,
    description  TEXT NOT NULL,
    FOREIGN KEY (groupCreatorID) REFERENCES users(userID)

);