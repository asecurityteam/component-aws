package aws

import (
	"context"
	"testing"

	"github.com/asecurityteam/settings/v2"
	"github.com/stretchr/testify/require"
)

func TestS3(t *testing.T) {
	cmp := NewS3Component()
	require.Equal(t, "s3", cmp.Settings().Name())
}

func TestS3Settings(t *testing.T) {
	src := settings.NewMapSource(map[string]interface{}{
		"s3": map[string]interface{}{},
	})
	p, err := NewS3(context.Background(), src)
	require.Nil(t, err)
	require.NotNil(t, p)
}
