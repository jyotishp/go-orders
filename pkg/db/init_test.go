package db_test

import (
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type mockDynamoDBClient struct {
    dynamodbiface.DynamoDBAPI
}

func (m *mockDynamoDBClient) CreateTable(input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
    return &dynamodb.CreateTableOutput{TableDescription: &dynamodb.TableDescription{TableName: input.TableName}}, nil
}

func (m *mockDynamoDBClient) DescribeTable(input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
    return &dynamodb.DescribeTableOutput{Table: &dynamodb.TableDescription{
        TableName: input.TableName,
    }}, nil
}