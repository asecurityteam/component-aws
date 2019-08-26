package aws

import (
	"context"

	"github.com/asecurityteam/settings"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

// SharedCredentialConfig contains settings for using the AWS shared profile
// configuration file to authenticate.
type SharedCredentialConfig struct {
	File    string `description:"The location of the shared profile configuration file. Leave blank for the default AWS location."`
	Profile string `description:"Name of the profile to use from the file."`
}

// Name of the configuration root.
func (*SharedCredentialConfig) Name() string {
	return "sharedprofile"
}

// StaticCredentialConfig contains settings for using a fixed set of
// credentials to authenticate.
type StaticCredentialConfig struct {
	ID     string `description:"The access key ID."`
	Secret string `descriptin:"The secret key."`
	Token  string `description:"Optional access token."`
}

// Name of the configuration root.
func (*StaticCredentialConfig) Name() string {
	return "static"
}

// AssumeRoleCredentialConfig contains settings for assuming a role.
type AssumeRoleCredentialConfig struct {
	Role       string `description:"The ARN of the role to assume."`
	ExternalID string `description:"External ID to use if using a cross-acount role."`
}

// Name of the configuration root.
func (*AssumeRoleCredentialConfig) Name() string {
	return "assumerole"
}

// SessionConfig contains settings for establishing an authenticated AWS session.
type SessionConfig struct {
	Region                     string `description:"The AWS region in which to authenticate."`
	Endpoint                   string `description:"Override the default AWS URL."`
	SharedCredentialConfig     *SharedCredentialConfig
	StaticCredentialConfig     *StaticCredentialConfig
	AssumeRoleCredentialConfig *AssumeRoleCredentialConfig
}

// Name of the configuration root.
func (*SessionConfig) Name() string {
	return "session"
}

// SessionComponent implements the Component interface and produces an AWS
// session object.
type SessionComponent struct{}

// NewSessionComponent constructs a default SessionComponent.
func NewSessionComponent() *SessionComponent {
	return &SessionComponent{}
}

// Settings generates the default configuration.
func (*SessionComponent) Settings() *SessionConfig {
	return &SessionConfig{
		SharedCredentialConfig:     &SharedCredentialConfig{},
		StaticCredentialConfig:     &StaticCredentialConfig{},
		AssumeRoleCredentialConfig: &AssumeRoleCredentialConfig{},
	}
}

// New generates a session for use with any AWS client.
func (*SessionComponent) New(ctx context.Context, conf *SessionConfig) (*session.Session, error) {
	ac := aws.NewConfig()
	if conf.Endpoint != "" {
		ac = ac.WithEndpoint(conf.Endpoint)
	}
	if conf.Region != "" {
		ac = ac.WithRegion(conf.Region)
	}
	var creds []credentials.Provider
	sf := conf.SharedCredentialConfig.File
	sp := conf.SharedCredentialConfig.Profile
	if sf != "" || sp != "" {
		creds = append(creds, &credentials.SharedCredentialsProvider{
			Filename: sf,
			Profile:  sp,
		})
	}
	stid := conf.StaticCredentialConfig.ID
	stsec := conf.StaticCredentialConfig.Secret
	sttok := conf.StaticCredentialConfig.Token
	if stid != "" {
		creds = append(creds, &credentials.StaticProvider{
			Value: credentials.Value{
				AccessKeyID:     stid,
				SecretAccessKey: stsec,
				SessionToken:    sttok,
			},
		})
	}
	if len(creds) < 1 {
		// If no special creds are set then use the default credential chain.
		creds = defaults.CredProviders(defaults.Config(), defaults.Handlers())
	}
	ac = ac.WithCredentials(credentials.NewChainCredentials(creds))
	ses, err := session.NewSession(ac)
	if err != nil {
		return nil, err
	}
	ar := conf.AssumeRoleCredentialConfig.Role
	aeid := conf.AssumeRoleCredentialConfig.ExternalID
	if ar != "" {
		p := &stscreds.AssumeRoleProvider{
			Client:   sts.New(ses),
			RoleARN:  ar,
			Duration: stscreds.DefaultDuration,
		}
		if aeid != "" {
			p.ExternalID = aws.String(aeid)
		}
		creds = append(creds, p)
	}
	ses.Config = ses.Config.WithCredentials(credentials.NewChainCredentials(creds))
	return ses, nil
}

// LoadSession is a convenience method for binding the source to the component.
func LoadSession(ctx context.Context, source settings.Source, c *SessionComponent) (*session.Session, error) {
	dst := new(session.Session)
	err := settings.NewComponent(ctx, source, c, dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// NewSession is the top-level entry point for creating an AWS auth session.
func NewSession(ctx context.Context, source settings.Source) (*session.Session, error) {
	return LoadSession(ctx, source, NewSessionComponent())
}
