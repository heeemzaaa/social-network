PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS users (
    userID TEXT NOT NULL PRIMARY KEY ,
    email VARCHAR(255) NOT NULL UNIQUE,
    firstName VARCHAR(50) NOT NULL,
    lastName VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL, 
    birthDate TEXT  NOT NULL,
    nickname VARCHAR(30) DEFAULT NULL,
    avatarPath TEXT,
    aboutMe TEXT,
    visibility TEXT NOT NULL DEFAULT 'public' CHECK (visibility IN ('public','private')),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE INDEX unique_nickname ON users(nickname) WHERE nickname IS NOT NULL;