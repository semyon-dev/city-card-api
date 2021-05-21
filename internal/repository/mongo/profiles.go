package mongo

import (
	"city-card-api/internal/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepo) GetProfileByID(id string) (models.UserProfile, error) {
	ctx := context.TODO()
	bsonId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{
		"_id", bsonId,
	}}
	var user models.UserProfile
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return models.UserProfile{}, err
	}
	return user, nil
}
