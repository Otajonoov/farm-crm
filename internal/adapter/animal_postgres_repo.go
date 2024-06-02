package adapter

import (
	"context"
	"errors"
	"farm-crm/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type animalPostgresRepo struct {
	db      *pgxpool.Pool
	factory domain.Factory
}

func NewAnimalPostgresRepo(db *pgxpool.Pool) domain.AnimalRepository {
	return &animalPostgresRepo{
		db: db,
	}
}

type Animal struct {
	guid        string
	animalType  string
	descriptive string
	weight      float32
	isIll       bool
	isFed       bool
	lastFedTime time.Time
	createdAt   time.Time
	updatedAt   time.Time
}

func (a animalPostgresRepo) Save(animal *domain.Animal) error {
	query :=
		`INSERT INTO animals(guid, 
                    		animal_type_guid,
                    		descriptive, 
                    		weight, 
                    		is_ill, 
                    		is_fed, 
                    		last_fed_time, 
                    		created_at, 
                    		updated_at
        )
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
                    		
                    		`

	_, err := a.db.Exec(context.Background(), query,
		animal.Guid(),
		animal.AnimalTypeGuid(),
		animal.Descriptive(),
		animal.Weight(),
		animal.IsIll(),
		animal.IsFed(),
		animal.LastFedTime(),
		animal.CreatedAt(),
		animal.UpdatedAt(),
	)
	if err != nil {
		return errors.New("error saving animal")
	}

	return nil
}

func (a animalPostgresRepo) GetByGuid(s string) (*domain.Animal, error) {
	var animal domain.Animal
	query := `
			SELECT a.guid, 
   			       a.animal_type_guid,
                   a.descriptive, 
                   a.weight,
                   a.is_ill,
                   a.is_fed, 
                   a.last_fed_time, 
                   a.created_at, 
                   a.updated_at
			FROM animals a
			WHERE a.guid = $1
		`

	err := a.db.QueryRow(context.Background(), query, s).Scan(
		animal.Guid(),
		animal.AnimalTypeGuid(),
		animal.Descriptive(),
		animal.Weight(),
		animal.IsIll(),
		animal.IsFed(),
		animal.LastFedTime(),
		animal.CreatedAt(),
		animal.UpdatedAt(),
	)
	if err != nil {
		return nil, errors.New("error fetching animal")
	}

	return &animal, nil
}

func (a animalPostgresRepo) GetAll() (animals []domain.Animal, err error) {
	query := `SELECT a.guid,
       				 a.animal_type_guid,
       				 a.descriptive,
       				 a.weight,
       				 a.is_ill,
       				 a.if_fed,
       				 a.last_fed_time,
       				 a.created_at,
       				 a.updated_at
				FROM animals a`

	rows, err := a.db.Query(context.Background(), query)
	if err != nil {
		return []domain.Animal{}, errors.New("error fetching animals")
	}
	defer rows.Close()

	for rows.Next() {
		var animal Animal
		if err := rows.Scan(
			&animal.guid,
			&animal.animalType,
			&animal.descriptive,
			&animal.weight,
			&animal.isIll,
			&animal.isFed,
			&animal.lastFedTime,
			&animal.createdAt,
			&animal.updatedAt,
		); err != nil {
			return nil, errors.New("failed to scan rows")
		}
		animalFactory := a.factory.ParseToDomain(animal.guid, animal.animalType, animal.descriptive, animal.weight, animal.isIll, animal.isFed, animal.lastFedTime, animal.createdAt, animal.updatedAt)
		animals = append(animals, *animalFactory)
	}

	return animals, nil
}

func (a animalPostgresRepo) Update(animal *domain.Animal) error {
	query := `UPDATE animals SET
				description = $1,
				weight = $2,
				is_ill = $3,
				is_fed = $4,
				last_fed_time = $5,
				updated_at = $6`

	_, err := a.db.Exec(context.Background(), query, animal.Descriptive(), animal.Weight(), animal.IsIll(), time.Now(), time.Now())
	if err != nil {
		return errors.New("error updating animal")
	}

	return nil
}

func (a animalPostgresRepo) Delete(guid string) error {
	query := `DELETE FROM animals WHERE guid = $1`
	_, err := a.db.Exec(context.Background(), query, guid)
	if err != nil {
		return errors.New("error deleting animal")
	}

	return nil
}
