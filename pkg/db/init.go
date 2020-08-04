package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Create DynamoDB service
func CreateDbService() *dynamodb.DynamoDB {
	return NewDbSession()
}

// Initialize the database
// Create new DynamoDB service and initializes the tables
func InitDb() {
	svc := CreateDbService()
	printError(CreateTableIfNotExists(svc, OrdersTableSchema()))
	printError(CreateTableIfNotExists(svc, RestaurantsTableSchema()))
	printError(CreateTableIfNotExists(svc, CustomersTableSchema()))
	printError(CreateTableIfNotExists(svc, ItemsTableSchema()))
}

// Print the error encountered during DB operations
func printError(err error) {
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			case dynamodb.ErrCodeResourceInUseException:
				fmt.Println(dynamodb.ErrCodeResourceInUseException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
}
