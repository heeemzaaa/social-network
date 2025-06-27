PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS group_events (
    eventID TEXT PRIMARY KEY,
    eventCreatorID  TEXT  NOT NULL,
    groupID TEXT NOT NULL,
    title  TEXT NOT NULL ,
    description NOT NULL,
    eventTime TEXT NOT NULL,
    FOREIGN KEY (groupID) REFERENCES groups(groupID),
    FOREIGN KEY (eventCreatorID) REFERENCES users(userID)
);
