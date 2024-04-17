package models

// Balance represents a balance in the database

type Balance struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	Balance       int64  `json:"balance"`
	Denomination  int    `json:"denomination"`
	Currency      string `json:"currency"`
	MaxWin        int64  `json:"max_win"`
	GameSessionID string `json:"game_session_id"`
}
