ALTER TABLE plates
    ALTER COLUMN price DROP DEFAULT,
    ALTER COLUMN plate_description DROP NOT NULL,
    ALTER COLUMN plate_description SET DATA TYPE text;