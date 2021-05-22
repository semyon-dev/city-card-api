package middlewares

import (
	"city-card-api/internal/services"
)

type Middlewares struct {
	services.AuthService
}

func NewHttpMiddleware(authService services.AuthService) *Middlewares {
	return &Middlewares{
		authService,
	}
}
