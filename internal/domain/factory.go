package domain

import (
	"time"
)

type Factory struct{}

func (a *Factory) ParseToDomain(guid, animalTypeGuid, descriptive string, weight float32, isIll, isFed, isWatered bool, lastFedTime, createdAt, updatedAt time.Time) *Animal {
	return &Animal{
		guid:           guid,
		animalTypeGuid: animalTypeGuid,
		descriptive:    descriptive,
		weight:         weight,
		isIll:          isIll,
		isFed:          isFed,
		isWatered:      isWatered,
		lastFedTime:    lastFedTime,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
	}
}

func (a *Factory) ParseToDomainFor(animalTypeGuid, descriptive string, weight float32, isIll, isFed, isWatered bool) *Animal {
	return &Animal{
		animalTypeGuid: animalTypeGuid,
		descriptive:    descriptive,
		weight:         weight,
		isIll:          isIll,
		isFed:          isFed,
		isWatered:      isWatered,
	}
}

func (a *Factory) ParseToDomainForAnimalType(Guid, animalType string, count int) *AnimalType {
	return &AnimalType{
		animalGuid: Guid,
		animalType: animalType,
		count:      count,
	}
}
