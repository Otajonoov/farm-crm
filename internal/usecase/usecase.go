package usecase

import (
	"farm-crm/internal/domain"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

type BaseUseCase struct{}

func (u *BaseUseCase) Error(msg string, err error) error {
	if len(strings.TrimSpace(msg)) != 0 {
		return fmt.Errorf("%v: %w", msg, err)
	}
	return err
}

func (u *BaseUseCase) beforeRequestForAnimal(animal *domain.Animal) {
	if animal.Guid() == "" {
		animal.SetGuid(uuid.New().String())
	}
	if animal.CreatedAt().IsZero() {
		animal.SetCreatedAt(time.Now())
	}
	if animal.UpdatedAt().IsZero() {
		animal.SetUpdatedAt(time.Now())
	}
}

func (u *BaseUseCase) beforeRequestForAnimalType(animal *domain.AnimalType) {
	if animal.AnimalGuid() == "" {
		animal.SetAnimalType(uuid.New().String())
	}
}
