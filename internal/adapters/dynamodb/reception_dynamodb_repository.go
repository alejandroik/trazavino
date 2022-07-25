package dynamodb

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ReceptionModel struct {
	UUID string `dynamodbav:"UUID"`

	StartTime time.Time `dynamodbav:"StartTime"`

	WineryUUID string `dynamodbav:"WineryUUID"`
	Winery     string `dynamodbav:"Winery"`

	TruckUUID string `dynamodbav:"TruckUUID"`
	Truck     string `dynamodbav:"Truck"`

	VineyardUUID string `dynamodbav:"VineyardUUID"`
	Vineyard     string `dynamodbav:"Vineyard"`

	GrapeTypeUUID string `dynamodbav:"GrapeTypeUUID"`
	GrapeType     string `dynamodbav:"GrapeType"`

	Weight int32 `dynamodbav:"Weight"`
	Sugar  int32 `dynamodbav:"Sugar"`

	EndTime     *time.Time `dynamodbav:"EndTime"`
	Hash        *string    `dynamodbav:"Hash"`
	Transaction *string    `dynamodbav:"Transaction"`
}

type ReceptionDynamoDbRepository struct {
	dynamodbClient *dynamodb.Client
}

func NewReceptionDynamoDbRepository(client *dynamodb.Client) ReceptionDynamoDbRepository {
	return ReceptionDynamoDbRepository{dynamodbClient: client}
}

func (r ReceptionDynamoDbRepository) receptionTable() *string {
	return aws.String("Reception")
}

func (r ReceptionDynamoDbRepository) AddReception(ctx context.Context, rc *entity.Reception) error {
	rm, err := r.marshalReception(rc)
	if err != nil {
		return err
	}

	var transactItems []types.TransactWriteItem

	condition := expression.AttributeExists(expression.Name("UUID"))
	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	transactItems = append(transactItems, types.TransactWriteItem{
		ConditionCheck: &types.ConditionCheck{
			Key:                      stringKey("UUID", rc.TruckUUID()),
			TableName:                aws.String("Truck"),
			ConditionExpression:      expr.Condition(),
			ExpressionAttributeNames: expr.Names(),
		},
	})

	transactItems = append(transactItems, types.TransactWriteItem{
		ConditionCheck: &types.ConditionCheck{
			Key:                      stringKey("UUID", rc.VineyardUUID()),
			TableName:                aws.String("Vineyard"),
			ConditionExpression:      expr.Condition(),
			ExpressionAttributeNames: expr.Names(),
		},
	})

	transactItems = append(transactItems, types.TransactWriteItem{
		ConditionCheck: &types.ConditionCheck{
			Key:                      stringKey("UUID", rc.GrapeTypeUUID()),
			TableName:                aws.String("GrapeType"),
			ConditionExpression:      expr.Condition(),
			ExpressionAttributeNames: expr.Names(),
		},
	})

	transactItems = append(transactItems, types.TransactWriteItem{
		Put: &types.Put{
			Item:      rm,
			TableName: r.receptionTable(),
		},
	})

	_, err = r.dynamodbClient.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{TransactItems: transactItems})
	return err
}

func (r ReceptionDynamoDbRepository) GetReception(ctx context.Context, receptionUUID string) (*entity.Reception, error) {
	response, err := r.dynamodbClient.GetItem(ctx, &dynamodb.GetItemInput{
		Key: stringKey("UUID", receptionUUID), TableName: r.receptionTable(),
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
	reception, err := r.GetReception(ctx, receptionUUID)
	if err != nil {
		return err
	}

	updReception, err := updateFn(ctx, reception)
	if err != nil {
		return err
	}

	rm, err := r.marshalReception(updReception)
	if err != nil {
		return err
	}

	if _, err = r.dynamodbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: r.receptionTable(), Item: rm,
	}); err != nil {
		return err
	}

	return nil
}

func (r ReceptionDynamoDbRepository) marshalReception(rc *entity.Reception) (document, error) {
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

	et := rc.EndTime()
	if !et.IsZero() {
		receptionModel.EndTime = &et
	}
	hash := rc.Hash()
	if hash != "" {
		receptionModel.Hash = &hash
	}
	transaction := rc.Transaction()
	if transaction != "" {
		receptionModel.Transaction = &transaction
	}

	return attributevalue.MarshalMap(receptionModel)
}

func (r ReceptionDynamoDbRepository) unmarshalReception(av document) (*entity.Reception, error) {
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
		rm.WineryUUID,
		rm.Winery,
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
