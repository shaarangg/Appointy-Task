package controllers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shaarangg/Appointy-Task/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertPost(database *mongo.Database) {
	postCollection := database.Collection("post")
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	currDate := time.Now()
	post1 := models.Post{"shaarangg.singh@gmail.com", "Hello this is my first post", "aksdasd.com", currDate.String()}
	insertResult, err := postCollection.InsertOne(ctx, post1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResult.InsertedID)
}
