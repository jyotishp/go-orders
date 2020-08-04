package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
    "log"
)

var (
	// Tables
	OrdersTable = "Orders"
	RestaurantsTable = "Restaurants"
	CustomersTable = "Customers"
	ItemsTable = "Items"

	// Keys
	orderId          = &dynamodb.KeySchemaElement{AttributeName: aws.String("OrderId"), KeyType: aws.String("HASH")}
	restaurantId     = &dynamodb.KeySchemaElement{AttributeName: aws.String("RestaurantId"), KeyType: aws.String("HASH")}
	customerId       = &dynamodb.KeySchemaElement{AttributeName: aws.String("CustomerId"), KeyType: aws.String("HASH")}
	itemId           = &dynamodb.KeySchemaElement{AttributeName: aws.String("ItemId"), KeyType: aws.String("HASH")}
	cuisine          = &dynamodb.KeySchemaElement{AttributeName: aws.String("Cuisine"), KeyType: aws.String("RANGE")}
	state            = &dynamodb.KeySchemaElement{AttributeName: aws.String("State"), KeyType: aws.String("RANGE")}

	// Attributes
	orderIdAttr      = &dynamodb.AttributeDefinition{AttributeName: aws.String("OrderId"), AttributeType: aws.String("S")}
	restaurantIdAttr = &dynamodb.AttributeDefinition{AttributeName: aws.String("RestaurantId"), AttributeType: aws.String("S")}
	customerIdAttr   = &dynamodb.AttributeDefinition{AttributeName: aws.String("CustomerId"), AttributeType: aws.String("S")}
	itemIdAttr       = &dynamodb.AttributeDefinition{AttributeName: aws.String("ItemId"), AttributeType: aws.String("S")}
	stateAttr        = &dynamodb.AttributeDefinition{AttributeName: aws.String("State"), AttributeType: aws.String("S")}
	cuisineAttr      = &dynamodb.AttributeDefinition{AttributeName: aws.String("Cuisine"), AttributeType: aws.String("S")}
)

// Schema for orders table
func OrdersTableSchema() *dynamodb.CreateTableInput {
	table := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			orderIdAttr, restaurantIdAttr, customerIdAttr, stateAttr,
		},
		KeySchema: []*dynamodb.KeySchemaElement{orderId},
		LocalSecondaryIndexes: []*dynamodb.LocalSecondaryIndex{
			{
				IndexName: restaurantId.AttributeName,
				KeySchema: []*dynamodb.KeySchemaElement{restaurantId, state},
			},
			{
				IndexName: customerId.AttributeName,
				KeySchema: []*dynamodb.KeySchemaElement{customerId, state},
			},
		},
		TableName: aws.String(ordersTable),
	}
	return table
}

// Schema for customers table
func CustomersTableSchema() *dynamodb.CreateTableInput {
	table := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{restaurantIdAttr, stateAttr},
		KeySchema:            []*dynamodb.KeySchemaElement{customerId},
		TableName:            aws.String(customersTable),
	}
	return table
}

// Schema for restaurants table
func RestaurantsTableSchema() *dynamodb.CreateTableInput {
	table := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{restaurantIdAttr, stateAttr},
		KeySchema:            []*dynamodb.KeySchemaElement{restaurantId},
		TableName:            aws.String(restaurantsTable),
	}
	return table
}

// Schema for items table
func ItemsTableSchema() *dynamodb.CreateTableInput {
	table := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{itemIdAttr, cuisineAttr},
		KeySchema:            []*dynamodb.KeySchemaElement{itemId, cuisine},
		TableName:            aws.String(itemsTable),
	}
	return table
}

// Create a new table
func CreateTable(svc dynamodbiface.DynamoDBAPI, table *dynamodb.CreateTableInput) error {
	out, err := svc.CreateTable(table)
	if err != nil {
		return fmt.Errorf("unable to create table: %v", err)
	}
	log.Printf("created table: %v", out)
	return nil
}

// Check if the given table exists
func TableExists(svc dynamodbiface.DynamoDBAPI, table string) bool {
	_, err := svc.DescribeTable(&dynamodb.DescribeTableInput{TableName: aws.String(table)})
	if err != nil {
		return false
	}
	return true
}

// Create a table if it doesn't exist
func CreateTableIfNotExists(svc dynamodbiface.DynamoDBAPI, table *dynamodb.CreateTableInput) error {
	if !TableExists(svc, *table.TableName) {
		return CreateTable(svc, table)
	}
	return nil
}
