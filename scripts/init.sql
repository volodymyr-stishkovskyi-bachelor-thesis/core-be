CREATE TABLE IF NOT EXISTS user_queries (
    id SERIAL PRIMARY KEY,
    query TEXT NOT NULL,
    response TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS scraped_data (
    id SERIAL PRIMARY KEY,
    source TEXT NOT NULL CHECK (source IN ('linkedin', 'leetcode')),
    url TEXT UNIQUE NOT NULL,
    raw_text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);