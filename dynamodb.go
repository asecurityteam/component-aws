package aws

import (
	"context"

	"github.com/asecurityteam/settings/v2"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDBConfig contains the settings for a DynamoDB client.
type DynamoDBConfig struct {
	Session *SessionConfig
}

// Name of the configuration root.
func (*DynamoDBConfig) Name() string {
	return "dynamodb"
}

// DynamoDBComponent is used to create a new DynamoDB client.
type DynamoDBComponent struct {
	Session *SessionComponent
}

// NewDynamoDBComponent generates a default component.
func NewDynamoDBComponent() *DynamoDBComponent {
	return &DynamoDBComponent{
		Session: NewSessionComponent(),
	}
}

// Settings generates the default settings for the component.
func (c *DynamoDBComponent) Settings() *DynamoDBConfig {
	return &DynamoDBConfig{
		Session: c.Session.Settings(),
	}
}

// New constructs a DynamoDB client.
func (c *DynamoDBComponent) New(ctx context.Context, conf *DynamoDBConfig) (*dynamodb.DynamoDB, error) {
	ses, err := c.Session.New(ctx, conf.Session)
	if err != nil {
		return nil, err
	}
	return dynamodb.New(ses), nil
}

// LoadDynamoDB is a convenience method for binding the source to the component.
func LoadDynamoDB(ctx context.Context, source settings.Source, c *DynamoDBComponent) (*dynamodb.DynamoDB, error) {
	dst := new(dynamodb.DynamoDB)
	err := settings.NewComponent(ctx, source, c, dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// NewDynamoDB is the top-level entry point for creating a DynamoDB client.
func NewDynamoDB(ctx context.Context, source settings.Source) (*dynamodb.DynamoDB, error) {
	return LoadDynamoDB(ctx, source, NewDynamoDBComponent())
}
