CREATE TABLE animals (
    guid CHAR(36) PRIMARY KEY NOT NULL,
    animal_type_guid CHAR(36) NOT NULL,
    descriptive VARCHAR(255) NOT NULL,
    weight FLOAT NOT NULL,
    is_ill BOOLEAN DEFAULT FALSE,
    is_fed BOOLEAN DEFAULT FALSE,
    is_watered BOOLEAN DEFAULT FALSE,
    last_feed_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (animal_type_guid) REFERENCES animal_type (guid) ON DELETE CASCADE
);
