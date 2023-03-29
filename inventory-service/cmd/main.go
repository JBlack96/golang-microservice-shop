package main

import (
	"github.com/JBlack96/golang-microservice-shop/inventory-service/internal/aws"
	localdynamo "github.com/JBlack96/golang-microservice-shop/inventory-service/internal/aws/dynamodb"
	"github.com/JBlack96/golang-microservice-shop/inventory-service/internal/domain"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	sess, err := aws.CreateSession("eu-west-1", "http://localhost:3000") // todo -> add config values
	if err != nil {
		log.Fatal("failed to create aws session")
	}

	db := dynamodb.New(sess)
	ims := localdynamo.New(db)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"health": "100%",
		})
	})

	r.GET("/items", func(c *gin.Context) {
		items, err := ims.GetAllItems()
		if err != nil {
			//todo do something with this error
			c.JSON(http.StatusInternalServerError, items)
		}

		c.JSON(http.StatusOK, items)
	})

	r.GET("/items/:id", func(c *gin.Context) {
		var item domain.InventoryItem
		if err := c.ShouldBindUri(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid item id"})
			return
		}

		item, err := ims.GetSingleItem("item id goes here")
		if err != nil {
			//todo do something with this error
			c.JSON(http.StatusInternalServerError, item)
		}

		c.JSON(http.StatusOK, item)
	})

	r.POST("/items", func(c *gin.Context) {
		item, err := ims.AddSingleItem(domain.InventoryItem{})
		if err != nil {
			//todo do something with this error
			c.JSON(http.StatusInternalServerError, item)
		}

		c.JSON(http.StatusOK, item)
	})

	r.Run()
}
