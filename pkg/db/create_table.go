package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func createAuthTable(tableName string) {
	panic("Implement createAuthTable!")
}

func CreateRestaurantsTable(tableName string) error {
	svc := createSession()
	key1 := []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("Name"),
			KeyType: aws.String("HASH"),
		},
	}
	key2 := []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("OrderCount"),
			KeyType: aws.String("HASH"),
		},
	}
	sgi := []*dynamodb.GlobalSecondaryIndex{
		{
			IndexName: aws.String("RName"),
			KeySchema: key1,
			Projection: &dynamodb.Projection{
				ProjectionType: aws.String("ALL"),
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(10),
				WriteCapacityUnits: aws.Int64(5),
			},
		},
		{
			IndexName: aws.String("OrderCounts"),
			KeySchema: key2,
			Projection: &dynamodb.Projection{
				ProjectionType: aws.String("ALL"),
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(10),
				WriteCapacityUnits: aws.Int64(5),
			},
		},
	}

	ip := &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("Name"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("OrderCount"),
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
		GlobalSecondaryIndexes: sgi,
	}

	_, err := svc.CreateTable(ip)
	if err != nil {
		printError(err)
		return err
	}
	fmt.Println("Successfully created ", tableName, " table.")
	return nil
}

func CreateItemsTable(tableName string) error {
	svc := createSession()
	keys := []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("RestaurantId"),
			KeyType: aws.String("HASH"),
		},
	}
	sgi := []*dynamodb.GlobalSecondaryIndex{
		{
			IndexName: aws.String("RId"),
			KeySchema: keys,
			Projection: &dynamodb.Projection{
				ProjectionType: aws.String("ALL"),
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(10),
				WriteCapacityUnits: aws.Int64(5),
			},
		},
	}
	ip := &dynamodb.CreateTableInput{
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
		GlobalSecondaryIndexes: sgi,
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits: aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := svc.CreateTable(ip)
	if err != nil {
		printError(err)
		return err
	}
	fmt.Println("Successfully created ", tableName, " table.")
	return nil
}

func createCustomersTable(tableName string) error {
	svc := createSession()
	ip := &dynamodb.CreateTableInput{
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

	_, err := svc.CreateTable(ip)
	if err != nil {
		printError(err)
		return err
	}
	fmt.Println("Successfully created ", tableName, " table.")
	return nil
}

func createOrdersTable(tableName string) error {
	svc := createSession()

	ip := &dynamodb.CreateTableInput{
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

	_, err := svc.CreateTable(ip)
	if err != nil {
		printError(err)
		return err
	}
	fmt.Println("Successfully created ", tableName, " table.")
	return nil
}

func CreateTable(tableName string) error {

	switch tableName {
	case "Restaurants":
		return CreateRestaurantsTable(tableName)
	case "Items":
		return CreateItemsTable(tableName)
	case "Orders":
		return createOrdersTable(tableName)
	case "Customers":
		return createCustomersTable(tableName)
	default:
		return fmt.Errorf("illegal table name entered")
	}
}

