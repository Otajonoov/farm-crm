package adapter

import (
	"context"
	"errors"
	"farm-crm/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type animalTypePostgresRepo struct {
	db      *pgxpool.Pool
	factory domain.Factory
}

func NewAnimalTypePostgresRepo(db *pgxpool.Pool) domain.AnimalTypeRepository {
	return &animalTypePostgresRepo{
		db: db,
	}
}

type AnimalType struct {
	guid  string
	types string
	count int
}

func (a animalTypePostgresRepo) Save(animalType *domain.AnimalType) error {
	query := `
			INSERT INTO animal_type(
                        guid, 
                        type, 
                        count
            ) 
			VALUES ($1, $2, $3)
			`

	_, err := a.db.Exec(context.Background(), query,
		animalType.AnimalGuid(),
		animalType.AnimalType(),
		animalType.Count(),
	)
	if err != nil {
		return errors.New("error saving animal type")
	}

	return nil
}

func (a animalTypePostgresRepo) GetAnimalTypeInfo() (animals []domain.AnimalType, err error) {
	query := `
			SELECT at.guid, 
			       at.type,
				   at.count
			FROM animal_type AS at
			`

	rows, err := a.db.Query(context.Background(), query)
	if err != nil {
		return []domain.AnimalType{}, errors.New("error saving animal type")
	}
	defer rows.Close()

	for rows.Next() {
		var animalType AnimalType
		if err := rows.Scan(
			&animalType.guid,
			&animalType.types,
			&animalType.count,
		); err != nil {
			return []domain.AnimalType{}, errors.New("error saving animal type")
		}
		animalTypeFactory := a.factory.ParseToDomainForAnimalType(animalType.guid, animalType.types, animalType.count)
		animals = append(animals, *animalTypeFactory)
	}

	return animals, nil
}

func (a animalTypePostgresRepo) Update(animalType domain.AnimalType) error {
	query := `
			UPDATE animal_type SET
			types = $1,
			count = $2,
			WHERE guid = $3`

	if _, err := a.db.Exec(context.Background(), query, animalType.AnimalType(), animalType.Count(), animalType.AnimalGuid()); err != nil {
		return errors.New("error update animal type")
	}

	return nil
}

func (a animalTypePostgresRepo) Delete(guid string) error {
	query := `DELETE FROM animal_type WHERE guid = $1`

	if _, err := a.db.Exec(context.Background(), query, guid); err != nil {
		return errors.New("error deleting animal type")
	}

	return nil
}
