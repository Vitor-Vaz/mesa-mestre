ALTER TABLE plates
    ALTER COLUMN price SET DEFAULT 0,
    ALTER COLUMN plate_description SET DATA TYPE varchar,
    ALTER COLUMN plate_description SET NOT NULL;