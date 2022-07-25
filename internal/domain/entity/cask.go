package entity

type Cask struct {
	ownedEntity
	cType   string
	isEmpty bool
}

func NewCask(id string, name string, cType string, isEmpty bool, wineryUUID string) (*Cask, error) {
	return &Cask{
		ownedEntity: ownedEntity{baseEntity{uuid: id, name: name}, wineryUUID},
		cType:       cType,
		isEmpty:     isEmpty,
	}, nil
}

func (c Cask) CType() string {
	return c.cType
}

func (c Cask) IsEmpty() bool {
	return c.isEmpty
}

func (c *Cask) UpdateIsEmpty(v bool) error {
	c.isEmpty = v

	return nil
}

func UnmarshalCaskFromDatabase(id string, name string, cType string, isEmpty bool, wineryUUID string) (*Cask, error) {
	return NewCask(id, name, cType, isEmpty, wineryUUID)
}
