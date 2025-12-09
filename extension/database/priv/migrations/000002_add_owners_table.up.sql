CREATE TABLE IF NOT EXISTS owners (
    id uuid primary key DEFAULT gen_random_uuid(),
    name varchar(255) not null,
    email varchar(255) not null unique,
    created_at timestamp with time zone default now()
); 