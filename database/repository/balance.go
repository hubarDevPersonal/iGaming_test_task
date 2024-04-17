package repository

import (
	"database/sql"
	"fmt"
	"iGaming/database/models"
)

type BalanceRepository interface {
	GetBalance(currency, gameSessionId string) (*models.Balance, error)
	UpdatePlayerBalance(userID string, amount int64, transactionType string) error
	CreateWallet(wallet *models.Balance) error
}

type balanceRepository struct {
	db *sql.DB
}

func NewBalanceRepository(db *sql.DB) BalanceRepository {
	return &balanceRepository{db: db}
}

func (w *balanceRepository) GetBalance(currency, gameSessionId string) (*models.Balance, error) {
	if w.db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT user_id, currency, balance, max_win, denomination FROM balance WHERE currency=$1 AND game_session_id=$2`
	wallet := &models.Balance{}

	err := w.db.QueryRow(query, currency, gameSessionId).Scan(&wallet.UserID, &wallet.Currency, &wallet.Balance, &wallet.MaxWin, &wallet.Denomination)
	if err != nil {
		return nil, fmt.Errorf("error fetching player balance: %v", err)
	}

	return wallet, nil
}

func (w *balanceRepository) UpdatePlayerBalance(userID string, amount int64, transactionType string) error {
	_, err := w.db.Exec("CALL update_player_balance($1, $2, $3)", userID, amount, transactionType)
	if err != nil {
		return fmt.Errorf("error updating player balance: %v", err)
	}
	return nil
}

func (w *balanceRepository) CreateWallet(wallet *models.Balance) error {
	query := `INSERT INTO balance (user_id, currency, balance, max_win, denomination) VALUES ($1, $2, $3, $4, $5)`
	_, err := w.db.Exec(query, wallet.UserID, wallet.Currency, wallet.Balance, wallet.MaxWin, wallet.Denomination)
	if err != nil {
		return fmt.Errorf("error creating player wallet: %v", err)
	}
	return nil
}
