PRAGMA foreign_keys =ON;

CREATE TABLE post_access (
    postID TEXT PRIMARY KEY,
    userID TEXT NOT NULL
);