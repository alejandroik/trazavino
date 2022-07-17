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

	StartTime time.Time  `dynamodbav:"StartTime"`
	EndTime   *time.Time `dynamodbav:"EndTime"`

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

func (r ReceptionModel) GetKey() (map[string]types.AttributeValue, error) {
	uuid, err := attributevalue.Marshal(r.UUID)
	if err != nil {
		return map[string]types.AttributeValue{}, err
	}

	return map[string]types.AttributeValue{"UUID": uuid}, nil
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

	transactItems = append(
		transactItems,
		types.TransactWriteItem{
			ConditionCheck: &types.ConditionCheck{
				Key:                      stringKey("UUID", rc.TruckUUID()),
				TableName:                aws.String("Truck"),
				ConditionExpression:      aws.String("attribute_exists(#U)"),
				ExpressionAttributeNames: map[string]string{"#U": "UUID"},
			},
		})

	transactItems = append(
		transactItems,
		types.TransactWriteItem{
			ConditionCheck: &types.ConditionCheck{
				Key:                      stringKey("UUID", rc.VineyardUUID()),
				TableName:                aws.String("Vineyard"),
				ConditionExpression:      aws.String("attribute_exists(#U)"),
				ExpressionAttributeNames: map[string]string{"#U": "UUID"},
			},
		})

	transactItems = append(
		transactItems,
		types.TransactWriteItem{
			ConditionCheck: &types.ConditionCheck{
				Key:                      stringKey("UUID", rc.GrapeTypeUUID()),
				TableName:                aws.String("GrapeType"),
				ConditionExpression:      aws.String("attribute_exists(#U)"),
				ExpressionAttributeNames: map[string]string{"#U": "UUID"},
			},
		})

	transactItems = append(
		transactItems,
		types.TransactWriteItem{
			Put: &types.Put{
				Item:      putReceptionItem,
				TableName: r.receptionTable(),
			},
		})

	_, err = r.dynamoDbClient.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{TransactItems: transactItems})
	return err
}

func (r ReceptionDynamoDbRepository) GetReception(ctx context.Context, receptionUUID string) (*entity.Reception, error) {
	rm := ReceptionModel{UUID: receptionUUID}
	rk, err := rm.GetKey()
	if err != nil {
		return nil, err
	}

	response, err := r.dynamoDbClient.GetItem(ctx, &dynamodb.GetItemInput{
		Key: rk, TableName: r.receptionTable(),
	})
	if err != nil {
		return nil, err
	}

	rc, err := r.unmarshalReception(response.Item)
	if err != nil {
		return nil, err
	}

	return rc, nil
}

func (r ReceptionDynamoDbRepository) UpdateReception(ctx context.Context, receptionUUID string, updateFn func(ctx context.Context, rc *entity.Reception) (*entity.Reception, error)) error {
	return nil
}

func (r ReceptionDynamoDbRepository) marshalReception(rc *entity.Reception) ReceptionModel {
	receptionModel := ReceptionModel{
		UUID:          rc.UUID(),
		StartTime:     rc.StartTime(),
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

func (r ReceptionDynamoDbRepository) unmarshalReception(av Document) (*entity.Reception, error) {
	rm := ReceptionModel{}
	err := attributevalue.UnmarshalMap(av, &rm)
	if err != nil {
		return nil, err
	}

	var endTime time.Time
	if rm.EndTime != nil {
		endTime = *rm.EndTime
	}

	var hash string
	if rm.Hash != nil {
		hash = *rm.Hash
	}

	var transaction string
	if rm.Transaction != nil {
		transaction = *rm.Transaction
	}

	return entity.UnmarshalReceptionFromDatabase(
		rm.UUID,
		rm.StartTime,
		rm.TruckUUID,
		rm.Truck,
		rm.VineyardUUID,
		rm.Vineyard,
		rm.GrapeTypeUUID,
		rm.GrapeType,
		rm.Weight,
		rm.Sugar,
		endTime,
		hash,
		transaction)
}
