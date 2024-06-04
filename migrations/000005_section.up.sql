CREATE TABLE IF NOT EXISTS "section" (
    "name" VARCHAR(255) NOT NULL
);

ALTER TABLE section
    ADD CONSTRAINT section_name_unique UNIQUE (name);
