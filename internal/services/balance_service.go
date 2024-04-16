package services

import "iGaming/database/repository"

type BalanceService interface {
	GetBalance(userID string) (int64, error)
}

type balanceService struct {
	PlayersRepo *repository.PlayersRepository
}

func NewBalanceService(playersRepo *repository.PlayersRepository) BalanceService {
	return &balanceService{PlayersRepo: playersRepo}
}

func (b balanceService) GetBalance(userID string) (int64, error) {
	//TODO implement me
	return 0, nil
}
