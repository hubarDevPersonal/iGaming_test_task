package services

import (
	"database/sql"
	"fmt"
	"iGaming/database/repository"
	"iGaming/internal/entities"
)

type BalanceService interface {
	GetBalance(requestType string, player *entities.Player) (*entities.Player, error)
}

type balanceService struct {
	PlayersRepo repository.PlayersRepository
	WalletRepo  repository.BalanceRepository
}

func NewBalanceService(db *sql.DB) BalanceService {
	return &balanceService{
		PlayersRepo: repository.NewPlayersRepository(db),
		WalletRepo:  repository.NewBalanceRepository(db),
	}
}

func (b *balanceService) GetBalance(requestType string, player *entities.Player) (*entities.Player, error) {
	switch requestType {
	case "balance":
		balance, err := b.WalletRepo.GetBalance(player.Balance.Currency, player.Balance.GameSessionID)
		if err != nil {
			return nil, err
		}
		player.UserID = balance.UserID
		playerData, err := b.PlayersRepo.GetPlayerData(player.UserID)
		if err != nil {
			return nil, err
		}
		player.ConvertPlayerToEntity(playerData)
		player.Balance.ConvertWalletToBalance(balance)
		return player, nil

	case "debit":
		err := b.WalletRepo.UpdatePlayerBalance(player.UserID, player.Balance.Balance, "debit")
		if err != nil {
			return player, err
		}
		return player, nil

	case "credit":
		err := b.WalletRepo.UpdatePlayerBalance(player.UserID, player.Balance.Balance, "credit")
		if err != nil {
			return player, err
		}
		return player, nil

	case "rollback":
		err := b.WalletRepo.UpdatePlayerBalance(player.UserID, player.Balance.Balance, "rollback")
		if err != nil {
			return player, err
		}
		return player, nil

	default:
		return nil, fmt.Errorf("invalid request type: %s", requestType)
	}
}
