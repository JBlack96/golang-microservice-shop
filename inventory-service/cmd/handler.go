package main

import (
	"fmt"
	"github.com/JBlack96/golang-microservice-shop/inventory-service/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"net/http"
)

type Handler struct {
	ims domain.InventoryManagementDBService
}

func NewHandler(ims domain.InventoryManagementDBService) *Handler {
	return &Handler{
		ims: ims,
	}
}

func (h *Handler) GetItem(c *gin.Context) {
	var item domain.InventoryItem
	if err := c.ShouldBindUri(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid item id"})
		return
	}

	item, err := h.ims.GetSingleItem(item.ID)
	if err != nil {
		//todo do something with this error
		c.JSON(http.StatusInternalServerError, item)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetItems(c *gin.Context) {
	items, err := h.ims.GetAllItems()
	if err != nil {
		//todo do something with this error
		c.JSON(http.StatusInternalServerError, items)
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) AddItem(c *gin.Context) {
	var newItem AddInventoryItemBody
	if err := c.ShouldBindBodyWith(&newItem, binding.JSON); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	item := domain.InventoryItem{
		ID:          uuid.New().String(),
		Name:        newItem.Name,
		Price:       newItem.Price,
		Description: newItem.Description,
		ImgSrc:      newItem.ImgSrc,
	}

	item, err := h.ims.AddSingleItem(item)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, item)
}

type AddInventoryItemBody struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	ImgSrc      string  `json:"imgSrc"`
	Description string  `json:"description"`
	// todo -> think about managing stock => (tradeoff of stock management within inventory service vs dedicated stock management service)
}
