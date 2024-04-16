package repository

import "database/sql"

type PlayersRepository interface {
	GetPlayerBalance(userID string) (int64, error)
}

type playersRepository struct {
	db *sql.DB
}

func NewPlayersRepository(db *sql.DB) PlayersRepository {
	return &playersRepository{db: db}
}

func (p playersRepository) GetPlayerBalance(userID string) (int64, error) {
	//TODO implement me
	return 0, nil
}
