package entity

type baseEntity struct {
	uuid string
	name string
}

func (e baseEntity) ID() string {
	return e.uuid
}

func (e baseEntity) Name() string {
	return e.name
}
