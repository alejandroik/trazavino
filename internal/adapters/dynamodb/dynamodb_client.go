package dynamodb

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/pkg/errors"
)

func NewDynamoDbClient(ctx context.Context) (*dynamodb.Client, error) {
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	dbEndpoint := os.Getenv("LOCAL_DB_ENDPOINT")
	if dbEndpoint == "" {
		return nil, errors.New("empty db endpoint")
		//return dynamodb.NewFromConfig(sdkConfig), nil
	}

	return dynamodb.NewFromConfig(
		sdkConfig,
		dynamodb.WithEndpointResolver(
			dynamodb.EndpointResolverFromURL(dbEndpoint)),
	), nil
}
