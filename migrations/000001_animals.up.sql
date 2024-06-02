CREATE TABLE IF NOT EXISTS animals
(
    guid CHAR(36) PRIMARY KEY NOT NULL,
    animal_type_guid CHAR(36) NOT NULL
    descriptive VARCHAR(255) NOT NULL,
    weight Float NOT NULL,
    is_ill BOOLEAN NOT NULL,
    is_fed BOOLEAN NOT NULL,
    last_feed_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

ALTER TABLE animals
    ADD CONSTRAINT animal_animal_type_foreign FOREIGN KEY(animal_type_guid) REFERENCES animal_type(guid) ON DELETE CASCADE;
