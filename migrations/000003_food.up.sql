CREATE TABLE IF NOT EXISTS food (
    guid CHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    mass FLOAT NOT NULL
);

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'food_guid_foreign') THEN
ALTER TABLE food
    ADD CONSTRAINT food_guid_foreign FOREIGN KEY (guid) REFERENCES animal_type (guid) ON DELETE CASCADE;
END IF;
END $$;