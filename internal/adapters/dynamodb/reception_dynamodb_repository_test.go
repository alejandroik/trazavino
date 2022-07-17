package dynamodb

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func newReception() *entity.Reception {
	receptionUUID := uuid.New().String()
	receptionTime := time.Now().Round(time.Second)
	truckUUID := "a57ccb31-a58b-4f6e-817e-b0a85fc652e4"
	truckLicense := "EEOO-990"
	vineyardUUID := "5da1acf8-d613-4af0-8b82-15493346bf56"
	vineyardName := "Los Andes"
	grapeTypeUUID := "6555c14c-07da-44c0-a85e-094c250448ea"
	grapeTypeName := "Rosada"
	weight := int32(2500)
	sugar := int32(15)

	rc, _ := entity.NewReception(
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

	return rc
}

func TestReceptionDynamoDbRepository_AddReception(t *testing.T) {
	rc := newReception()
	ctx := context.Background()

	os.Setenv("LOCAL_DB_ENDPOINT", "http://localhost:8000")
	client, err := NewDynamoDbClient(ctx)
	require.NoError(t, err)
	recRepo := NewReceptionDynamoDbRepository(client)
	err = recRepo.AddReception(ctx, rc)
	require.NoError(t, err)
}

func TestReceptionDynamoDbRepository_Marshal(t *testing.T) {
	rc := newReception()
	ctx := context.Background()

	os.Setenv("LOCAL_DB_ENDPOINT", "http://localhost:8000")
	client, err := NewDynamoDbClient(ctx)
	require.NoError(t, err)
	recRepo := NewReceptionDynamoDbRepository(client)

	rm := recRepo.marshalReception(rc)
	putReceptionItem, err := attributevalue.MarshalMap(rm)
	require.NoError(t, err)
	reception, err := recRepo.unmarshalReception(putReceptionItem)
	require.NoError(t, err)
	t.Log(reception)
}

func TestReceptionDynamoDbRepository_GetReception(t *testing.T) {
	ctx := context.Background()

	os.Setenv("LOCAL_DB_ENDPOINT", "http://localhost:8000")
	client, err := NewDynamoDbClient(ctx)
	require.NoError(t, err)
	recRepo := NewReceptionDynamoDbRepository(client)

	rc, err := recRepo.GetReception(ctx, "54293b87-be54-401b-8efe-983b7d109ee3")
	require.NoError(t, err)
	t.Log(rc)
}
