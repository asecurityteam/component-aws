//go:build integration

package tests

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	aws "github.com/asecurityteam/component-aws"
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/require"
)

func TestDynamo(t *testing.T) {
	numGen, err := rand.Int(rand.Reader, big.NewInt(27))
	if err != nil {
		require.NoError(t, err)
	}
	tableName := fmt.Sprintf("abc")
	partitionKey := fmt.Sprintf("%d", numGen.Int64())
	region := os.Getenv("DYNAMO_REGION")
	endpoint := os.Getenv("DYNAMO_ENDPOINT")

	cmp := aws.NewDynamoDBComponent()
	conf := cmp.Settings()
	conf.Session.Region = region
	conf.Session.Endpoint = endpoint

	client, err := cmp.New(context.Background(), conf)
	require.Nil(t, err)

	start := time.Now()
	for time.Since(start) < 10*time.Second {
		_, err = client.CreateTable(&dynamodb.CreateTableInput{
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: awssdk.String(partitionKey),
					AttributeType: awssdk.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: awssdk.String(partitionKey),
					KeyType:       awssdk.String("HASH"),
				},
			},
			TableName: awssdk.String(tableName),
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  awssdk.Int64(10),
				WriteCapacityUnits: awssdk.Int64(10),
			},
		})
		if err == nil {
			break
		}
		t.Log(err)
	}
	require.Nil(t, err)
}
