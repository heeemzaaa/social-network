PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS group_events (
    eventID TEXT PRIMARY KEY,
    groupID TEXT NOT NULL,
    title  TEXT NOT NULL ,
    description NOT NULL,
    FOREIGN KEY (groupID) REFERENCES groups(groupID)
);
