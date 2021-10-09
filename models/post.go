package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id"`
	Caption   string             `bson:"caption`
	Image_url string             `bson:"image_url`
	Timestamp time.Time          `bson:"timestamp`
}
