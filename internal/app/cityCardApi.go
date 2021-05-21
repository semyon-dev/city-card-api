package app

import (
	"city-card-api/internal/config"
	"city-card-api/internal/delivery/http"
	"city-card-api/internal/repository/mongo"
	"city-card-api/internal/services"
	"fmt"
	"log"
	"os"
)

func Run() {
	config.ConnectToMongoDB()
	mongoDB := config.Mongo.Database("goods-scanner")
	// Repository
	profileDBRepo := mongo.NewMongoRepository(mongoDB, "lists")
	// Services
	// TODO: fix cache repository
	profileServices := services.NewProfileService(profileDBRepo, profileDBRepo)
	myServices := services.NewServices(profileServices)
	// Delivery
	// HTTP
	server := http.NewHttpServer(myServices)
	httpEngine := server.StartHTTP()
	httpPort := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	err := httpEngine.Run(httpPort)
	if err != nil {
		log.Fatal(err)
	}
}
