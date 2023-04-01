package dynamodb

import (
	"github.com/JBlack96/golang-microservice-shop/inventory-service/internal/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type InventoryManagementService struct {
	tableName string
	db        dynamodbiface.DynamoDBAPI
}

func New(db dynamodbiface.DynamoDBAPI, tableName string) *InventoryManagementService {
	return &InventoryManagementService{
		db:        db,
		tableName: tableName,
	}
}

func (ims *InventoryManagementService) GetAllItems() ([]domain.InventoryItem, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(ims.tableName),
	}

	output, err := ims.db.Scan(input)
	if err != nil {
		return []domain.InventoryItem{}, err
	}

	var items []domain.InventoryItem
	if err = dynamodbattribute.UnmarshalListOfMaps(output.Items, &items); err != nil {
		return []domain.InventoryItem{}, err
	}

	return items, nil
}

func (ims *InventoryManagementService) GetSingleItem(ID string) (domain.InventoryItem, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(ims.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(ID)},
		},
	}

	output, err := ims.db.GetItem(input)
	if err != nil {
		return domain.InventoryItem{}, err
	}

	var item domain.InventoryItem
	if err = dynamodbattribute.UnmarshalMap(output.Item, &item); err != nil {
		return domain.InventoryItem{}, err
	}

	return item, nil
}

func (ims *InventoryManagementService) AddSingleItem(item domain.InventoryItem) (domain.InventoryItem, error) {
	marshalledItem, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return domain.InventoryItem{}, err
	}

	input := &dynamodb.PutItemInput{
		TableName:           aws.String(ims.tableName),
		Item:                marshalledItem,
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	}

	if _, err := ims.db.PutItem(input); err != nil {
		return domain.InventoryItem{}, err
	}

	return item, nil
}
