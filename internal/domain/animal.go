package domain

import "time"

type Animal struct {
	guid           string
	animalTypeGuid string
	descriptive    string
	weight         float32
	isIll          bool
	isFed          bool
	lastFedTime    time.Time
	createdAt      time.Time
	updatedAt      time.Time
}

func (a *Animal) Guid() string {
	return a.guid
}

func (a *Animal) SetGuid(guid string) {
	a.guid = guid
}

func (a *Animal) AnimalTypeGuid() string {
	return a.animalTypeGuid
}

func (a *Animal) SetAnimalTypeGuid(animalTypeGuid string) {
	a.animalTypeGuid = animalTypeGuid
}

func (a *Animal) Descriptive() string {
	return a.descriptive
}

func (a *Animal) SetDescriptive(descriptive string) {
	a.descriptive = descriptive
}

func (a *Animal) Weight() float32 {
	return a.weight
}

func (a *Animal) SetWeight(weight float32) {
	a.weight = weight
}

func (a *Animal) IsIll() bool {
	return a.isIll
}

func (a *Animal) SetIsIll(isIll bool) {
	a.isIll = isIll
}

func (a *Animal) IsFed() bool {
	return a.isFed
}

func (a *Animal) SetIsFed(isFed bool) {
	a.isFed = isFed
}

func (a *Animal) LastFedTime() time.Time {
	return a.lastFedTime
}

func (a *Animal) SetLastFedTime(lastFedTime time.Time) {
	a.lastFedTime = lastFedTime
}

func (a *Animal) CreatedAt() time.Time {
	return a.createdAt
}

func (a *Animal) SetCreatedAt(createdAt time.Time) {
	a.createdAt = createdAt
}

func (a *Animal) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a *Animal) SetUpdatedAt(updatedAt time.Time) {
	a.updatedAt = updatedAt
}

type AnimalRepository interface {
	Save(*Animal) error
	GetByGuid(string) (*Animal, error)
	GetAll() ([]Animal, error)
	Update(*Animal) error
	Delete(string) error
}

type AnimalUseCase interface {
	Create(*Animal) error
	GetByGuid(string) (*Animal, error)
	GetAll() ([]Animal, error)
	Update(*Animal) error
	Delete(string) error
}
