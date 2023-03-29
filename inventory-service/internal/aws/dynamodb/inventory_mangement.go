package dynamodb

import (
	"github.com/JBlack96/golang-microservice-shop/inventory-service/internal/domain"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type InventoryManagementService struct {
	db dynamodbiface.DynamoDBAPI
}

func New(db dynamodbiface.DynamoDBAPI) *InventoryManagementService {
	return &InventoryManagementService{
		db: db,
	}
}

func (ims *InventoryManagementService) GetAllItems() ([]domain.InventoryItem, error) {
	return []domain.InventoryItem{}, nil
}

func (ims *InventoryManagementService) GetSingleItem(ID string) (domain.InventoryItem, error) {
	return domain.InventoryItem{}, nil
}

func (ims *InventoryManagementService) AddSingleItem(item domain.InventoryItem) (domain.InventoryItem, error) {
	return domain.InventoryItem{}, nil
}
