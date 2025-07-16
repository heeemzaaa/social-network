PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS groups (
    groupID TEXT NOT NULL PRIMARY KEY,
    groupCreatorID TEXT NOT NULL,
    title  VARCHAR(100) NOT NULL UNIQUE,
    imagePath TEXT,
    description  VARCHAR(1000) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (groupCreatorID) REFERENCES users(userID)
);