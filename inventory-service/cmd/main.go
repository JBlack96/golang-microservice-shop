package main

import (
	"fmt"
	"github.com/JBlack96/golang-microservice-shop/inventory-service/internal/aws"
	localdynamo "github.com/JBlack96/golang-microservice-shop/inventory-service/internal/aws/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	loadConfig()

	sess, err := aws.CreateSession(viper.GetString("aws.region"), viper.GetString("aws.endpoint"))
	if err != nil {
		log.Fatal("failed to create aws session")
	}

	if viper.GetString("inventory_db.table_name") == "" {
		log.Fatal("failed to get required config value")
	}

	db := dynamodb.New(sess)
	ims := localdynamo.New(db, viper.GetString("inventory_db.table_name"))

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

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
