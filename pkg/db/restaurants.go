package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
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

func GetRestaurantName(tableName string, restaurantName string) ([]models.Restaurant, error) {
	filter := expression.Name("Name").Equal(expression.Value(restaurantName))
	proj := expression.NamesList(expression.Name("Id"),
		expression.Name("Name"), expression.Name("Address"), expression.Name("Items"))
	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(proj).Build()
	restaurantList := make([]models.Restaurant, 0)
	if err != nil {
		printError(err)
		return restaurantList, err
	}
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
		ExpressionAttributeNames: expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression: expr.Filter(),
		ProjectionExpression: expr.Projection(),
	}
	svc := createSession()
	res, err := svc.Scan(params)
	for _, item := range res.Items {
		restaurant := models.Restaurant{}
		err := dynamodbattribute.UnmarshalMap(item, &restaurant)
		if err != nil {
			printError(err)
			return restaurantList, err
		}
		restaurantList = append(restaurantList, restaurant)
	}
	return restaurantList, nil
}

func GetRestaurantSGI(tableName string, restaurantName string) ([]models.Restaurant, error) {
	restaurantList := make([]models.Restaurant, 0)
	ip := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		IndexName: aws.String("RName"),
		KeyConditionExpression: aws.String("#n=:n"),
		ExpressionAttributeNames: map[string]*string{
			"#n": aws.String("Name"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {
				S: aws.String(restaurantName),
			},
		},
	}
	svc := createSession()
	res, err := svc.Query(ip)
	if err != nil {
		printError(err)
		return restaurantList, err
	}
	for _, item := range res.Items {
		restaurant := models.Restaurant{}
		err := dynamodbattribute.UnmarshalMap(item, &restaurant)
		if err != nil {
			printError(err)
			return restaurantList, err
		}
		restaurantList = append(restaurantList, restaurant)
	}
	return restaurantList, nil
}

func InsertRestaurant(tableName string, createRestaurant models.Restaurant) (models.Restaurant, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		printError(err)
		return models.Restaurant{}, err
	}

	createRestaurant.Id = int32(uid.ID())
	ip, err := dynamodbattribute.MarshalMap(buildRestaurant(createRestaurant))
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
	createRestaurant.Items = insertItems(createRestaurant.Id, createRestaurant.Items, false)
	_, err = UpdateRestaurant(tableName, createRestaurant, false)
	if err != nil {
		printError(err)
		return models.Restaurant{}, err
	}
	return createRestaurant, nil
	
}

func UpdateRestaurant(tableName string, updateRestaurant models.Restaurant, updateItems bool) (models.Restaurant, error) {
	type KeyInput struct {
		Id int32
	}

	key, err := dynamodbattribute.MarshalMap(KeyInput{Id: updateRestaurant.Id})
	if err != nil {
		printError(err)
		return models.Restaurant{}, err
	}

	if updateItems {
		updateRestaurant.Items = insertItems(updateRestaurant.Id, updateRestaurant.Items, false)
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

func GetAllItems(tableName string, filter models.ItemFilter) ([]models.Item, error) {
	items := make([]models.Item, 0)
	restaurant, err := GetRestaurant(tableName, filter.RestaurantId)
	if err != nil {
		printError(err)
		return items, err
	}

	if filter.Min == 0 && filter.Max == 0 {
		return restaurant.Items, nil
	}

	for _, item := range restaurant.Items {
		if item.Amount >= filter.Min && item.Amount <= filter.Max {
			items = append(items, item)
		}
	}
	return items, nil

}
