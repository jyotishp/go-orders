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
	orderId          = &dynamodb.KeySchemaElement{AttributeName: aws.String("order_id"), KeyType: aws.String("HASH")}
	restaurantId     = &dynamodb.KeySchemaElement{AttributeName: aws.String("restaurant_id"), KeyType: aws.String("HASH")}
	customerId       = &dynamodb.KeySchemaElement{AttributeName: aws.String("customer_id"), KeyType: aws.String("HASH")}
	itemId           = &dynamodb.KeySchemaElement{AttributeName: aws.String("item_id"), KeyType: aws.String("HASH")}
	cuisine          = &dynamodb.KeySchemaElement{AttributeName: aws.String("cuisine"), KeyType: aws.String("RANGE")}
	state            = &dynamodb.KeySchemaElement{AttributeName: aws.String("state"), KeyType: aws.String("RANGE")}

	// Attributes
	orderIdAttr      = &dynamodb.AttributeDefinition{AttributeName: aws.String("order_id"), AttributeType: aws.String("S")}
	restaurantIdAttr = &dynamodb.AttributeDefinition{AttributeName: aws.String("restaurant_id"), AttributeType: aws.String("S")}
	customerIdAttr   = &dynamodb.AttributeDefinition{AttributeName: aws.String("customer_id"), AttributeType: aws.String("S")}
	itemIdAttr       = &dynamodb.AttributeDefinition{AttributeName: aws.String("item_id"), AttributeType: aws.String("S")}
	stateAttr        = &dynamodb.AttributeDefinition{AttributeName: aws.String("state"), AttributeType: aws.String("S")}
	cuisineAttr      = &dynamodb.AttributeDefinition{AttributeName: aws.String("cuisine"), AttributeType: aws.String("S")}
)

// Schema for orders table
func OrdersTableSchema() *dynamodb.CreateTableInput {
	table := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			orderIdAttr, restaurantIdAttr, customerIdAttr,
		},
		KeySchema: []*dynamodb.KeySchemaElement{orderId},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: restaurantId.AttributeName,
				KeySchema: []*dynamodb.KeySchemaElement{restaurantId},
				Projection: &dynamodb.Projection{ProjectionType: aws.String("ALL")},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(10),
				},
			},
			{
				IndexName: customerId.AttributeName,
				KeySchema: []*dynamodb.KeySchemaElement{customerId},
				Projection: &dynamodb.Projection{ProjectionType: aws.String("ALL")},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(10),
				},
			},
		},
		TableName: aws.String(OrdersTable),
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	}
	return table
}

// Schema for customers table
func CustomersTableSchema() *dynamodb.CreateTableInput {
	table := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{customerIdAttr},
		KeySchema:            []*dynamodb.KeySchemaElement{customerId},
		TableName:            aws.String(CustomersTable),
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	}
	return table
}

// Schema for restaurants table
func RestaurantsTableSchema() *dynamodb.CreateTableInput {
	table := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{restaurantIdAttr},
		KeySchema:            []*dynamodb.KeySchemaElement{restaurantId},
		TableName:            aws.String(RestaurantsTable),
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	}
	return table
}

// Schema for items table
func ItemsTableSchema() *dynamodb.CreateTableInput {
	table := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{itemIdAttr, cuisineAttr},
		KeySchema:            []*dynamodb.KeySchemaElement{itemId, cuisine},
		TableName:            aws.String(ItemsTable),
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
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
	log.Printf("checking if table %v exists", table)
	_, err := svc.DescribeTable(&dynamodb.DescribeTableInput{TableName: aws.String(table)})
	if err != nil {
		return false
	}
	return true
}

// Create a table if it doesn't exist
func CreateTableIfNotExists(svc dynamodbiface.DynamoDBAPI, table *dynamodb.CreateTableInput) error {
	if !TableExists(svc, *table.TableName) {
		log.Printf("creating table: %v", *table.TableName)
		return CreateTable(svc, table)
	}
	return nil
}
