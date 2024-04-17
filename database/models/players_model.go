package models

// Player represents a player in the database

type Player struct {
	UserID   string
	UserNick string
	JPKey    string // Jackpot accumulation group ID
}
