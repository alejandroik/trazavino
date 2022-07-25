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

type MacerationModel struct {
	UUID string `dynamodbav:"UUID"`

	StartTime time.Time `dynamodbav:"StartTime"`

	WineryUUID string `dynamodbav:"WineryUUID"`
	Winery     string `dynamodbav:"Winery"`

	ReceptionUUID      string    `dynamodbav:"ReceptionUUID"`
	ReceptionStartTime time.Time `dynamodbav:"Reception"`

	WarehouseUUID string `dynamodbav:"WarehouseUUID"`
	Warehouse     string `dynamodbav:"Warehouse"`

	EndTime      *time.Time `dynamodbav:"EndTime"`
	PreviousUUID *string    `dynamodbav:"PreviousUUID"`

	Hash        *string `dynamodbav:"Hash"`
	Transaction *string `dynamodbav:"Transaction"`
}

type MacerationDynamoDbRepository struct {
	dynamodbClient *dynamodb.Client
}

func NewMacerationDynamodbRepository(client *dynamodb.Client) MacerationDynamoDbRepository {
	return MacerationDynamoDbRepository{dynamodbClient: client}
}

func (r MacerationDynamoDbRepository) macerationTable() *string {
	return aws.String("Maceration")
}

func (r MacerationDynamoDbRepository) AddMaceration(ctx context.Context, mc *entity.Maceration) error {
	mm, err := r.marshalMaceration(mc)
	if err != nil {
		return nil
	}

	var transactItems []types.TransactWriteItem

	condition := expression.AttributeExists(expression.Name("UUID"))
	update := expression.Set(expression.Name("EndTime"), expression.Value(mc.StartTime()))
	expr, err := expression.NewBuilder().WithUpdate(update).WithCondition(condition).Build()
	if err != nil {
		return err
	}

	transactItems = append(transactItems, types.TransactWriteItem{
		Update: &types.Update{
			Key:                       stringKey("UUID", mc.ReceptionUUID()),
			TableName:                 aws.String("Reception"),
			UpdateExpression:          expr.Update(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			ConditionExpression:       expr.Condition(),
		},
	})

	update = expression.Set(expression.Name("IsEmpty"), expression.Value(false))
	expr, err = expression.NewBuilder().WithUpdate(update).WithCondition(condition).Build()
	if err != nil {
		return err
	}

	transactItems = append(transactItems, types.TransactWriteItem{
		Update: &types.Update{
			Key:                       stringKey("UUID", mc.WarehouseUUID()),
			TableName:                 aws.String("Warehouse"),
			UpdateExpression:          expr.Update(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			ConditionExpression:       expr.Condition(),
		},
	})

	transactItems = append(transactItems, types.TransactWriteItem{
		Put: &types.Put{
			Item:      mm,
			TableName: r.macerationTable(),
		},
	})

	_, err = r.dynamodbClient.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{TransactItems: transactItems})
	return err
}

func (r MacerationDynamoDbRepository) GetMaceration(ctx context.Context, macerationUUID string) (*entity.Maceration, error) {
	response, err := r.dynamodbClient.GetItem(ctx, &dynamodb.GetItemInput{
		Key: stringKey("UUID", macerationUUID), TableName: r.macerationTable(),
	})
	if err != nil {
		return nil, err
	}

	mc, err := r.unmarshalMaceration(response.Item)
	if err != nil {
		return nil, err
	}

	return mc, nil
}

func (r MacerationDynamoDbRepository) UpdateMaceration(ctx context.Context, macerationUUID string, updateFn func(ctx context.Context, m *entity.Maceration) (*entity.Maceration, error)) error {
	maceration, err := r.GetMaceration(ctx, macerationUUID)
	if err != nil {
		return err
	}

	updMaceration, err := updateFn(ctx, maceration)
	if err != nil {
		return err
	}

	rm, err := r.marshalMaceration(updMaceration)
	if err != nil {
		return err
	}

	if _, err = r.dynamodbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: r.macerationTable(), Item: rm,
	}); err != nil {
		return err
	}

	return nil
}

func (r MacerationDynamoDbRepository) marshalMaceration(mc *entity.Maceration) (document, error) {
	macerationModel := MacerationModel{
		UUID:               mc.UUID(),
		StartTime:          mc.StartTime(),
		ReceptionUUID:      mc.ReceptionUUID(),
		ReceptionStartTime: mc.ReceptionStartTime(),
		WarehouseUUID:      mc.WarehouseUUID(),
		Warehouse:          mc.WarehouseName(),
	}

	et := mc.EndTime()
	if !et.IsZero() {
		macerationModel.EndTime = &et
	}
	hash := mc.Hash()
	if hash != "" {
		macerationModel.Hash = &hash
	}
	transaction := mc.Transaction()
	if transaction != "" {
		macerationModel.Transaction = &transaction
	}

	return attributevalue.MarshalMap(macerationModel)
}

func (r MacerationDynamoDbRepository) unmarshalMaceration(av document) (*entity.Maceration, error) {
	mm := MacerationModel{}
	err := attributevalue.UnmarshalMap(av, &mm)
	if err != nil {
		return nil, err
	}

	var endTime time.Time
	if mm.EndTime != nil {
		endTime = *mm.EndTime
	}

	var previousUUID string
	if mm.PreviousUUID != nil {
		previousUUID = *mm.PreviousUUID
	}

	var hash string
	if mm.Hash != nil {
		hash = *mm.Hash
	}

	var transaction string
	if mm.Transaction != nil {
		transaction = *mm.Transaction
	}

	return entity.UnmarshalMacerationFromDatabase(
		mm.UUID,
		mm.StartTime,
		mm.WineryUUID,
		mm.Winery,
		mm.ReceptionUUID,
		mm.ReceptionStartTime,
		mm.WarehouseUUID,
		mm.Warehouse,
		endTime,
		previousUUID,
		hash,
		transaction,
	)
}
