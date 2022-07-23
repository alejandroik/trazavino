package entity

type GrapeType struct {
	baseEntity
}

func NewGrapeType(id string, name string) (*GrapeType, error) {
	return &GrapeType{baseEntity{uuid: id, name: name}}, nil
}

func UnmarshalGrapeTypeFromDatabase(id string, name string) (*GrapeType, error) {
	return NewGrapeType(id, name)
}

func (t GrapeType) ID() string {
	return t.uuid
}

func (t GrapeType) Name() string {
	return t.name
}
