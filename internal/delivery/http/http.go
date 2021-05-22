package http

import (
	v1 "city-card-api/internal/delivery/http/handlers/v1"
	"city-card-api/internal/delivery/http/middlewares"
	"city-card-api/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type httpServer struct {
	v1          *v1.HttpV1
	middlewares *middlewares.Middlewares
}

func NewHttpServer(profile services.ProfileService, auth services.AuthService, pay services.PayService) *httpServer {
	return &httpServer{
		v1:          v1.NewHttpV1(profile, auth, pay),
		middlewares: middlewares.NewHttpMiddleware(auth),
	}
}

func (server *httpServer) StartHTTP() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	myCors := cors.DefaultConfig()
	myCors.AllowAllOrigins = true
	myCors.AddAllowHeaders("Authorization")
	router.Use(cors.New(myCors))
	// prometheus
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/auth/login", server.v1.Login)
	v1.POST("/auth/register", server.v1.Register)
	v1.POST("/auth/refresh", server.v1.Refresh)

	v1.Use(server.middlewares.Auth())
	{
		pay := v1.Group("/pay")
		pay.GET("/balance", server.v1.Balance)
		pay.POST("/money", server.v1.AddMoney)
		pay.POST("/request", server.v1.RequestPay)
		pay.POST("/approve", server.v1.ApprovePay)

	}
	return router
}
