package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

func GetOrder(tableName string, id int32) (Order, error) {
	type Input struct {
		Id int32
	}

	key, err := dynamodbattribute.MarshalMap(Input{Id: id})
	if err != nil {
		printError(err)
		return Order{}, err
	}

	ip := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: key,
	}

	svc := createSession()

	res, err := svc.GetItem(ip)
	if err != nil {
		printError(err)
		return Order{}, err
	}

	order := Order{}

	err = dynamodbattribute.UnmarshalMap(res.Item, &order)
	if err != nil {
		printError(err)
		return Order{}, err
	}

	return order, nil
}

func InsertOrder(tableName string, createOrder Order) (Order, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		printError(err)
		return Order{}, err
	}

	createOrder.Id = int32(uid.ID())

	ip, err := dynamodbattribute.MarshalMap(createOrder)
	if err != nil {
		printError(err)
		return Order{}, nil
	}

	svc := createSession()
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: ip,
	}

	_, err = svc.PutItem(input)
	if err != nil {
		printError(err)
		return Order{}, nil
	}
	return createOrder, nil
}

func UpdateOrder(tableName string, updateOrder OrderIp) (Order, error) {
	type KeyInput struct {
		Id int32
	}

	key, err := dynamodbattribute.MarshalMap(KeyInput{Id: updateOrder.Id})
	if err != nil {
		printError(err)
		return Order{}, err
	}

	orderNoKey := removeOrderId(updateOrder)
	omap, err := dynamodbattribute.MarshalMap(orderNoKey)
	if err != nil {
		printError(err)
		return Order{}, err
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: key,
		ExpressionAttributeValues: omap,
		ExpressionAttributeNames: map[string]*string{
			"#dur": aws.String("Duration"),
			"#tm": aws.String("Time"),
			"#itms": aws.String("Items"),
		},
		UpdateExpression: aws.String("set Discount=:d, Amount=:amt, PaymentMethod=:pm, " +
			"Rating=:r, #dur=:dn, Cuisine=:c, #tm=:t, Verified=:v, " +
			"Customer=:ctmr, Restaurant=:rt, #itms=:itms"),
	}

	svc := createSession()
	_, err = svc.UpdateItem(input)
	if err != nil {
		printError(err)
		return Order{}, err
	}

	op := toNormalOrder(updateOrder)

	return op, nil
}