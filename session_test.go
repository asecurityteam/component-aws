package aws

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/asecurityteam/settings"
	"github.com/stretchr/testify/require"
)

func TestSession(t *testing.T) {
	cmp := NewSessionComponent()
	s := cmp.Settings()
	require.Equal(t, "session", s.Name())
	require.Equal(t, "sharedprofile", s.SharedCredentialConfig.Name())
	require.Equal(t, "static", s.StaticCredentialConfig.Name())
	require.Equal(t, "assumerole", s.AssumeRoleCredentialConfig.Name())

	s.Region = fmt.Sprintf("%d", rand.Int63())
	s.Endpoint = fmt.Sprintf("%d", rand.Int63())
	s.SharedCredentialConfig.File = fmt.Sprintf("%d", rand.Int63())
	s.SharedCredentialConfig.Profile = fmt.Sprintf("%d", rand.Int63())
	s.StaticCredentialConfig.ID = fmt.Sprintf("%d", rand.Int63())
	s.StaticCredentialConfig.Secret = fmt.Sprintf("%d", rand.Int63())
	s.StaticCredentialConfig.Token = fmt.Sprintf("%d", rand.Int63())
	s.AssumeRoleCredentialConfig.Role = fmt.Sprintf("%d", rand.Int63())
	s.AssumeRoleCredentialConfig.ExternalID = fmt.Sprintf("%d", rand.Int63())

	ses, err := cmp.New(context.Background(), s)
	require.Nil(t, err)
	require.NotNil(t, ses)
}

func TestSessionSettings(t *testing.T) {
	src := settings.NewMapSource(map[string]interface{}{
		"session": map[string]interface{}{
			"region":        fmt.Sprintf("%d", rand.Int63()),
			"endpoint":      fmt.Sprintf("%d", rand.Int63()),
			"sharedprofile": map[string]interface{}{},
			"static":        map[string]interface{}{},
			"assumerole":    map[string]interface{}{},
		},
	})
	p, err := NewSession(context.Background(), src)
	require.Nil(t, err)
	require.NotNil(t, p)
}
