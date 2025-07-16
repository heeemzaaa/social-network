PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS group_events (
    eventID TEXT NOT NULL PRIMARY KEY,
    eventCreatorID  TEXT  NOT NULL,
    groupID TEXT NOT NULL,
    title  VARCHAR(100 ) NOT NULL ,
    description VARCHAR (1000) NOT NULL,
    eventTime TEXT NOT NULL,
    createdAt  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (groupID) REFERENCES groups(groupID),
    FOREIGN KEY (eventCreatorID) REFERENCES users(userID)
);
