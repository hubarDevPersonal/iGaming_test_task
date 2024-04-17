package dto

import "iGaming/internal/entities"

// api	string	Name of operation, “balance” in this case
// data	object
// data.gameSessionId	string	Game Session Id
// data.currency	string	Name of users currency
type RequestBalance struct {
	Api  string `json:"api"`
	Data Data   `json:"data,omitempty"`
}

type Data struct {
	GameSessionId string   `json:"game_session_id,omitempty"`
	Currency      string   `json:"currency,omitempty"`
	Denomination  int      `json:"denomination"`
	MaxWin        int      `json:"max_win,omitempty"`
	UserId        string   `json:"user_id,omitempty"`
	JpKey         string   `json:"jp_key,omitempty"`
	SpinMeta      SpinMeta `json:"spin_meta,omitempty"`
	BetMeta       BetMeta  `json:"bet_meta,omitempty"`
	Notes         Notes    `json:"notes,omitempty"`
}

type SpinMeta struct {
	Lines        int `json:"lines,omitempty"`
	BetPerLine   int `json:"bet_per_line,omitempty"`
	TotalBet     int `json:"total_bet,omitempty"`
	SymbolMatrix int `json:"symbol_matrix,omitempty"`
}

type BetMeta struct {
	Bets int `json:"bets"`
}

type Notes struct {
	Notes string `json:"notes"`
}

func ToPlayer(dto RequestBalance) *entities.Player {
	return &entities.Player{
		Balance: &entities.Balance{
			Currency:      dto.Data.Currency,
			GameSessionID: dto.Data.GameSessionId,
		},
	}
}
