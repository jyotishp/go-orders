package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/jyotishp/go-orders/pkg/models"
)

func GetRestaurant(tableName string, id int32) (models.Restaurant, error) {
	type Input struct {
		Id int32
	}

	key, err := dynamodbattribute.MarshalMap(Input{Id: id})
	if err != nil {
		printError(err)
		return models.Restaurant{}, err
	}

	ip := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: key,
	}

	svc := createSession()

	res, err := svc.GetItem(ip)
	if err != nil {
		printError(err)
		return models.Restaurant{}, err
	}

	restaurant := models.Restaurant{}

	err = dynamodbattribute.UnmarshalMap(res.Item, &restaurant)
	if err != nil {
		printError(err)
		return models.Restaurant{}, err
	}

	return restaurant, nil
}

func InsertRestaurant(tableName string, createRestaurant models.Restaurant) (models.Restaurant, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		printError(err)
		return models.Restaurant{}, err
	}

	createRestaurant.Id = int32(uid.ID())
	ip, err := dynamodbattribute.MarshalMap(createRestaurant)
	if err != nil {
		printError(err)
		return models.Restaurant{}, nil
	}

	svc := createSession()
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: ip,
	}

	_, err = svc.PutItem(input)
	if err != nil {
		printError(err)
		return models.Restaurant{}, nil
	}

	return createRestaurant, nil
	
}

func UpdateRestaurant(tableName string, updateRestaurant models.Restaurant) (models.Restaurant, error) {
	type KeyInput struct {
		Id int32
	}

	key, err := dynamodbattribute.MarshalMap(KeyInput{Id: updateRestaurant.Id})
	if err != nil {
		printError(err)
		return models.Restaurant{}, err
	}

	rmap, err := dynamodbattribute.MarshalMap(restaurantMap(updateRestaurant))
	if err != nil {
		printError(err)
		return models.Restaurant{}, err
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: key,
		ExpressionAttributeValues: rmap,
		ExpressionAttributeNames: map[string]*string{
			"#n": aws.String("Name"),
			"#itms": aws.String("Items"),
		},
		UpdateExpression: aws.String("set #n=:rn, Address=:radr, #itms=:ritms"),
	}

	svc := createSession()
	_, err = svc.UpdateItem(input)
	if err != nil {
		printError(err)
		return models.Restaurant{}, err
	}

	return updateRestaurant, nil
}
