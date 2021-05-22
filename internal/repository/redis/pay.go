package redis

import (
	"city-card-api/internal/models"
	"context"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *redisRepo) RequestPay(userID string) (string, error) {
	// Generate pair uuid->userID
	ctx := context.TODO()
	myUuid, _ := uuid.NewUUID()
	// TODO: fix ttl
	err := r.client.Set(ctx, myUuid.String(), userID, 0).Err()
	return myUuid.String(), err
}

func (r *redisRepo) GetUserIDByPayToken(payToken string) (string, error) {
	ctx := context.TODO()
	val, err := r.client.Get(ctx, payToken).Result()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return val, nil
}

func (r *redisRepo) CreateCard(userID primitive.ObjectID) (models.Card, error) {
	return models.Card{}, nil
}

func (r *redisRepo) GetBalance(userID string) (float64, error) {
	return 0, nil
}

func (r *redisRepo) AddMoney(userID string, amount float64) (float64, error) {
	return 0, nil
}

func (r *redisRepo) SubmitPay(toUserID, fromUserID string, amount float64) error {
	return nil
}
