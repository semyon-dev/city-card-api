package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepo struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database, collectionName string) *mongoRepo {
	return &mongoRepo{
		collection: db.Collection(collectionName),
	}
}
