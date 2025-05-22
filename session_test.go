package aws

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"

	"github.com/asecurityteam/settings/v2"
	"github.com/stretchr/testify/require"
)

func TestSession(t *testing.T) {
	cmp := NewSessionComponent()
	s := cmp.Settings()
	require.Equal(t, "session", s.Name())
	require.Equal(t, "sharedprofile", s.SharedCredentialConfig.Name())
	require.Equal(t, "static", s.StaticCredentialConfig.Name())
	require.Equal(t, "assumerole", s.AssumeRoleCredentialConfig.Name())
	numGen, err := rand.Int(rand.Reader, big.NewInt(27))
	if err != nil {
		require.NoError(t, err)
	}
	s.Region = fmt.Sprintf("%d", numGen.Int64())
	s.Endpoint = fmt.Sprintf("%d", numGen.Int64())
	s.SharedCredentialConfig.File = fmt.Sprintf("%d", numGen.Int64())
	s.SharedCredentialConfig.Profile = fmt.Sprintf("%d", numGen.Int64())
	s.StaticCredentialConfig.ID = fmt.Sprintf("%d", numGen.Int64())
	s.StaticCredentialConfig.Secret = fmt.Sprintf("%d", numGen.Int64())
	s.StaticCredentialConfig.Token = fmt.Sprintf("%d", numGen.Int64())
	s.AssumeRoleCredentialConfig.Role = fmt.Sprintf("%d", numGen.Int64())
	s.AssumeRoleCredentialConfig.ExternalID = fmt.Sprintf("%d", numGen.Int64())

	ses, err := cmp.New(context.Background(), s)
	require.Nil(t, err)
	require.NotNil(t, ses)
}

func TestSessionSettings(t *testing.T) {
	numGen, err := rand.Int(rand.Reader, big.NewInt(27))
	if err != nil {
		require.NoError(t, err)
	}
	src := settings.NewMapSource(map[string]interface{}{
		"session": map[string]interface{}{
			"region":        fmt.Sprintf("%d", numGen.Int64()),
			"endpoint":      fmt.Sprintf("%d", numGen.Int64()),
			"sharedprofile": map[string]interface{}{},
			"static":        map[string]interface{}{},
			"assumerole":    map[string]interface{}{},
		},
	})
	p, err := NewSession(context.Background(), src)
	require.Nil(t, err)
	require.NotNil(t, p)
}
