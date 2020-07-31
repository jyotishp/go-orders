package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func checkError(err error) {
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			case dynamodb.ErrCodeResourceInUseException:
				fmt.Println(dynamodb.ErrCodeResourceInUseException, aerr.Error())

			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}
}

func createSession() *dynamodb.DynamoDB {

	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region: aws.String("Mumbai"),
	}))

	return dynamodb.New(sess)
}

func checkTable(tableName string) bool  {

	svc := createSession()

	input := &dynamodb.ListTablesInput{}

	for {
		result, err := svc.ListTables(input)
		if err != nil {
			checkError(err)
			return false
		}

		for _, table := range result.TableNames {
			if *table == tableName {
				return true
			}
		}
		return false
	}
}