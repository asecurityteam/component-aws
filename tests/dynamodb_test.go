// +build integration

package tests

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	aws "github.com/asecurityteam/component-aws"
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/require"
)

func TestDynamo(t *testing.T) {
	tableName := fmt.Sprintf("%d", rand.Int63())
	partitionKey := fmt.Sprintf("%d", rand.Int63())
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
				&dynamodb.AttributeDefinition{
					AttributeName: awssdk.String(partitionKey),
					AttributeType: awssdk.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				&dynamodb.KeySchemaElement{
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
