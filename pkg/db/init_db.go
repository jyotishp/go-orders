package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func printError(err error) {
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
		return
	}
}

func DisableSSL() bool {
	res := os.Getenv("DISABLE_SSL")
	if len(res) == 0 {
		return false
	}
	return true
}

func createSession() *dynamodb.DynamoDB {

	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String(os.Getenv("DB_ENDPOINT")),
		DisableSSL: aws.Bool(DisableSSL()),
	}))

	return dynamodb.New(sess)
}

func checkTable(tableName string) bool  {

	svc := createSession()

	input := &dynamodb.ListTablesInput{}

	for {
		result, err := svc.ListTables(input)
		if err != nil {
			printError(err)
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