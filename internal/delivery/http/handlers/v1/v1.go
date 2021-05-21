package v1

import "city-card-api/internal/services"

type HttpV1 struct {
	profile services.ProfileService
	auth    services.AuthService
}

func NewHttpV1(profile services.ProfileService, auth services.AuthService) *HttpV1 {
	return &HttpV1{
		profile: profile,
		auth:    auth,
	}
}
