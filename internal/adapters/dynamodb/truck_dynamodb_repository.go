package dynamodb

type TruckModel struct {
	UUID string `dynamodbav:"UUID"`

	License string `dynamodbav:"License"`
}

func (t TruckModel) getKey() document {
	return stringKey("UUID", t.UUID)
}
