CREATE TABLE IF NOT EXISTS dining_tables (
    id uuid primary key,
    table_number integer not null unique,
    capacity integer not null,
    table_status varchar not null,
    created_at timestamp without time zone default now() not null,
    updated_at timestamp without time zone default now() not null
);