package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jyotishp/go-orders/pkg/models"
	"log"
)

func PutItem(tableName string, model interface{})  {
	if ok := checkTable(tableName); !ok {
		log.Fatal("Invalid Table Name passed to Put Item")
		return
	}

	var item models.Model
	var ok bool

	switch tableName {
	case "Customers":
		item, ok = model.(models.Customer)
	case "Orders":
		item, ok = model.(models.Order)
	case "Restaurants":
		item, ok = model.(models.Restaurant)
	}

	if !ok {
		log.Fatal("Wrong Request received in Put Item")
		return
	}

	svc := createSession()

	ip, err := dynamodbattribute.MarshalMap(item)
	checkError(err)

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: ip,
	}

	_, err = svc.PutItem(input)
	checkError(err)
}
