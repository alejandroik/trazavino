package dynamodb

type GrapeTypeModel struct {
	UUID string `dynamodbav:"UUID"`

	Name string `dynamodbav:"Name"`
}

func (gt GrapeTypeModel) getKey() document {
	return stringKey("UUID", gt.UUID)
}
