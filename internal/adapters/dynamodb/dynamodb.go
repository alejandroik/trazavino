package dynamodb

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
)

type Document map[string]types.AttributeValue

func stringKey(k string, v string) Document {
	return Document{k: &types.AttributeValueMemberS{Value: v}}
}

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
