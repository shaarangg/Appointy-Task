package controllers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shaarangg/Appointy-Task/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertUser(database *mongo.Database) {
	userCollection := database.Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	user1 := models.User{"shaarang", "shaarngg.singh@gmail.com", "password"}
	insertResult, err := userCollection.InsertOne(ctx, user1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResult.InsertedID)
}
