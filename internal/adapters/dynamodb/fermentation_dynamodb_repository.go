package dynamodb

import (
	"context"
	"time"

	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type FermentationModel struct {
	UUID string `dynamodbav:"UUID"`

	StartTime time.Time `dynamodbav:"StartTime"`

	ReceptionUUID      string    `dynamodbav:"ReceptionUUID"`
	ReceptionStartTime time.Time `dynamodbav:"Reception"`

	WarehouseUUID string `dynamodbav:"WarehouseUUID"`
	Warehouse     string `dynamodbav:"Warehouse"`

	EndTime      *time.Time `dynamodbav:"EndTime"`
	PreviousUUID *string    `dynamodbav:"PreviousUUID"`

	Hash        *string `dynamodbav:"Hash"`
	Transaction *string `dynamodbav:"Transaction"`
}

type FermentationDynamodbRepository struct {
	dynamodbClient *dynamodb.Client
}

func NewFermentationDynamodbRepository(client *dynamodb.Client) FermentationDynamodbRepository {
	return FermentationDynamodbRepository{dynamodbClient: client}
}

func (r FermentationDynamodbRepository) fermentationTable() *string {
	return aws.String("Fermentation")
}

func (r FermentationDynamodbRepository) AddFermentation(ctx context.Context, f *entity.Fermentation) error {
	return nil
}

func (r FermentationDynamodbRepository) GetFermentation(ctx context.Context, fermentationUUID string) (*entity.Fermentation, error) {
	return nil, nil
}

func (r FermentationDynamodbRepository) UpdateFermentation(
	ctx context.Context,
	fermentationUUID string,
	updateFn func(ctx context.Context, f *entity.Fermentation) (*entity.Fermentation, error),
) error {
	return nil
}
