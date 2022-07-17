package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type TruckModel struct {
	UUID string `dynamodbav:"UUID"`

	License string `dynamodbav:"License"`
}

func (t TruckModel) GetKey() (map[string]types.AttributeValue, error) {
	uuid, err := attributevalue.Marshal(t.UUID)
	if err != nil {
		return map[string]types.AttributeValue{}, err
	}

	return map[string]types.AttributeValue{"UUID": uuid}, nil
}
