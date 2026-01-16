CREATE TABLE IF NOT EXISTS dining_tables (
    id uuid primary key,
    table_number integer,
    capacity integer,
    table_status varchar,
    created_at timestamp without time zone default now() not null,
    updated_at timestamp without time zone default now() not null
);