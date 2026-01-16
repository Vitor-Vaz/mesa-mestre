CREATE TABLE IF NOT EXISTS plates (
    id uuid primary key DEFAULT gen_random_uuid(),
    price float8 not null,
    plate_name varchar not null,
    plate_description text,
    created_at timestamp with time zone default now()
);