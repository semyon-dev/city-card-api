package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserProfile struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name"`
}
