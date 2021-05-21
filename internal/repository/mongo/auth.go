package mongo

import (
	"city-card-api/internal/models"
)

func (r *mongoRepo) CreateUser(user models.UserWithPassword) (models.UserProfile, error) {
	return models.UserProfile{}, nil
}

func (r *mongoRepo) ReadByLoginAndPass(login, pass string) (models.UserProfile, error) {
	return models.UserProfile{}, nil
}

func (r *mongoRepo) ChangePassword(userID int, currentPass, newPass string) error {
	return nil
}

func (r *mongoRepo) UpdateProfile(name, surname, mname string) error {
	return nil
}

func (r *mongoRepo) DeleteUser(userID int) error {
	return nil
}
