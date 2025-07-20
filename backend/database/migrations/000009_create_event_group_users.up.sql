PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS group_event_users (
    ID  TEXT NOT NULL PRIMARY KEY,
    eventID TEXT NOT NULL,
    groupID TEXT NOT NULL,
    userID TEXT NOT NULL,
    actionChosen INTEGER NOT NULL DEFAULT 0 CHECK (actionChosen IN (-1,0,1)),
    FOREIGN KEY (groupID) REFERENCES groups(groupID),
    FOREIGN KEY (eventID) REFERENCES group_events(eventID),
    FOREIGN KEY (userID)  REFERENCES users(userID),
    UNIQUE (eventID, groupID, userID)
);




