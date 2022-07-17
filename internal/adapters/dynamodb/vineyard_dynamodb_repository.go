package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type VineyardModel struct {
	UUID string `dynamodbav:"UUID"`

	Name string `dynamodbav:"Name"`
}

func (v VineyardModel) GetKey() (map[string]types.AttributeValue, error) {
	uuid, err := attributevalue.Marshal(v.UUID)
	if err != nil {
		return map[string]types.AttributeValue{}, err
	}

	return map[string]types.AttributeValue{"UUID": uuid}, nil
}
