package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Update - Replaces an existing resource in the repository
func (h DynamodbDatabase) Update(resource interface{}) error {
	row, err := dynamodbattribute.MarshalMap(resource)

	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      row,
		TableName: aws.String(h.Table),
	}

	_, err = h.Connection.PutItem(input)

	if err != nil {
		return err
	}
	return nil
}