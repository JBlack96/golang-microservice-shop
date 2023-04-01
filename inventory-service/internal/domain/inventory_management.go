package domain

import "github.com/google/uuid"

type InventoryItem struct {
	ID          string  `json:"id" uri:"id" binding:"required,uuid"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	ImgSrc      string  `json:"imgSrc"`
	Description string  `json:"description"`
	// todo -> think about managing stock => (tradeoff of stock management within inventory service vs dedicated stock management service)
}

func (i *InventoryItem) SetRandomID() {
	id, _ := uuid.NewUUID()
	i.ID = id.String()
}

type InventoryManagementDBService interface {
	GetAllItems() ([]InventoryItem, error)
	GetSingleItem(ID string) (InventoryItem, error)
	AddSingleItem(item InventoryItem) (InventoryItem, error)
}
