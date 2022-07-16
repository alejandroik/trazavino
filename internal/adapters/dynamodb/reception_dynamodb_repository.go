package dynamodb

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ReceptionModel struct {
	UUID string `dynamodbav:"UUID"`

	Time time.Time `dynamodbav:"Time"`

	TruckUUID string `dynamodbav:"TruckUUID"`
	Truck     string `dynamodbav:"Truck"`

	VineyardUUID string `dynamodbav:"VineyardUUID"`
	Vineyard     string `dynamodbav:"Vineyard"`

	GrapeTypeUUID string `dynamodbav:"GrapeTypeUUID"`
	GrapeType     string `dynamodbav:"GrapeType"`

	Weight int32 `dynamodbav:"Weight"`
	Sugar  int32 `dynamodbav:"Sugar"`

	Hash        *string `dynamodbav:"Hash"`
	Transaction *string `dynamodbav:"Transaction"`
}

type ReceptionDynamoDbRepository struct {
	dynamoDbClient *dynamodb.Client
}

func NewReceptionDynamoDbRepository(client *dynamodb.Client) ReceptionDynamoDbRepository {
	return ReceptionDynamoDbRepository{dynamoDbClient: client}
}

func (r ReceptionDynamoDbRepository) receptionTable() *string {
	return aws.String("Reception")
}

func (r ReceptionDynamoDbRepository) AddReception(ctx context.Context, rc *entity.Reception) error {
	rm := r.marshalReception(rc)

	putReceptionItem, err := attributevalue.MarshalMap(rm)
	if err != nil {
		return err
	}

	var transactItems []types.TransactWriteItem
	transactItems = append(transactItems, types.TransactWriteItem{
		ConditionCheck: nil,
		Put: &types.Put{
			Item:      putReceptionItem,
			TableName: r.receptionTable(),
		},
	})

	_, err = r.dynamoDbClient.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{TransactItems: transactItems})
	return err
}

func (r ReceptionDynamoDbRepository) GetReception(ctx context.Context, receptionUUID string) (*entity.Reception, error) {
	return nil, nil
}

func (r ReceptionDynamoDbRepository) UpdateReception(ctx context.Context, receptionUUID string, updateFn func(ctx context.Context, rc *entity.Reception) (*entity.Reception, error)) error {
	return nil
}

func (r ReceptionDynamoDbRepository) marshalReception(rc *entity.Reception) ReceptionModel {
	receptionModel := ReceptionModel{
		UUID:          rc.UUID(),
		Time:          rc.Time(),
		TruckUUID:     rc.TruckUUID(),
		Truck:         rc.TruckLicense(),
		VineyardUUID:  rc.VineyardUUID(),
		Vineyard:      rc.VineyardName(),
		GrapeTypeUUID: rc.GrapeTypeUUID(),
		GrapeType:     rc.GrapeTypeName(),
		Weight:        rc.Weight(),
		Sugar:         rc.Sugar(),
	}

	return receptionModel
}
