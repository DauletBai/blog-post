-- 001_init.sql
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL
);

-- Damp posts
INSERT INTO posts (title, content) VALUES ('Welcome to my blog', 'This is the first post on this blog.');
INSERT INTO posts (title, content) VALUES ('Go is awesome', 'Let me you why Go is one of the best language.');
INSERT INTO posts (title, content) VALUES ('SQLite and Go', 'How to use SQLite in Go applications.');