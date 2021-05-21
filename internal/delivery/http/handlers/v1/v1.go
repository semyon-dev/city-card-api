package v1

import "city-card-api/internal/services"

type HttpV1 struct {
	services *services.Services
}

func NewHttpV1(s *services.Services) *HttpV1 {
	return &HttpV1{
		services: s,
	}
}
