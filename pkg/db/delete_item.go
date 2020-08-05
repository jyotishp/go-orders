package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// DeleteKey accepts an Id and and table name and deletes the entry in the db.
// This function onlly caters to the Customers and Orders tables.
func DeleteKey(tableName string, id int32) error {
	if !checkTable(tableName) {
		if tableName == "Customers" {
			createCustomersTable(tableName)
		} else if tableName == "Orders" {
			createOrdersTable(tableName)
		}
	}

	type KeyInput struct {
		Id int32
	}

	key, err := dynamodbattribute.MarshalMap(KeyInput{Id: id})
	if err != nil {
		printError(err)
		return err
	}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: key,
	}

	svc := createSession()
	_, err = svc.DeleteItem(input)
	if err != nil {
		printError(err)
		return err
	}

	return nil
}

// DeleteTable deletes the restaurant of the given table name.
func DeleteTable(tableName string) error {
	input := &dynamodb.DeleteTableInput{
		TableName: aws.String(tableName),
	}
	svc := createSession()

	_, err := svc.DeleteTable(input)
	if err != nil {
		printError(err)
		return err
	}
	fmt.Println("Succesfully deleted" + tableName + " table.")
	return nil
}
