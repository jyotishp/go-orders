package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func createAuthTable(tableName string) {
	panic("Implement createAuthTable!")
}

func CreateTable(tableName string) {

	svc := createSession()

	ip := &dynamodb.CreateTableInput{}

	if tableName == "Items" {
		ip = &dynamodb.CreateTableInput{
			TableName: aws.String(tableName),
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("ItemId"),
					AttributeType: aws.String("N"),
				},
				{
					AttributeName: aws.String("RestaurantId"),
					AttributeType: aws.String("N"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("RestaurantId"),
					KeyType: aws.String("HASH"),
				},
				{
					AttributeName: aws.String("ItemId"),
					KeyType: aws.String("RANGE"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits: aws.Int64(10),
				WriteCapacityUnits: aws.Int64(5),
			},
		}
	} else {
		ip = &dynamodb.CreateTableInput{
			TableName: aws.String(tableName),
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("Id"),
					AttributeType: aws.String("N"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("Id"),
					KeyType: aws.String("HASH"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits: aws.Int64(10),
				WriteCapacityUnits: aws.Int64(5),
			},
		}
	}

	_, err := svc.CreateTable(ip)
	if err != nil {
		printError(err)
		return
	}
	fmt.Println("Successfully created ", tableName, " table.")
}

