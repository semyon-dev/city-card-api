package config

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once

func ConnectToMongoDB() {
	once.Do(connectToMongoDB)
}

var Mongo *mongo.Client

func connectToMongoDB() {
	var mongoURI string
	if mongoURI = os.Getenv("MONGO_URI"); mongoURI == "" {
		log.Fatal("Error setup mongo uri for connect. Env: MONGO_URI")
	}
	var err error
	Mongo, err = mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = Mongo.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// _, err = Mongo.Database("goods-scanner").Collection("lists").Indexes().CreateOne(
	// 	context.Background(),
	// 	mongo.IndexModel{
	// 		Keys:    bson.D{{Key: "email", Value: 1}, {Key: "telephone", Value: 1}, {Key: "login", Value: 1}},
	// 		Options: options.Index().SetUnique(true),
	// 	},
	// )
	// err = Mongo.Ping(ctx)
	if err != nil {
		log.Fatal("Create index error:", err)
	}
	// defer client.Disconnect(ctx)
}
