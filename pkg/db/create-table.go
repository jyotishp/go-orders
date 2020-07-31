package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func createOrderTable(tableName string) {
	svc := createSession()

	ip := &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("OrderId"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("OrderId"),
				KeyType: aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits: aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := svc.CreateTable(ip)
	if err != nil {
		checkError(err)
		return
	}
	fmt.Println("Successfully created ", tableName, " table.")
}

func createRestaurantTable(tableName string) {
	svc := createSession()

	ip := &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
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
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits: aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := svc.CreateTable(ip)
	if err != nil {
		checkError(err)
		return
	}
	fmt.Println("Successfully created ", tableName, " table.")
}

func createCustomerTable(tableName string) {
	svc := createSession()

	ip := &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("CustomerId"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("CustomerId"),
				KeyType: aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits: aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := svc.CreateTable(ip)
	if err != nil {
		checkError(err)
		return
	}
	fmt.Println("Successfully created ", tableName, " table.")
}

func createAuthTable(tableName string) {
	panic("Implement createAuthTable!")
}

func CreateTable(tableName string) {

	switch tableName {
	case "Orders":
		createOrderTable(tableName)
	case "Restaurants":
		createRestaurantTable(tableName)
	case "Customers":
		createCustomerTable(tableName)
	case "Auth":
		createAuthTable(tableName)
	default:
		log.Fatal("Wrong tableName passed")
	}
}

