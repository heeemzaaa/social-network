PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS group_membership (
    groupID TEXT NOT NULL,
    userID TEXT NOT NULL,.
    PRIMARY KEY (groupID, userID),
    FOREIGN KEY (groupID) REFERENCES groups(groupID) ON DELETE CASCADE,
    FOREIGN KEY (userID)  REFERENCES users(userID) ON DELETE CASCADE
);



