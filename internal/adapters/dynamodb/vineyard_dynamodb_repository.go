package dynamodb

type VineyardModel struct {
	UUID string `dynamodbav:"UUID"`

	Name string `dynamodbav:"Name"`
}

func (v VineyardModel) getKey() document {
	return stringKey("UUID", v.UUID)
}
