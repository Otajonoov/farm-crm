package domain

type Food struct {
	guid string
	name string
	mass float64
}

func (f *Food) Guid() string {
	return f.guid
}

func (f *Food) SetGuid(guid string) {
	f.guid = guid
}

func (f *Food) Name() string {
	return f.name
}

func (f *Food) SetName(name string) {
	f.name = name
}

func (f *Food) Mass() float64 {
	return f.mass
}

func (f *Food) SetMass(mass float64) {
	f.mass = mass
}
