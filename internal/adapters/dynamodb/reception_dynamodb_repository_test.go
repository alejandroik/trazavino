package dynamodb

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestReceptionDynamoDbRepository_AddReception(t *testing.T) {
	receptionUUID := uuid.New().String()
	receptionTime := time.Now().Round(time.Second)
	truckUUID := uuid.New().String()
	truckLicense := "EEOO-990"
	vineyardUUID := uuid.New().String()
	vineyardName := "Los Andes"
	grapeTypeUUID := uuid.New().String()
	grapeTypeName := "Rosada"
	weight := int32(2500)
	sugar := int32(15)

	rc, err := entity.NewReception(
		receptionUUID,
		receptionTime,
		truckUUID,
		truckLicense,
		vineyardUUID,
		vineyardName,
		grapeTypeUUID,
		grapeTypeName,
		weight,
		sugar)
	require.NoError(t, err)

	ctx := context.Background()

	os.Setenv("LOCAL_DB_ENDPOINT", "http://localhost:8000")
	client, err := NewDynamoDbClient(ctx)
	require.NoError(t, err)
	recRepo := NewReceptionDynamoDbRepository(client)
	err = recRepo.AddReception(ctx, rc)
	require.NoError(t, err)
}
