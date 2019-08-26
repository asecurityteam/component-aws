package aws

import (
	"context"

	"github.com/asecurityteam/settings"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Config contains the settings for a S3 client.
type S3Config struct {
	Session *SessionConfig
}

// Name of the configuration root.
func (*S3Config) Name() string {
	return "s3"
}

// S3Component is used to create a new S3 client.
type S3Component struct {
	Session *SessionComponent
}

// NewS3Component generates a default component.
func NewS3Component() *S3Component {
	return &S3Component{
		Session: NewSessionComponent(),
	}
}

// Settings generates the default settings for the component.
func (c *S3Component) Settings() *S3Config {
	return &S3Config{
		Session: c.Session.Settings(),
	}
}

// New constructs a S3 client.
func (c *S3Component) New(ctx context.Context, conf *S3Config) (*s3.S3, error) {
	ses, err := c.Session.New(ctx, conf.Session)
	if err != nil {
		return nil, err
	}
	return s3.New(ses), nil
}

// LoadS3 is a convenience method for binding the source to the component.
func LoadS3(ctx context.Context, source settings.Source, c *S3Component) (*s3.S3, error) {
	dst := new(s3.S3)
	err := settings.NewComponent(ctx, source, c, dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// NewS3 is the top-level entry point for creating an S3 client.
func NewS3(ctx context.Context, source settings.Source) (*s3.S3, error) {
	return LoadS3(ctx, source, NewS3Component())
}
