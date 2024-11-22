
CREATE TABLE articles (
    id SERIAL PRIMARY KEY, 
    title TEXT, 
    content TEXT NOT NULL, 
    author VARCHAR(100), 
    created_time TIMESTAMP DEFAULT NOW()
); 