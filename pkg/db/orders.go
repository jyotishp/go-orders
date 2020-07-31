package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/jyotishp/go-orders/pkg/models"
)

func GetOrder(tableName string, id int32) (models.Order, error) {
	type Input struct {
		Id int32
	}

	key, err := dynamodbattribute.MarshalMap(Input{Id: id})
	if err != nil {
		printError(err)
		return models.Order{}, err
	}

	ip := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: key,
	}

	svc := createSession()

	res, err := svc.GetItem(ip)
	if err != nil {
		printError(err)
		return models.Order{}, err
	}

	order := models.Order{}

	err = dynamodbattribute.UnmarshalMap(res.Item, &order)
	if err != nil {
		printError(err)
		return models.Order{}, err
	}

	return order, nil
}

func InsertOrder(tableName string, createOrder models.Order) (models.Order, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		printError(err)
		return models.Order{}, err
	}

	createOrder.Id = int32(uid.ID())
	ip, err := dynamodbattribute.MarshalMap(createOrder)
	if err != nil {
		printError(err)
		return models.Order{}, nil
	}

	svc := createSession()
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: ip,
	}

	_, err = svc.PutItem(input)
	if err != nil {
		printError(err)
		return models.Order{}, nil
	}

	return createOrder, nil
}