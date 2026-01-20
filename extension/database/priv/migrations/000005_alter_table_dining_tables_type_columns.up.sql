ALTER TABLE dining_tables
    ALTER COLUMN id SET DEFAULT gen_random_uuid(),
    ALTER COLUMN table_number SET NOT NULL,
    ALTER COLUMN capacity SET NOT NULL,
    ALTER COLUMN table_status SET NOT NULL;

ALTER TABLE dining_tables
    ADD CONSTRAINT dining_tables_table_number_unique UNIQUE (table_number);