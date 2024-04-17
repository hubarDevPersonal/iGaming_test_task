package entities

import "iGaming/database/models"

type Balance struct {
	UserID        string
	Currency      string
	Balance       int64
	MaxWin        int64
	Denomination  int64
	GameSessionID string
}

func (b *Balance) ConvertWalletToBalance(wallet *models.Balance) {
	b.UserID = wallet.UserID
	b.Currency = wallet.Currency
	b.Balance = wallet.Balance
	b.MaxWin = wallet.MaxWin
	b.Denomination = int64(wallet.Denomination) // Ensure type compatibility
	b.GameSessionID = wallet.GameSessionID
}
