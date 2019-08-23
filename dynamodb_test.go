package aws

import (
	"context"
	"testing"

	"github.com/asecurityteam/settings"
	"github.com/stretchr/testify/require"
)

func TestDynamoDB(t *testing.T) {
	cmp := NewDynamoDBComponent()
	require.Equal(t, "dynamodb", cmp.Settings().Name())
}

func TestDynamoDBSettings(t *testing.T) {
	src := settings.NewMapSource(map[string]interface{}{
		"dynamodb": map[string]interface{}{},
	})
	p, err := NewDynamoDB(context.Background(), src)
	require.Nil(t, err)
	require.NotNil(t, p)
}
