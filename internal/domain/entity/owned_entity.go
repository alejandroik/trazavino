package entity

type ownedEntity struct {
	baseEntity
	wineryUUID string
}

func (e ownedEntity) WineryUUID() string {
	return e.wineryUUID
}
