package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type GrapeTypeModel struct {
	UUID string `dynamodbav:"UUID"`

	Name string `dynamodbav:"Name"`
}

func (gt GrapeTypeModel) GetKey() (map[string]types.AttributeValue, error) {
	uuid, err := attributevalue.Marshal(gt.UUID)
	if err != nil {
		return map[string]types.AttributeValue{}, err
	}

	return map[string]types.AttributeValue{"UUID": uuid}, nil
}
