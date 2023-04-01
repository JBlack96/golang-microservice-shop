package setup

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSetup(t *testing.T) {
	time.Sleep(30 * time.Second)
	t.Run("create dynamodb table - inventory", func(t *testing.T) {
		cfg := aws.NewConfig().
			WithRegion("eu-west-1").
			WithEndpoint("http://localhost:4566").
			WithCredentials(
				credentials.NewStaticCredentials("test", "test", "test-token"),
			)
		sess := session.Must(session.NewSession(cfg))
		db := dynamodb.New(sess)

		createTableInput := dynamodb.CreateTableInput{
			TableName: aws.String("inventory"),
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{AttributeName: aws.String("id"), AttributeType: aws.String("S")},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{AttributeName: aws.String("id"), KeyType: aws.String("HASH")},
			},
			BillingMode: aws.String("PAY_PER_REQUEST"),
		}

		if _, err := db.CreateTable(&createTableInput); err != nil {
			assert.FailNow(t, "failed to create dynamodb inventory table", err)
		}
	})
}
