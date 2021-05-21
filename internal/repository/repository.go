package repository

import (
	"city-card-api/internal/models"
)

type Profiler interface {
	GetProfileByID(id string) (models.UserProfile, error)
}

type UserRepository interface {
	CreateUser(user models.UserWithPassword) (models.UserProfile, error)
	ReadByLoginAndPass(login, pass string) (models.UserProfile, error)
	ChangePassword(userID int, currentPass, newPass string) error
	UpdateProfile(name, surname, mname string) error
	DeleteUser(userID int) error
}
