package repository

import (
	"city-card-api/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

type CardRepository interface {
	CreateCard(userID primitive.ObjectID) (models.Card, error)
	GetBalance(userID string) (float64, error)
	AddMoney(userID string, amount float64) (float64, error)
	RequestPay(userID string) (string, error)
	GetUserIDByPayToken(payToken string) (string, error)
	SubmitPay(toUserID, fromUserID string, amount float64) error
}
