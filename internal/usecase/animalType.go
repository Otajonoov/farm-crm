package usecase

import (
	"errors"
	"farm-crm/internal/domain"
)

type animalTypeUseCase struct {
	BaseUseCase
	animalTypeRepo domain.AnimalTypeRepository
}

func NewAnimalTypeUseCase(animalTypeRepo domain.AnimalTypeRepository) domain.AnimalTypeUseCase {
	return &animalTypeUseCase{
		animalTypeRepo: animalTypeRepo}
}

func (a *animalTypeUseCase) Create(animalType *domain.AnimalType) error {
	a.beforeRequestForAnimalType(animalType)

	err := a.animalTypeRepo.Save(animalType)
	if err != nil {
		return errors.New("animal type create fail")
	}

	return nil
}

func (a *animalTypeUseCase) GetAnimalTypeInfo() ([]domain.AnimalType, error) {
	info, err := a.animalTypeRepo.GetAnimalTypeInfo()
	if err != nil {
		return nil, errors.New("failed to retrieve animal info")
	}

	return info, nil
}

func (a *animalTypeUseCase) Update(animalType domain.AnimalType) error {
	err := a.animalTypeRepo.Update(animalType)
	if err != nil {
		return errors.New("animal type update failed")
	}

	return nil
}

func (a *animalTypeUseCase) Delete(guid string) error {
	err := a.animalTypeRepo.Delete(guid)
	if err != nil {
		return errors.New("animal type delete failed")
	}

	return nil
}
