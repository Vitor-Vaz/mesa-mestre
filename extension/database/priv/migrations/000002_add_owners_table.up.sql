CREATE TABLE IF NOT EXISTS owners (
    id uuid primary key,
    name VARCHAR(255),
    email VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);