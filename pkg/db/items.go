package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

func GetItem(tableName string, restaurantId int32, itemId int32) (Item, error) {
	if !checkTable(tableName) {
		CreateItemsTable(tableName)
	}
	type Input struct {
		ItemId int32
		RestaurantId int32
	}

	key, err := dynamodbattribute.MarshalMap(Input{ItemId: itemId, RestaurantId: restaurantId})
	if err != nil {
		printError(err)
		return Item{}, err
	}

	ip := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: key,
	}
	svc := createSession()
	res, err := svc.GetItem(ip)
	if err != nil {
		printError(err)
		return Item{}, err
	}
	opItem := Item{}
	err = dynamodbattribute.UnmarshalMap(res.Item, &opItem)
	if err != nil {
		printError(err)
		return Item{}, err
	}
	if opItem == (Item{}) {
		return opItem, fmt.Errorf("no item found for given restaurantId and itemId")
	}
	opItem.Id = itemId
	return opItem, nil
}

func InsertItem(tableName string, restaurantId int32, item Item, updateRestaurants bool) (Item ,error) {
	if !checkTable(tableName) {
		CreateItemsTable(tableName)
	}

	uid, err := uuid.NewUUID()
	if err != nil {
		printError(err)
		return Item{}, err
	}

	item.Id = int32(uid.ID())

	ip, err := dynamodbattribute.MarshalMap(createItem(restaurantId, item))
	if err != nil {
		printError(err)
		return Item{}, err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: ip,
	}

	svc := createSession()
	_, err = svc.PutItem(input)
	if err != nil {
		printError(err)
		return Item{}, err
	}
	if updateRestaurants {
		restaurant, err := GetRestaurant("Restaurants", restaurantId)
		if err != nil {
			printError(err)
			return Item{}, err
		}
		restaurant.Items = append(restaurant.Items, item)
		_, err = UpdateRestaurant("Restaurants", restaurant, false)
		if err != nil {
			printError(err)
			return Item{}, err
		}
	}
	return item, nil
}

func UpdateItem(tableName string, restaurantId int32, item Item) (Item, error) {
	if !checkTable(tableName) {
		CreateItemsTable(tableName)
	}

	type KeyInput struct {
		RestaurantId int32
		ItemId int32
	}

	key, err := dynamodbattribute.MarshalMap(KeyInput{ItemId: item.Id, RestaurantId: restaurantId})
	if err != nil {
		printError(err)
		return Item{}, err
	}

	imap, err := dynamodbattribute.MarshalMap(updateItemMap(createItem(restaurantId, item)))

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: key,
		ExpressionAttributeValues: imap,
		UpdateExpression: aws.String("set Cuisine=:c, Amount=:a, Discount=:d"),
	}

	svc := createSession()
	_, err = svc.UpdateItem(input)
	if err != nil {
		printError(err)
		return Item{}, err
	}
	restaurant, err := GetRestaurant("Restaurants", restaurantId)
	if err != nil {
		printError(err)
		return Item{}, err
	}

	for idx, itm := range restaurant.Items {
		if itm.Id == item.Id {
			restaurant.Items[idx] = item
		}
	}
	_, err = UpdateRestaurant("Restaurants", restaurant, false)
	if err != nil {
		printError(err)
		return Item{}, err
	}

	return item, nil
}

func DeleteItem(tableName string, restaurantId, itemId int32) error {
	if !checkTable(tableName) {
		CreateItemsTable(tableName)
	}

	type KeyInput struct {
		RestaurantId, ItemId int32
	}

	key, err := dynamodbattribute.MarshalMap(KeyInput{RestaurantId: restaurantId, ItemId: itemId})
	if err != nil {
		printError(err)
		return err
	}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: key,
	}

	restaurant, err := GetRestaurant("Restaurants", restaurantId)
	if err != nil {
		printError(err)
		return err
	}
	itms := make([]Item, 0)
	for _, itm := range restaurant.Items {
		if itm.Id != itemId {
			itms = append(itms, itm)
		}
	}
	restaurant.Items = itms
	_, err = UpdateRestaurant("Restaurants", restaurant, false)
	if err != nil {
		printError(err)
		return err
	}

	svc := createSession()
	_, err = svc.DeleteItem(input)
	return err
}
