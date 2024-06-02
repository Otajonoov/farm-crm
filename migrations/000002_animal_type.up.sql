CREATE TABLE IF NOT EXISTS animal_type (
    guid CHAR(36) PRIMARY KEY NOT NULL,
    "type" VARCHAR(50) NOT NULL,
    "count" INTEGER
);
