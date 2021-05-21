package repository

import "city-card-api/internal/models"

type Profiler interface {
	GetProfileByID(id string) (models.UserProfile, error)
}
