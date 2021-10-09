package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	Uid       primitive.ObjectID `bson:"uid, omitempty"`
	Caption   string             `bson:"caption, omitempty"`
	Image_url string             `bson:"imageurl, omitempty"`
	Timestamp string             `bson:"timestamp, omitempty"`
}
