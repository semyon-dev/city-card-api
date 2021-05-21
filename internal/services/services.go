package services

import (
	"city-card-api/internal/models"
	"city-card-api/internal/repository"
)

type profileService struct {
	db    repository.Profiler
	cache repository.Profiler
}
type Services struct {
	profileService *profileService
}

func NewProfileService(profileRepoDB repository.Profiler, profileRepoCache repository.Profiler) *profileService {
	return &profileService{
		db:    profileRepoDB,
		cache: profileRepoCache,
	}
}

func NewServices(profileService *profileService) *Services {
	return &Services{
		profileService: profileService,
	}
}

type ProfileService interface {
	GetProfileByID(id, accessToken string) (models.UserProfile, error)
}

// type Services interface {
// 	ProfileService
// }
