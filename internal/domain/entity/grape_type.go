package entity

type GrapeType struct {
	baseEntity
}

func NewGrapeType(id int64, name string) (*GrapeType, error) {
	return &GrapeType{baseEntity{id: id, name: name}}, nil
}

func (t GrapeType) ID() int64 {
	return t.id
}

func (t GrapeType) Name() string {
	return t.name
}
