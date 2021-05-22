package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Card struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	UserID  primitive.ObjectID `json:"userID" bson:"userID"`
	Balance float64            `json:"balance"`
}

type Transaction struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Token    string             `json:"token" bson:"token"`
	FromCard string             `json:"fromCard"`
	ToCard   string             `json:"toCard"`
	Amount   string             `json:"amount"`
}
