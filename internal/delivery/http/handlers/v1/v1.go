package v1

import "city-card-api/internal/services"

type HttpV1 struct {
	profile services.ProfileService
	auth    services.AuthService
	pay     services.PayService
}

func NewHttpV1(profile services.ProfileService, auth services.AuthService, pay services.PayService) *HttpV1 {
	return &HttpV1{
		profile: profile,
		auth:    auth,
		pay:     pay,
	}
}
