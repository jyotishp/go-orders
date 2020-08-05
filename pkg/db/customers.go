package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/uuid"
	pb "github.com/jyotishp/go-orders/pkg/proto"
	"log"
)

// Create a new customer entity and add it to the database
func AddCustomer(svc dynamodbiface.DynamoDBAPI, c *pb.CreateCustomer) (*pb.Customer, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed create new uuid: %v", err)
	}
	customer := &pb.Customer{
		CustomerId:      id.String(),
		Name:    c.Name,
		Address: c.Address,
	}
	log.Printf("%v", customer)
	err = AddObject(svc, CustomersTable, customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

// Fetch customer data from database by customer ID
func GetCustomer(svc dynamodbiface.DynamoDBAPI, i *pb.CustomerId) (*pb.Customer, error) {
	res, err := GetObjectById(svc, CustomersTable, *customerId.AttributeName, i.CustomerId)
	if err != nil {
		return nil, err
	}
	log.Printf("%v", res)
	return dbToCustomer(res)
}

// Convert the DB outputs to customer struct
func dbToCustomer(output map[string]*dynamodb.AttributeValue) (*pb.Customer, error) {
	customer := &pb.Customer{}
	err := dynamodbattribute.UnmarshalMap(output, &customer)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal object: %v", err)
	}
	return customer, nil
}

// Update the customer data in the database by customer ID
func UpdateCustomer(svc dynamodbiface.DynamoDBAPI, i *pb.UpdateCustomer) (*pb.Customer, error) {
	values := map[string]*dynamodb.AttributeValue{
		":n": { S: aws.String(i.Customer.Name) },
		":a": {
			M: map[string]*dynamodb.AttributeValue{
				"Line1": { S: aws.String(i.Customer.Address.Line1) },
				"Line2": { S: aws.String(i.Customer.Address.Line2) },
				"City": { S: aws.String(i.Customer.Address.City) },
				"State": { S: aws.String(i.Customer.Address.State) },
			},
		},
	}
	names := map[string]*string{
		"#n": aws.String("Name"),
		"#a": aws.String("Address"),
	}
	expr := "set #n=:n, #a=:a"
	err := UpdateObjectById(
		svc,
		CustomersTable,
		*customerId.AttributeName,
		i.CustomerId,
		values,
		names,
		expr,
		)
	if err != nil {
		return nil, err
	}
	return GetCustomer(svc, &pb.CustomerId{CustomerId: i.CustomerId})
}

// List all the customers in the DB
func ListCustomers(svc dynamodbiface.DynamoDBAPI) (*pb.CustomerList, error) {
	out, err := svc.Scan(&dynamodb.ScanInput{TableName: aws.String(CustomersTable)})
	customerList := make([]*pb.Customer, *out.Count)
	for idx, customer := range out.Items {
		err = dynamodbattribute.UnmarshalMap(customer, &customerList[idx])
		if err != nil {
			return nil, err
		}
	}
	return &pb.CustomerList{Customer: customerList}, nil
}

// Delete a customer from database by customer ID
func DeleteCustomer(svc dynamodbiface.DynamoDBAPI, i *pb.CustomerId) (*pb.Empty, error) {
	err := DeleteObjectById(svc, CustomersTable, *customerId.AttributeName, i.CustomerId)
	return &pb.Empty{}, err
}