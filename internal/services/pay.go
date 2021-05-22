package services

import (
	_ "city-card-api/internal/models"
	_ "city-card-api/internal/repository"
)

func (s *payService) Balance(userID string) (float64, error) {
	balance, err := s.db.GetBalance(userID)
	return balance, err
}

func (s *payService) RequestPay(userID string) (string, error) {
	return "", nil
}

func (s *payService) SubmitPay(payToken string, amount int) error {
	return nil
}

func (s *payService) AddMoney(userID string, amount float64) (float64, error) {
	money, err := s.db.AddMoney(userID, amount)
	return money, err
}
