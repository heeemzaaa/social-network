PRAGMA foreign_keys =ON;

CREATE TABLE IF NOT EXISTS group_requests(
  requestID TEXT NOT NULL PRIMARY KEY, 
  senderID  TEXT NOT NULL ,
  receiverID  TEXT NOT NULL,
  typeRequest  TEXT NOT NULL CHECK (typeRequest IN ("join-request", "invitation-request"))
);
