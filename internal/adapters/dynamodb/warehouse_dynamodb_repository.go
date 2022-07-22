package dynamodb

import (
	"context"

	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type WarehouseModel struct {
	UUID string `dynamodbav:"UUID"`

	Name string `dynamodbav:"Name"`
}

type WarehouseDynamodbRepository struct {
	dynamodbClient *dynamodb.Client
}

func NewWarehouseDynamodbRepository(client *dynamodb.Client) WarehouseDynamodbRepository {
	return WarehouseDynamodbRepository{dynamodbClient: client}
}

func (r WarehouseDynamodbRepository) warehouseTable() *string {
	return aws.String("Warehouse")
}

func (r WarehouseDynamodbRepository) AddWarehouse(ctx context.Context, warehouse *entity.Warehouse) error {
	return nil
}

func (r WarehouseDynamodbRepository) GetWarehouse(ctx context.Context, warehouseId int64) (*entity.Warehouse, error) {
	return nil, nil
}

func (r WarehouseDynamodbRepository) ListWarehouses(ctx context.Context, offset int32, limit int32) ([]*entity.Warehouse, error) {
	return nil, nil
}

func (r WarehouseDynamodbRepository) UpdateWarehouse(
	ctx context.Context,
	warehouseId int64,
	updateFn func(ctx context.Context, wh *entity.Warehouse) (*entity.Warehouse, error),
) error {
	return nil
}
