CREATE TABLE IF NOT EXISTS orders (
    id uuid primary key DEFAULT gen_random_uuid(),
    order_status varchar not null,
    plate_id uuid references plates(id),
    dining_table_id uuid references dining_tables(id),
    total_price float8 not null,
    created_at timestamp with time zone default now()
);