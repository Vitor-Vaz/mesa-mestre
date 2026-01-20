ALTER TABLE dining_tables
    ALTER COLUMN id DROP DEFAULT,
    ALTER COLUMN table_number DROP NOT NULL,
    ALTER COLUMN capacity DROP NOT NULL,
    ALTER COLUMN table_status DROP NOT NULL;

ALTER TABLE dining_tables
    DROP CONSTRAINT dining_tables_table_number_unique;