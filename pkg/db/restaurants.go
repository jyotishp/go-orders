package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
	pb "github.com/jyotishp/go-orders/pkg/proto"
)

// GetRestaurant fetches the restuarant of given Id and returns it alongwith any error that it may encounter.
func GetRestaurant(tableName string, id int32) (Restaurant, error) {
	if !checkTable(tableName) {
		CreateRestaurantsTable(tableName)
	}

	type Input struct {
		Id int32
	}

	key, err := dynamodbattribute.MarshalMap(Input{Id: id})
	if err != nil {
		printError(err)
		return Restaurant{}, err
	}

	ip := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: key,
	}

	svc := createSession()

	res, err := svc.GetItem(ip)
	if err != nil {
		printError(err)
		return Restaurant{}, err
	}

	restaurant := Restaurant{}

	err = dynamodbattribute.UnmarshalMap(res.Item, &restaurant)
	if err != nil {
		printError(err)
		return Restaurant{}, err
	}
	if restaurant.Id == 0 {
		return Restaurant{}, fmt.Errorf("no restaurant found for given id")
	}
	return restaurant, nil
}

// GetRestaurantName scans the table and returns all the restuarants whose name matches the given name alongwith any error that it may encounter.
func GetRestaurantName(tableName string, restaurantName string) ([]Restaurant, error) {
	if !checkTable(tableName) {
		CreateRestaurantsTable(tableName)
	}

	filter := expression.Name("Name").Equal(expression.Value(restaurantName))
	proj := expression.NamesList(expression.Name("Id"),
		expression.Name("Name"), expression.Name("Address"), expression.Name("Items"))
	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(proj).Build()
	restaurantList := make([]Restaurant, 0)
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
		restaurant := Restaurant{}
		err := dynamodbattribute.UnmarshalMap(item, &restaurant)
		if err != nil {
			printError(err)
			return restaurantList, err
		}
		restaurantList = append(restaurantList, restaurant)
	}
	return restaurantList, nil
}

// GetRestaurantSGI fetches all the restuarants of given name and returns the list alongiwth any error that it may encounter.
// The querying is done based on the Secondary Global Index created called "RName".
func GetRestaurantSGI(tableName string, restaurantName string) ([]Restaurant, error) {
	if !checkTable(tableName) {
		CreateRestaurantsTable(tableName)
	}

	svc := createSession()

	restaurantList := make([]Restaurant, 0)
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

	res, err := svc.Query(ip)
	if err != nil {
		printError(err)
		return restaurantList, err
	}

	for _, item := range res.Items {
		restaurant := Restaurant{}
		err := dynamodbattribute.UnmarshalMap(item, &restaurant)
		if err != nil {
			printError(err)
			return restaurantList, err
		}
		restaurantList = append(restaurantList, restaurant)
	}

	return restaurantList, nil
}

// InsertRestaurant gets Restaurant Params (without ID) and assigns it a unique Id. It then inserts this restaurant into the db.
// It returns this along with any error that it may encounter.
func InsertRestaurant(tableName string, createRestaurant Restaurant) (Restaurant, error) {
	if !checkTable(tableName) {
		CreateRestaurantsTable(tableName)
	}

	svc := createSession()

	uid, err := uuid.NewUUID()
	if err != nil {
		printError(err)
		return Restaurant{}, err
	}

	createRestaurant.Id = int32(uid.ID())
	items := createRestaurant.Items
	createRestaurant.Items = []Item{}

	ip, err := dynamodbattribute.MarshalMap(createRestaurant)
	if err != nil {
		printError(err)
		return Restaurant{}, nil
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: ip,
	}

	_, err = svc.PutItem(input)
	if err != nil {
		printError(err)
		return Restaurant{}, nil
	}

	createRestaurant.Items = insertItems(createRestaurant.Id, items, false)
	op, err := UpdateRestaurant(tableName, createRestaurant, false)
	if err != nil {
		printError(err)
		return Restaurant{}, err
	}

	return op, nil
}

// UpdateRestaurant gets a Restaurant and updates it in the db. It also accepts a bool which if true will update the Items table as well.
// It returns the Restaurant and any error that it may encounter.
func UpdateRestaurant(tableName string, updateRestaurant Restaurant, updateItems bool) (Restaurant, error) {
	if !checkTable(tableName) {
		CreateRestaurantsTable(tableName)
	}

	type KeyInput struct {
		Id int32
	}

	key, err := dynamodbattribute.MarshalMap(KeyInput{Id: updateRestaurant.Id})
	if err != nil {
		fmt.Println("REACH1")
		printError(err)
		return Restaurant{}, err
	}

	if updateItems {
		updateRestaurant.Items = insertItems(updateRestaurant.Id,
			updateRestaurant.Items, false)
	}

	newRestaurant := removeRestId(updateRestaurant)
	rmap, err := dynamodbattribute.MarshalMap(newRestaurant)
	if err != nil {
		printError(err)
		fmt.Println("REACH2")
		return Restaurant{}, err
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: key,
		ExpressionAttributeValues: rmap,
		ExpressionAttributeNames: map[string]*string{
			"#n": aws.String("Name"),
			"#itms": aws.String("Items"),
		},
		UpdateExpression: aws.String("set #n=:n, Address=:a, #itms=:itms"),
	}

	svc := createSession()

	_, err = svc.UpdateItem(input)
	if err != nil {
		printError(err)
		fmt.Println("REACH3")
		return Restaurant{}, err
	}

	op := toNormalRestaurant(newRestaurant)
	op.Id = updateRestaurant.Id
	op.OrderCount = updateRestaurant.OrderCount
	fmt.Println("REACH")
	return op, nil
}

// GetAllItems accepts a restuarantId and a min and max value and returns
//all items of the restaurant whose amount falls into this value range alongwith any error that it may encounter.
func GetAllItems(tableName string, filter *pb.ItemsFilter) ([]Item, error) {
	if !checkTable(tableName) {
		CreateRestaurantsTable(tableName)
	}

	items := make([]Item, 0)
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

// DeleteRestaurant deletes the Restaurant from the db.
// It also deletes all items belonging to this restaurant from the Items table.
// BatchWrite alongwith DeleteRequest was used to do this efficiently.
func DeleteRestaurant(tableName string, restaurantId int32) error {

	type KeyInput struct {
		Id int32
	}

	type DeleteKey struct {
		RestaurantId, ItemId int32
	}

	svc := createSession()

	restaurant, err := GetRestaurant(tableName, restaurantId)
	items := restaurant.Items
	if err != nil {
		printError(err)
		return err
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

// updateRestaurantCount updates the atomic counter for the value OrderCount in the Restaurants table.
func updateRestaurantCount(tableName string, id int32) error {
	type KeyInput struct {
		Id int32
	}

	key, err := dynamodbattribute.MarshalMap(KeyInput{Id: id})
	if err != nil {
		printError(err)
	}

	ip := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: key,
		UpdateExpression: aws.String("set OrderCount=OrderCount+:inc"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":inc": {
				N: aws.String(fmt.Sprint("1")),
			},
		},
	}
	svc := createSession()
	_, err = svc.UpdateItem(ip)
	if err != nil {
		printError(err)
		return err
	}
	return nil
}