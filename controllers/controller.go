package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/shaarangg/Appointy-Task/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserRequests(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		fmt.Println("GET req at user")
		response.Header().Add("content-type", "application/json")
		userID := strings.TrimPrefix(request.URL.Path, "/users/")
		fmt.Println(userID)
		id, _ := primitive.ObjectIDFromHex(userID)
		fmt.Println(id)
		var user models.User
		collection := DB.Database("appointy").Collection("user")
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := collection.FindOne(ctx, models.User{ID: id}).Decode(&user)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		}
		json.NewEncoder(response).Encode(user)
	case "POST":
		fmt.Println("POST req at user")
		response.Header().Add("content-type", "application/json")
		var user models.User
		json.NewDecoder(request.Body).Decode(&user)
		fmt.Println(user.Name, user.Email, user.Password)
		collection := DB.Database("appointy").Collection("user")
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		result, _ := collection.InsertOne(ctx, user)
		json.NewEncoder(response).Encode(result)
	default:
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{"message": "not found"}`))
	}
}
