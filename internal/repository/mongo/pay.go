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
func (r *mongoRepo) withDrawMoney(fromUserID string, amount float64) error {
	log.Println("fromUser", fromUserID, "amount", amount)
	// withdraw from user
	ctx := context.TODO()
	witdraw := -amount
	objId, err := primitive.ObjectIDFromHex(fromUserID)
	if err != nil {
		return err
	}
	filterFind := bson.D{{
		"$and", []bson.D{
			bson.D{{"userID", objId}},
			bson.D{{"balance", bson.D{{"$gte", amount}}}}},
	}}
	filterUpdate := bson.D{{
		"$inc", bson.D{{
			"balance", witdraw,
		}}}}
	_, err = r.collection.UpdateOne(ctx, filterFind, filterUpdate)
	return err
}
func (r *mongoRepo) incrementMoney(toUserID string, amount float64) error {
	// log.Println("fromUser", fromUserID, "amount", amount)
	// withdraw from user
	ctx := context.TODO()
	// witdraw := -amount
	objId, err := primitive.ObjectIDFromHex(toUserID)
	if err != nil {
		return err
	}
	filterFind := bson.D{{"userID", objId}}
	filterUpdate := bson.D{{
		"$inc", bson.D{{
			"balance", amount,
		}}}}
	var card models.Card
	_, err = r.collection.UpdateOne(ctx, filterFind, filterUpdate)
	log.Println(card)
	return err
}
func (r *mongoRepo) SubmitPay(toUserID, fromUserID string, amount float64) error {
	err := r.withDrawMoney(fromUserID, amount)
	if err != nil {
		return err
	}
	err = r.incrementMoney(toUserID, amount)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoRepo) GetUserIDByPayToken(payToken string) (string, error) {
	return "", nil
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
