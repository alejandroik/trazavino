package entity

type Bottle struct {
	ownedEntity
}

func NewBottle(id string, name string, wineryUUID string) (*Bottle, error) {
	return &Bottle{ownedEntity{baseEntity{uuid: id, name: name}, wineryUUID}}, nil
}

func UnmarshalBottleFromDatabase(id string, name string, wineryUUID string) (*Bottle, error) {
	return NewBottle(id, name, wineryUUID)
}
