package entity

type Cask struct {
	baseEntity
	cType   string
	isEmpty bool
}

func NewCask(id string, name string, cType string, isEmpty bool) (*Cask, error) {
	return &Cask{
		baseEntity: baseEntity{uuid: id, name: name},
		cType:      cType,
		isEmpty:    isEmpty,
	}, nil
}

func (c Cask) ID() string {
	return c.uuid
}

func (c Cask) Name() string {
	return c.name
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

func UnmarshalCaskFromDatabase(id string, name string, cType string, isEmpty bool) (*Cask, error) {
	return NewCask(id, name, cType, isEmpty)
}
