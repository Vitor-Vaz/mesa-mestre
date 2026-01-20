-- name: InsertDiningTable :exec
INSERT INTO dining_tables (table_number, capacity, table_status)
VALUES ($1, $2, $3);