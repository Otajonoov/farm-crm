CREATE TABLE IF NOT EXISTS warehouse(
    guid CHAR(36) PRIMARY KEY NOT NULL,
    section_name CHAR(36) NOT NULL,
    food_guid CHAR(36) NOT NULL,
    drug_guid BIGINT NOT NULL,
);

ALTER TABLE warehouse
    ADD CONSTRAINT warehouse_food_guid_foreign FOREIGN KEY(food_guid) REFERENCES food (guid) ON DELETE CASCADE;

ALTER TABLE warehouse
    ADD CONSTRAINT warehouse_drug_guid_foreign FOREIGN KEY(drug_guid) REFERENCES drugs (guid) ON DELETE CASCADE;

ALTER TABLE warehouse
    ADD CONSTRAINT warehouse_section_name_foreign FOREIGN KEY(section_name) REFERENCES section (name) ON DELETE CASCADE;
