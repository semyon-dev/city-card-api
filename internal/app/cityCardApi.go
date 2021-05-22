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
	mongoDB := config.Mongo.Database("city-card")
	// Repository
	profileDBRepo := mongo.NewMongoRepository(mongoDB, "profiles")
	authDBRepo := mongo.NewMongoRepository(mongoDB, "profiles")
	payDBRepo := mongo.NewMongoRepository(mongoDB, "cards")
	// Services
	// TODO: fix cache repository
	profileServices := services.NewProfileService(profileDBRepo, profileDBRepo)
	payServices := services.NewPayService(payDBRepo, payDBRepo)
	authServices := services.NewAuthService(authDBRepo, authDBRepo, payDBRepo, payDBRepo)

	// myServices := services.NewServices(profileServices, authServices)
	// Delivery
	// HTTP
	server := http.NewHttpServer(profileServices, authServices, payServices)
	httpEngine := server.StartHTTP()
	httpPort := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	err := httpEngine.Run(httpPort)
	if err != nil {
		log.Fatal(err)
	}
}
