package mongo

import (
	"city-card-api/internal/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepo) CreateCard(userID primitive.ObjectID) (models.Card, error) {
	ctx := context.TODO()
	card := models.Card{
		ID:      primitive.NewObjectID(),
		UserID:  userID,
		Balance: 0,
	}
	_, err := r.collection.InsertOne(ctx, card)
	return card, err
}

func (r *mongoRepo) GetBalance(userID string) (float64, error) {
	var card models.Card
	ctx := context.TODO()
	objId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println("error get obj id from hex.", err)
		return 0, err
	}
	filter := bson.D{{
		"userID", objId,
	}}
	err = r.collection.FindOne(ctx, filter).Decode(&card)
	if err != nil {
		log.Println("error get balance. userID:", userID)
	}
	return card.Balance, err
}

func (r *mongoRepo) RequestPay(userID string) (string, error) {
	return "", nil
}

func (r *mongoRepo) SubmitPay(payToken string, amount float64) error {
	return nil
}

func (r *mongoRepo) AddMoney(userID string, amount float64) (float64, error) {
	var card models.Card
	ctx := context.TODO()
	objId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	filterFind := bson.D{{
		"userID", objId,
	}}
	filterUpdate := bson.D{{
		"$inc", bson.D{{"balance", amount}},
	}}
	result := float64(0)
	err = r.collection.FindOneAndUpdate(ctx, filterFind, filterUpdate).Decode(&card)
	if err == nil {
		result = card.Balance + amount
	}
	return result, err
}
