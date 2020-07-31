package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jyotishp/go-orders/pkg/models"
	"github.com/google/uuid"
)

func GetCustomer(tableName string, id int32) (models.Customer, error) {
	type Input struct {
		Id int32
	}

	key, err := dynamodbattribute.MarshalMap(Input{Id: id})
	if err != nil {
		printError(err)
		return models.Customer{}, err
	}

	ip := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: key,
	}

	svc := createSession()

	res, err := svc.GetItem(ip)
	if err != nil {
		printError(err)
		return models.Customer{}, err
	}

	customer := models.Customer{}

	err = dynamodbattribute.UnmarshalMap(res.Item, &customer)
	if err != nil {
		printError(err)
		return models.Customer{}, err
	}

	return customer, nil
}

func CreateCustomer(tableName string, createCustomer models.Customer) (models.Customer, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		printError(err)
		return models.Customer{}, err
	}

	createCustomer.Id = int32(uid.ID())
	ip, err := dynamodbattribute.MarshalMap(createCustomer)
	if err != nil {
		printError(err)
		return models.Customer{}, nil
	}

	svc := createSession()
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: ip,
	}

	_, err = svc.PutItem(input)
	if err != nil {
		printError(err)
		return models.Customer{}, nil
	}

	return createCustomer, nil
}
