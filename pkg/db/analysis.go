package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetTopRestaurants(tableName string, size int32) ([]Restaurant, error) {
	ip := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		KeyConditionExpression: aws.String("Dummy=:z"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":z": {
				N: aws.String(fmt.Sprintf("0")),
			},
		},
		IndexName: aws.String("OrderCounts"),
		Limit: aws.Int64(int64(size)),
		ScanIndexForward: aws.Bool(false),
	}
	restaurantList := make([]Restaurant, 0)
	svc := createSession()
	res, err := svc.Query(ip)
	if err != nil {
		printError(err)
		return restaurantList, err
	}
	for _, rest := range res.Items {
		tmp := Restaurant{}
		err := dynamodbattribute.UnmarshalMap(rest, &tmp)
		if err != nil {
			printError(err)
			return restaurantList, err
		}
		restaurantList = append(restaurantList, tmp)
	}
	return restaurantList, nil
}

func GetWorstRestaurants(tableName string, size int32) ([]Restaurant, error) {
	ip := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		KeyConditionExpression: aws.String("Dummy=:z"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":z": {
				N: aws.String(fmt.Sprintf("0")),
			},
		},
		IndexName: aws.String("OrderCounts"),
		Limit: aws.Int64(int64(size)),
		ScanIndexForward: aws.Bool(true),
	}
	restaurantList := make([]Restaurant, 0)
	svc := createSession()
	res, err := svc.Query(ip)
	if err != nil {
		printError(err)
		return restaurantList, err
	}
	for _, rest := range res.Items {
		tmp := Restaurant{}
		err := dynamodbattribute.UnmarshalMap(rest, &tmp)
		if err != nil {
			printError(err)
			return restaurantList, err
		}
		restaurantList = append(restaurantList, tmp)
	}
	return restaurantList, nil
}
