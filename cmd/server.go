package main

import (
	"database/sql"
	"iGaming/internal/api"
)

type App struct {
	Routes *api.Routes
}

func NewApp(db *sql.DB) *App {
	return &App{
		Routes: api.NewRoutes(db),
	}
}
