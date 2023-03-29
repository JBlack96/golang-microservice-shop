package domain

type InventoryItem struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	ImgSrc      string  `json:"imgSrc"`
	Description string  `json:"description"`
	// todo -> think about managing stock => (tradeoff of stock management within inventory service vs dedicated stock management service)
}

type InventoryManagementDBService interface {
	GetAllItems() ([]InventoryItem, error)
	GetSingleItem(ID string) (InventoryItem, error)
	AddSingleItem(item InventoryItem) (InventoryItem, error)
}
