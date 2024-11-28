-- Script to create a database schema for an article.

CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    content TEXT NOT NULL,
    author VARCHAR(255),
    created_time TIMESTAMP DEFAULT NOW()
);    