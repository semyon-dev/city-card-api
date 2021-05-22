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
	db       repository.UserRepository
	cache    repository.UserRepository
	dbPay    repository.CardRepository
	cachePay repository.CardRepository
}

type payService struct {
	cardDB    repository.CardRepository
	cardCache repository.CardRepository
	// payDB     repository.PayRepository
	// payCache  repository.PayRepository
}

func NewProfileService(profileRepoDB repository.Profiler, profileRepoCache repository.Profiler) *profileService {
	return &profileService{
		db:    profileRepoDB,
		cache: profileRepoCache,
	}
}

func NewPayService(cardDB repository.CardRepository, cardCache repository.CardRepository) *payService {
	return &payService{
		cardDB:    cardDB,
		cardCache: cardCache,
		// payDB:     payDB,
		// payCache:  payCahce,
	}
}

func NewAuthService(db repository.UserRepository, cache repository.UserRepository, dbPay repository.CardRepository, cachePay repository.CardRepository) *authService {
	return &authService{
		db:       db,
		cache:    cache,
		dbPay:    dbPay,
		cachePay: cachePay,
	}
}

type ProfileService interface {
	GetProfileByID(id string) (models.UserProfile, error)
}

type AuthService interface {
	Login(login, pass string) (models.UserProfile, models.Tokens, error)
	Register(user models.UserWithPassword) (models.UserProfile, models.Tokens, error)
	Decode(token string) (*models.UserJWT, error)
}

type PayService interface {
	Balance(userID string) (float64, error)
	AddMoney(userID string, amount float64) (float64, error)
	RequestPay(userID string) (string, error)
	SubmitPay(toUserID, payToken string, amount float64) error
}
