package services

import "city-card-api/internal/models"

func (s *profileService) GetProfileByID(id string) (models.UserProfile, error) {
	profile, err := s.db.GetProfileByID(id)
	return profile, err
}
