package domain

type AnimalType struct {
	animalGuid string
	animalType string
	count      int
}

func (a *AnimalType) AnimalGuid() string {
	return a.animalGuid
}

func (a *AnimalType) SetAnimalGuid(animalGuid string) {
	a.animalGuid = animalGuid
}

func (a *AnimalType) AnimalType() string {
	return a.animalType
}

func (a *AnimalType) SetAnimalType(animalType string) {
	a.animalType = animalType
}

func (a *AnimalType) Count() int {
	return a.count
}

func (a *AnimalType) SetCount(count int) {
	a.count = count
}

type AnimalTypeRepository interface {
	Save(*AnimalType) error
	GetAnimalTypeInfo() ([]AnimalType, error)
	Update(AnimalType) error
	Delete(string) error
}

type AnimalTypeUseCase interface {
	Create(*AnimalType) error
	GetAnimalTypeInfo() ([]AnimalType, error)
	Update(AnimalType) error
	Delete(string) error
}
