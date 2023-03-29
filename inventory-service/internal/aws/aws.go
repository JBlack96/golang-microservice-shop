package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// CreateSession creates a new AWS session.
func CreateSession(region, endpoint string) (sess *session.Session, err error) {
	// Create AWS Config.
	if len(region) == 0 {
		return nil, fmt.Errorf("no valid AWS_REGION found")
	}
	cfg := aws.NewConfig().WithRegion(region)

	// If there is an endpoint, override credentials for localstack integration tests
	if endpoint != "" {
		cfg.
			WithEndpoint(endpoint).
			WithCredentials(
				credentials.NewStaticCredentials("test", "test", "test-token"),
			)
	}

	// Create AWS Session.
	sess, err = session.NewSession(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating AWS session")
	}

	return
}
