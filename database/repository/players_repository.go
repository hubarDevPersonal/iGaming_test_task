package repository

import (
	"database/sql"
	"iGaming/database/models"
)

type PlayersRepository interface {
	GetPlayerData(userID string) (*models.Player, error)
	CreatePlayer(player *models.Player) error
}

type playersRepository struct {
	db *sql.DB
}

func NewPlayersRepository(db *sql.DB) PlayersRepository {
	return &playersRepository{db: db}
}

func (p *playersRepository) GetPlayerData(userID string) (*models.Player, error) {
	query := `SELECT user_id, user_nick, jp_key FROM players WHERE user_id=$1`
	player := &models.Player{}

	err := p.db.QueryRow(query, userID).Scan(&player.UserID, &player.UserNick, &player.JPKey)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (p *playersRepository) CreatePlayer(player *models.Player) error {
	query := `INSERT INTO players (user_id, user_nick, jp_key) VALUES ($1, $2, $3)`
	_, err := p.db.Exec(query, player.UserID, player.UserNick, player.JPKey)
	if err != nil {
		return err
	}
	return nil
}
