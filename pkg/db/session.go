package db

import (
    "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "log"
)

// Create a new AWS session
func NewDbSession() *dynamodb.DynamoDB {
    dbEndpoint := GetEnv("DB_ENDPOINT", "localhost:8080")
    session, err := session.NewSession(&aws.Config{
        Endpoint: aws.String(dbEndpoint),
    })
    if err != nil {
        log.Fatalf("failed to create new session: %v", err)
    }
    svc := dynamodb.New(session)
    return svc
}