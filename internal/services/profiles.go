package services

import "city-card-api/internal/models"

func (s *Services) GetProfileByID(id string) (models.UserProfile, error) {
	profile, err := s.profileService.db.GetProfileByID(id)
	return profile, err
}
