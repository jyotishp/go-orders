package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"log"
	"os"
)


// Get value of an env variable and return a default if it doesn't
// exist
func GetEnv(key, defaultVal string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultVal
	}
	return value
}

// Add an object to the database
func AddObject(svc dynamodbiface.DynamoDBAPI, table string, data interface{}) error {
	av, err := dynamodbattribute.MarshalMap(data)
	if err != nil {
		return fmt.Errorf("failed to marshal new customer: %v", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(table),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return fmt.Errorf("failed to add item to database: %v", err)
	}
	return nil
}

func GetObjectById(svc dynamodbiface.DynamoDBAPI, table, key, value string) (map[string]*dynamodb.AttributeValue, error) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			key: {
				S: aws.String(value),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("could not fetch from db: %v", err)
	}
	if result.Item == nil {
		return nil, fmt.Errorf("entity not found: %v", err)
	}
	log.Printf("%v, %v", result, result.Item)
	return result.Item, nil
}

func UpdateObjectById(
	svc dynamodbiface.DynamoDBAPI,
	table, key, value string,
	attrValues map[string]*dynamodb.AttributeValue,
	attrNames map[string]*string,
	updateExpr string,
	) error {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: attrValues,
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			key: {
				S: aws.String(value),
			},
		},
		ExpressionAttributeNames:     attrNames,
		UpdateExpression: aws.String(updateExpr),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		return fmt.Errorf("faield to update object: %v", err)
	}
	return nil
}

func DeleteObjectById(svc dynamodbiface.DynamoDBAPI, table, key, value string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			key: {
				S: aws.String(value),
			},
		},
		TableName: aws.String(table),
	}
	_, err := svc.DeleteItem(input)
	if err != nil {
		return fmt.Errorf("failed to delete object: %v", err)
	}
	return nil
}
//
//func createItem(restaurantId int32, item Item) ItemIp {
//	return ItemIp{
//		RestaurantId: restaurantId,
//		Name: item.Name,
//		ItemId: item.Id,
//		Cuisine: item.Cuisine,
//		Discount: item.Discount,
//		Amount: item.Amount,
//	}
//}
//
//func updateItemMap(item ItemIp) ItemUp {
//	return ItemUp{
//		Cuisine: item.Cuisine,
//		Discount: item.Discount,
//		Amount: item.Amount,
//	}
//}
//
//func insertItems(restaurantId int32, items []Item, updateRestaurants bool) []Item {
//	op := make([]Item, 0)
//	for _, item := range items {
//		opItem, err := InsertItem("Items", restaurantId, item, updateRestaurants)
//		if err != nil {
//			fmt.Println("REACH4")
//			printError(err)
//			return op
//		}
//		op = append(op, opItem)
//	}
//	return op
//}
//
//func removeCustId(customer Customer) CustomerNoId {
//	return CustomerNoId{
//		Name: customer.Name,
//		Address: customer.Address,
//	}
//}
//
//func toNormalOrder(order OrderIp) Order {
//	return Order{
//		Id:            order.Id,
//		Discount:      order.Discount,
//		Amount:        order.Amount,
//		PaymentMethod: order.PaymentMethod,
//		Rating:        order.Rating,
//		Duration:      order.Duration,
//		Cuisine:       order.Cuisine,
//		Time:          order.Time,
//		Verified:      order.Verified,
//		Customer:      order.Customer,
//		Restaurant:    order.Restaurant,
//		Items:         order.Items,
//	}
//}
//
//func toNormalRestaurant(restaurant RestaurantNoId) Restaurant {
//	return Restaurant{
//		Name:    restaurant.Name,
//		Address: restaurant.Address,
//		Items:   restaurant.Items,
//	}
//}
//
//func removeOrderId(order OrderIp) OrderNoId  {
//	return OrderNoId{
//		Discount:      order.Discount,
//		Amount:        order.Amount,
//		PaymentMethod: order.PaymentMethod,
//		Rating:        order.Rating,
//		Duration:      order.Duration,
//		Cuisine:       order.Cuisine,
//		Time:          order.Time,
//		Verified:      order.Verified,
//		Customer:      order.Customer,
//		Restaurant:    order.Restaurant,
//		Items:         order.Items,
//	}
//}
//
//func removeRestId(restaurant Restaurant) RestaurantNoId {
//	return RestaurantNoId{
//		Name:    restaurant.Name,
//		Address: restaurant.Address,
//		Items:   restaurant.Items,
//	}
//}