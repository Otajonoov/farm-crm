CREATE TABLE IF NOT EXISTS drugs (
                                     guid CHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    count BIGINT NOT NULL,
    animal_type_guid CHAR(36) NOT NULL
    );

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'drugs_animal_type_guid_foreign') THEN
ALTER TABLE drugs
    ADD CONSTRAINT drugs_animal_type_guid_foreign FOREIGN KEY (animal_type_guid) REFERENCES animal_type (guid) ON DELETE CASCADE;
END IF;
END $$;
