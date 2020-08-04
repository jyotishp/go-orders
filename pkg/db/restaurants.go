package db
//
//import (
//	"fmt"
//	"github.com/aws/aws-sdk-go/aws"
//	"github.com/aws/aws-sdk-go/service/dynamodb"
//	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
//	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
//	"github.com/google/uuid"
//	pb "github.com/jyotishp/go-orders/pkg/proto"
//)
//
//func GetRestaurant(tableName string, id int32) (Restaurant, error) {
//	type Input struct {
//		Id int32
//	}
//
//	key, err := dynamodbattribute.MarshalMap(Input{Id: id})
//	if err != nil {
//		printError(err)
//		return Restaurant{}, err
//	}
//
//	ip := &dynamodb.GetItemInput{
//		TableName: aws.String(tableName),
//		Key: key,
//	}
//
//	svc := createSession()
//
//	res, err := svc.GetItem(ip)
//	if err != nil {
//		printError(err)
//		return Restaurant{}, err
//	}
//
//	restaurant := Restaurant{}
//
//	err = dynamodbattribute.UnmarshalMap(res.Item, &restaurant)
//	if err != nil {
//		printError(err)
//		return Restaurant{}, err
//	}
//	if restaurant.Id == 0 {
//		return Restaurant{}, fmt.Errorf("no restaurant found for given id")
//	}
//	return restaurant, nil
//}
//
//func GetRestaurantName(tableName string, restaurantName string) ([]Restaurant, error) {
//	filter := expression.Name("Name").Equal(expression.Value(restaurantName))
//	proj := expression.NamesList(expression.Name("Id"),
//		expression.Name("Name"), expression.Name("Address"), expression.Name("Items"))
//	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(proj).Build()
//	restaurantList := make([]Restaurant, 0)
//	if err != nil {
//		printError(err)
//		return restaurantList, err
//	}
//	params := &dynamodb.ScanInput{
//		TableName: aws.String(tableName),
//		ExpressionAttributeNames: expr.Names(),
//		ExpressionAttributeValues: expr.Values(),
//		FilterExpression: expr.Filter(),
//		ProjectionExpression: expr.Projection(),
//	}
//	svc := createSession()
//	res, err := svc.Scan(params)
//	for _, item := range res.Items {
//		restaurant := Restaurant{}
//		err := dynamodbattribute.UnmarshalMap(item, &restaurant)
//		if err != nil {
//			printError(err)
//			return restaurantList, err
//		}
//		restaurantList = append(restaurantList, restaurant)
//	}
//	return restaurantList, nil
//}
//
//func GetRestaurantSGI(tableName string, restaurantName string) ([]Restaurant, error) {
//
//	svc := createSession()
//
//	restaurantList := make([]Restaurant, 0)
//	ip := &dynamodb.QueryInput{
//		TableName: aws.String(tableName),
//		IndexName: aws.String("RName"),
//		KeyConditionExpression: aws.String("#n=:n"),
//		ExpressionAttributeNames: map[string]*string{
//			"#n": aws.String("Name"),
//		},
//		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
//			":n": {
//				S: aws.String(restaurantName),
//			},
//		},
//	}
//
//	res, err := svc.Query(ip)
//	if err != nil {
//		printError(err)
//		return restaurantList, err
//	}
//
//	for _, item := range res.Items {
//		restaurant := Restaurant{}
//		err := dynamodbattribute.UnmarshalMap(item, &restaurant)
//		if err != nil {
//			printError(err)
//			return restaurantList, err
//		}
//		restaurantList = append(restaurantList, restaurant)
//	}
//
//	return restaurantList, nil
//}
//
//func InsertRestaurant(tableName string, createRestaurant Restaurant) (Restaurant, error) {
//
//	svc := createSession()
//
//	uid, err := uuid.NewUUID()
//	if err != nil {
//		printError(err)
//		return Restaurant{}, err
//	}
//
//	createRestaurant.Id = int32(uid.ID())
//	items := createRestaurant.Items
//	createRestaurant.Items = []Item{}
//
//	ip, err := dynamodbattribute.MarshalMap(createRestaurant)
//	if err != nil {
//		printError(err)
//		return Restaurant{}, nil
//	}
//
//	input := &dynamodb.PutItemInput{
//		TableName: aws.String(tableName),
//		Item: ip,
//	}
//
//	_, err = svc.PutItem(input)
//	if err != nil {
//		printError(err)
//		return Restaurant{}, nil
//	}
//
//	createRestaurant.Items = insertItems(createRestaurant.Id, items, false)
//	op, err := UpdateRestaurant(tableName, createRestaurant, false)
//	if err != nil {
//		printError(err)
//		return Restaurant{}, err
//	}
//
//	return op, nil
//}
//
//func UpdateRestaurant(tableName string, updateRestaurant Restaurant, updateItems bool) (Restaurant, error) {
//
//	type KeyInput struct {
//		Id int32
//	}
//
//	key, err := dynamodbattribute.MarshalMap(KeyInput{Id: updateRestaurant.Id})
//	if err != nil {
//		fmt.Println("REACH1")
//		printError(err)
//		return Restaurant{}, err
//	}
//
//	if updateItems {
//		updateRestaurant.Items = insertItems(updateRestaurant.Id,
//			updateRestaurant.Items, false)
//	}
//
//	newRestaurant := removeRestId(updateRestaurant)
//	rmap, err := dynamodbattribute.MarshalMap(newRestaurant)
//	if err != nil {
//		printError(err)
//		fmt.Println("REACH2")
//		return Restaurant{}, err
//	}
//
//	input := &dynamodb.UpdateItemInput{
//		TableName: aws.String(tableName),
//		Key: key,
//		ExpressionAttributeValues: rmap,
//		ExpressionAttributeNames: map[string]*string{
//			"#n": aws.String("Name"),
//			"#itms": aws.String("Items"),
//		},
//		UpdateExpression: aws.String("set #n=:n, Address=:a, #itms=:itms"),
//	}
//
//	svc := createSession()
//
//	_, err = svc.UpdateItem(input)
//	if err != nil {
//		printError(err)
//		fmt.Println("REACH3")
//		return Restaurant{}, err
//	}
//
//	op := toNormalRestaurant(newRestaurant)
//	op.Id = updateRestaurant.Id
//	fmt.Println("REACH")
//	return op, nil
//}
//
//func GetAllItems(tableName string, filter *pb.ItemsFilter) ([]Item, error) {
//
//	items := make([]Item, 0)
//	restaurant, err := GetRestaurant(tableName, filter.RestaurantId)
//	if err != nil {
//		printError(err)
//		return items, err
//	}
//
//	if filter.Min == 0 && filter.Max == 0 {
//		return restaurant.Items, nil
//	}
//
//	for _, item := range restaurant.Items {
//		if item.Amount >= filter.Min && item.Amount <= filter.Max {
//			items = append(items, item)
//		}
//	}
//
//	return items, nil
//}
//
//func DeleteRestaurant(tableName string, restaurantId int32) error {
//
//	type KeyInput struct {
//		Id int32
//	}
//
//	type DeleteKey struct {
//		RestaurantId, ItemId int32
//	}
//
//	svc := createSession()
//
//	restaurant, err := GetRestaurant(tableName, restaurantId)
//	items := restaurant.Items
//	if err != nil {
//		printError(err)
//		return err
//	}
//
//	req := make([]*dynamodb.WriteRequest, 0)
//	for _, item := range items {
//		tmpKey, err := dynamodbattribute.MarshalMap(DeleteKey{
//			RestaurantId: restaurantId, ItemId: item.Id})
//		if err != nil {
//			printError(err)
//			return err
//		}
//		tmp := &dynamodb.WriteRequest{
//			DeleteRequest: &dynamodb.DeleteRequest{
//				Key: tmpKey,
//			},
//		}
//		req = append(req, tmp)
//	}
//
//	ip := &dynamodb.BatchWriteItemInput{
//		RequestItems: map[string][]*dynamodb.WriteRequest{
//			"Items": req,
//		},
//	}
//	_, err = svc.BatchWriteItem(ip)
//	if err != nil {
//		printError(err)
//		return err
//	}
//
//	key, err := dynamodbattribute.MarshalMap(KeyInput{Id: restaurantId})
//	if err != nil {
//		printError(err)
//		return err
//	}
//
//	input := &dynamodb.DeleteItemInput{
//		TableName: aws.String(tableName),
//		Key: key,
//	}
//
//	_, err = svc.DeleteItem(input)
//	if err != nil {
//		printError(err)
//		return err
//	}
//	return nil
//}