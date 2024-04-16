package dto

// api	string	Name of operation, “balance” in this case
// data	object
// data.gameSessionId	string	Game Session Id
// data.currency	string	Name of users currency
type RequestBalance struct {
	Api  string `json:"api"`
	Data Data   `json:"data"`
}

type Data struct {
	GameSessionId string   `json:"gameSessionId"`
	Currency      string   `json:"currency"`
	Denomination  int      `json:"denomination"`
	MaxWin        int      `json:"maxWin"`
	UserId        string   `json:"userId"`
	JpKey         string   `json:"jpKey"`
	SpinMeta      SpinMeta `json:"spinMeta"`
	BetMeta       BetMeta  `json:"betMeta"`
	Notes         Notes    `json:"notes"`
}

type SpinMeta struct {
	Lines        int `json:"lines"`
	BetPerLine   int `json:"betPerLine"`
	TotalBet     int `json:"totalBet"`
	SymbolMatrix int `json:"symbolMatrix"`
}

type BetMeta struct {
	Bets int `json:"bets"`
}

//type Bet struct {
//
//}

type Notes struct {
	Notes string `json:"notes"`
}
