package services

import (
	"city-card-api/internal/models"
	"city-card-api/internal/repository"
)

type profileService struct {
	db    repository.Profiler
	cache repository.Profiler
}

type authService struct {
	db    repository.UserRepository
	cache repository.UserRepository
}

// type services struct {
// 	*profileService
// 	*authService
// }

func NewProfileService(profileRepoDB repository.Profiler, profileRepoCache repository.Profiler) *profileService {
	return &profileService{
		db:    profileRepoDB,
		cache: profileRepoCache,
	}
}

// func NewServices(profileService *profileService, authService *authService) *services {
// 	return &services{
// 		profileService,
// 		authService,
// 	}
// }

type ProfileService interface {
	GetProfileByID(id string) (models.UserProfile, error)
}

type AuthService interface {
	Login(login, pass string) (models.UserProfile, models.Tokens, error)
	Register(user models.UserWithPassword) (models.UserProfile, models.Tokens, error)
	Decode(token string) (*models.UserJWT, error)
}

// type Services interface {
// 	ProfileService
// 	AuthService
// }
