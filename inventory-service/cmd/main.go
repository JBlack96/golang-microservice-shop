package main

import (
	"github.com/JBlack96/golang-microservice-shop/inventory-service/internal/aws"
	localdynamo "github.com/JBlack96/golang-microservice-shop/inventory-service/internal/aws/dynamodb"
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

	handler := NewHandler(ims)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"health": "100%",
		})
	})

	r.GET("/items", handler.GetItems)
	r.POST("/items", handler.AddItem)
	r.GET("/items/:id", handler.GetItem)

	r.Run()
}
