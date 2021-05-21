package http

import (
	v1 "city-card-api/internal/delivery/http/handlers/v1"
	"city-card-api/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type httpServer struct {
	v1 *v1.HttpV1
}

func NewHttpServer(services *services.Services) *httpServer {
	return &httpServer{
		v1: v1.NewHttpV1(services),
	}
}

func (server *httpServer) StartHTTP() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.Use(cors.Default())
	// prometheus
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	api := router.Group("/api")
	v1 := api.Group("/v1")

	lists := v1.Group("/lists")
	lists.GET("/", server.v1.Hello)

	return router
}
