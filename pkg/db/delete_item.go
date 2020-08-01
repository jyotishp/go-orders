package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func DeleteItem(tableName string, id int32) error {
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
