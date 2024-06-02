package usecase

import (
	"errors"
	"farm-crm/internal/domain"
	"time"
)

type animalUseCase struct {
	BaseUseCase
	animalRepo domain.AnimalRepository
}

func NewAnimalUseCase(animalRepo domain.AnimalRepository) domain.AnimalUseCase {
	return &animalUseCase{
		animalRepo: animalRepo,
	}
}

func (a animalUseCase) Create(animal *domain.Animal) error {
	a.beforeRequestForAnimal(animal)

	err := a.animalRepo.Save(animal)
	if err != nil {
		return errors.New("error creating animal")
	}

	return nil
}

func (a animalUseCase) GetByGuid(guid string) (*domain.Animal, error) {
	animal, err := a.animalRepo.GetByGuid(guid)
	if err != nil {
		return nil, errors.New("error fetching animal")
	}

	return animal, nil
}

func (a animalUseCase) GetAll() ([]domain.Animal, error) {
	animals, err := a.animalRepo.GetAll()
	if err != nil {
		return nil, errors.New("error fetching list of animals")
	}

	return animals, nil
}

func (a animalUseCase) Update(animal *domain.Animal) error {
	existAnimal, err := a.animalRepo.GetByGuid(animal.Guid())
	if err != nil {
		return errors.New("error fetching animal")
	}
	if animal.Descriptive() != "" {
		existAnimal.SetDescriptive(animal.Descriptive())
	}
	if animal.Weight() <= 0 {
		existAnimal.SetWeight(animal.Weight())
	}
	existAnimal.SetIsIll(animal.IsIll())
	existAnimal.SetIsFed(animal.IsFed())
	existAnimal.SetLastFedTime(animal.LastFedTime())

	existAnimal.SetUpdatedAt(time.Now())
	err = a.animalRepo.Save(existAnimal)
	if err != nil {
		return errors.New("error updating animal")
	}

	return nil
}

func (a animalUseCase) Delete(guid string) error {
	err := a.animalRepo.Delete(guid)
	if err != nil {
		return errors.New("error deleting animal")
	}

	return nil
}
