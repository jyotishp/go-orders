package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jyotishp/go-orders/pkg/models"
)

func DeleteRestaurant(tableName string, restaurantId int32) error {
	svc := createSession()
	items, err := GetAllItems("Restaurants", models.ItemFilter{RestaurantId: restaurantId})
	if err != nil {
		printError(err)
		return err
	}
	type DeleteKey struct {
		RestaurantId, ItemId int32
	}
	req := make([]*dynamodb.WriteRequest, 0)
	for _, item := range items {
		tmpKey, err := dynamodbattribute.MarshalMap(DeleteKey{
			RestaurantId: restaurantId, ItemId: item.Id})
		if err != nil {
			printError(err)
			return err
		}
		tmp := &dynamodb.WriteRequest{
			DeleteRequest: &dynamodb.DeleteRequest{
				Key: tmpKey,
			},
		}
		req = append(req, tmp)
	}

	ip := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			"Items": req,
		},
	}
	_, err = svc.BatchWriteItem(ip)
	if err != nil {
		printError(err)
		return err
	}

	type KeyInput struct {
		Id int32
	}

	key, err := dynamodbattribute.MarshalMap(KeyInput{Id: restaurantId})
	if err != nil {
		printError(err)
		return err
	}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: key,
	}

	_, err = svc.DeleteItem(input)
	if err != nil {
		printError(err)
		return err
	}
	return nil
}

func DeleteKey(tableName string, id int32) error {
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
