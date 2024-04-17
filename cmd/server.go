package main

import (
	"database/sql"
	"iGaming/config"
	"iGaming/internal/api"
)

type App struct {
	Routes *api.Routes
}

func NewApp(db *sql.DB, cfg *config.Repository) *App {
	return &App{
		Routes: api.NewRoutes(db, cfg),
	}
}
