package domain

type Warehouse struct {
	guid        string
	sectionName string
	foodGuid    string
	drugGuid    string
}

func (w *Warehouse) Guid() string {
	return w.guid
}

func (w *Warehouse) SetGuid(guid string) {
	w.guid = guid
}

func (w *Warehouse) SectionName() string {
	return w.sectionName
}

func (w *Warehouse) SetSectionName(sectionName string) {
	w.sectionName = sectionName
}

func (w *Warehouse) FoodGuid() string {
	return w.foodGuid
}

func (w *Warehouse) SetFoodGuid(foodGuid string) {
	w.foodGuid = foodGuid
}

func (w *Warehouse) DrugGuid() string {
	return w.drugGuid
}

func (w *Warehouse) SetDrugGuid(drugGuid string) {
	w.drugGuid = drugGuid
}
