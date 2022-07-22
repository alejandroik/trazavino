package dynamodb

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func newMaceration() (*entity.Maceration, error) {
	macerationUUID := uuid.New().String()
	macerationTime := time.Now().Round(time.Second)
	receptionUUID := "9b29aba2-9e71-4662-8865-82e3f236e96f"
	receptionTime, err := time.Parse("2006-01-02T15:04:05", "2022-07-22T12:29:42")
	if err != nil {
		return nil, err
	}
	warehouseUUID := "93f24e5f-59b8-488d-a583-dc6d948140bb"
	warehouseName := "EEEEEEE"

	mc, err := entity.NewMaceration(
		macerationUUID,
		macerationTime,
		receptionUUID,
		receptionTime,
		warehouseUUID,
		warehouseName,
	)
	if err != nil {
		return nil, err
	}

	return mc, nil
}

func TestMacerationDynamoDbRepository_AddMaceration(t *testing.T) {
	ctx := context.Background()
	os.Setenv("LOCAL_DB_ENDPOINT", "http://localhost:8000")
	client, err := NewDynamoDbClient(ctx)
	require.NoError(t, err)

	mc, err := newMaceration()
	require.NoError(t, err)
	macRepo := NewMacerationDynamodbRepository(client)
	err = macRepo.AddMaceration(ctx, mc)
	require.NoError(t, err)
}
