package services

import (
	_ "city-card-api/internal/models"
	_ "city-card-api/internal/repository"
	"log"
)

func (s *payService) Balance(userID string) (float64, error) {
	balance, err := s.cardDB.GetBalance(userID)
	return balance, err
}

func (s *payService) RequestPay(userID string) (string, error) {
	token, err := s.cardCache.RequestPay(userID)
	if err != nil {
		log.Println(err)
	}
	// log.Println("token", token)
	return token, err
}

func (s *payService) SubmitPay(toUserID, payToken string, amount float64) error {
	fromUserID, err := s.cardCache.GetUserIDByPayToken(payToken)
	if err != nil {
		return err
	}
	err = s.cardDB.SubmitPay(toUserID, fromUserID, amount)
	return err
}

func (s *payService) AddMoney(userID string, amount float64) (float64, error) {
	money, err := s.cardDB.AddMoney(userID, amount)
	return money, err
}
