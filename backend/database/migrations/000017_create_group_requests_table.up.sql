PRAGMA foreign_keys =ON;

CREATE TABLE IF NOT EXISTS group_requests(
  requestID TEXT NOT NULL PRIMARY KEY, 
  senderID  TEXT NOT NULL ,
  receiverID  TEXT NOT NULL,
  groupID  TEXT NOT NULL,
  FOREIGN KEY (senderID) REFERENCES users(userID) ON DELETE CASCADE,
  FOREIGN KEY (receiverID)  REFERENCES users(userID) ON DELETE CASCADE,
  FOREIGN KEY (groupID) REFERENCES groups(groupID) ON DELETE CASCADE,
  UNIQUE (senderID,senderID ,groupID),
  typeRequest  TEXT NOT NULL CHECK (typeRequest IN ("join-request", "invitation-request"))
);
