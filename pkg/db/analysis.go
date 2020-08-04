package db

import (
	"go-orders/src/github.com/aws/aws-sdk-go/aws"
	"go-orders/src/github.com/aws/aws-sdk-go/service/dynamodb"
	"go-orders/src/github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetTopRestaurants(tableName string) ([]Restaurant, error) {
	ip := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
		ProjectionExpression: aws.String("Id, OrderCount"),
		IndexName: aws.String("OrderCounts"),
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
