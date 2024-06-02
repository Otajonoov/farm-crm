package domain

import (
	"time"
)

type Factory struct{}

func (a *Factory) ParseToDomain(guid, animalTypeGuid, descriptive string, weight float32, isIll, isFed bool, lastFedTime, createdAt, updatedAt time.Time) *Animal {
	return &Animal{
		guid:           guid,
		animalTypeGuid: animalTypeGuid,
		descriptive:    descriptive,
		weight:         weight,
		isIll:          isIll,
		isFed:          isFed,
		lastFedTime:    lastFedTime,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
	}
}

func (a *Factory) ParseToDomainFor(animalTypeGuid, descriptive string, weight float32, isIll, isFed bool) *Animal {
	return &Animal{
		animalTypeGuid: animalTypeGuid,
		descriptive:    descriptive,
		weight:         weight,
		isIll:          isIll,
		isFed:          isFed,
	}
}
