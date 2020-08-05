package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

// GetCustomer accepts an Id and returns the Customer of that Id from the db.
// If no such customer exists or it encounters an error, then it returns said error.
func GetCustomer(tableName string, id int32) (Customer, error) {

	if !checkTable(tableName) {
		createCustomersTable(tableName)
	}

	type Input struct {
		Id int32
	}

	key, err := dynamodbattribute.MarshalMap(Input{Id: id})
	if err != nil {
		printError(err)
		return Customer{}, err
	}

	ip := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: key,
	}

	svc := createSession()

	res, err := svc.GetItem(ip)
	if err != nil {
		printError(err)
		return Customer{}, err
	}

	customer := Customer{}

	err = dynamodbattribute.UnmarshalMap(res.Item, &customer)
	if err != nil {
		printError(err)
		return Customer{}, err
	}
	if customer == (Customer{}) {
		return Customer{}, fmt.Errorf("no customer found for given id")
	}
	return customer, nil
}

// InsertCustomer accepts a Customer Params, without an Id, and inserts this Customer into the db after assigning a unique ID.
// It returns the Customer (Id included) and any error that it may encounter.
func InsertCustomer(tableName string, createCustomer Customer) (Customer, error) {
	if !checkTable(tableName) {
		createCustomersTable(tableName)
	}

	uid, err := uuid.NewUUID()
	if err != nil {
		printError(err)
		return Customer{}, err
	}

	createCustomer.Id = int32(uid.ID())
	ip, err := dynamodbattribute.MarshalMap(createCustomer)
	if err != nil {
		printError(err)
		return Customer{}, nil
	}

	svc := createSession()
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: ip,
	}

	_, err = svc.PutItem(input)
	if err != nil {
		printError(err)
		return Customer{}, nil
	}

	return createCustomer, nil
}

// UpdateCustomer accepts a Customer and updates this Customer into the db.
// It returns the Customer and any error that it may encounter.
func UpdateCustomer(tableName string, updateCustomer Customer) (Customer, error) {
	if !checkTable(tableName) {
		createCustomersTable(tableName)
	}

	type KeyInput struct {
		Id int32
	}

	key, err := dynamodbattribute.MarshalMap(KeyInput{Id: updateCustomer.Id})
	if err != nil {
		printError(err)
		return Customer{}, err
	}
	customer := removeCustId(updateCustomer)
	if err != nil {
		printError(err)
		return Customer{}, err
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: key,
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {
				S: aws.String(customer.Name),
			},
			":a": {
				M: map[string]*dynamodb.AttributeValue{
					"Line1": {
						S: aws.String(customer.Address.Line1),
					},
					"Line2": {
						S: aws.String(customer.Address.Line2),
					},
					"City": {
						S: aws.String(customer.Address.City),
					},
					"State": {
						S: aws.String(customer.Address.State),
					},
				},
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#n": aws.String("Name"),
			"#a": aws.String("Address"),
		},
		UpdateExpression: aws.String("set #n=:n, #a=:a"),
	}

	svc := createSession()
	_, err = svc.UpdateItem(input)
	if err != nil {
		printError(err)
		return Customer{}, err
	}

	return updateCustomer, nil
}

// GetAllCustomers returns all the Customers of the table.
func GetAllCustomers(tableName string) ([]Customer, error) {
	if !checkTable(tableName) {
		createCustomersTable(tableName)
	}

	ip := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	op := make([]Customer, 0)
	svc := createSession()
	res, err := svc.Scan(ip)
	if err != nil {
		printError(err)
		return op, err
	}
	for _, item := range res.Items {
		customer := Customer{}
		err = dynamodbattribute.UnmarshalMap(item, &customer)
		if err != nil {
			printError(err)
			return op, err
		}
		op = append(op, customer)
	}
	return op, nil
}