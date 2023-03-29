package main

import (
	"github.com/JBlack96/golang-microservice-shop/inventory-service/internal/domain"
	"github.com/gin-gonic/gin"
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

	item, err := h.ims.GetSingleItem("item id goes here")
	if err != nil {
		//todo do something with this error
		c.JSON(http.StatusInternalServerError, item)
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetItems(c *gin.Context) {
	items, err := h.ims.GetAllItems()
	if err != nil {
		//todo do something with this error
		c.JSON(http.StatusInternalServerError, items)
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) AddItem(c *gin.Context) {
	item, err := h.ims.AddSingleItem(domain.InventoryItem{})
	if err != nil {
		//todo do something with this error
		c.JSON(http.StatusInternalServerError, item)
	}

	c.JSON(http.StatusOK, item)
}
