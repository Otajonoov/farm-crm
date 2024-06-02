package domain

type Drugs struct {
	guid  string
	name  string
	count int
}

func (d *Drugs) Guid() string {
	return d.guid
}

func (d *Drugs) SetGuid(guid string) {
	d.guid = guid
}

func (d *Drugs) Name() string {
	return d.name
}

func (d *Drugs) SetName(name string) {
	d.name = name
}

func (d *Drugs) Count() int {
	return d.count
}

func (d *Drugs) SetCount(count int) {
	d.count = count
}
