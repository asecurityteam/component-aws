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
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/require"
)

func TestS3(t *testing.T) {
	bucket := fmt.Sprintf("TEST%d", rand.Int63())
	region := os.Getenv("S3_REGION")
	endpoint := os.Getenv("S3_ENDPOINT")

	cmp := aws.NewS3Component()
	conf := cmp.Settings()
	conf.Session.Region = region
	conf.Session.Endpoint = endpoint
	client, err := cmp.New(context.Background(), conf)
	require.Nil(t, err)

	start := time.Now()
	for time.Since(start) < 10*time.Second {
		_, err = client.CreateBucket(&s3.CreateBucketInput{
			Bucket: awssdk.String(bucket),
		})
		if err == nil {
			break
		}
		t.Log(err)
	}
	require.Nil(t, err)
}
