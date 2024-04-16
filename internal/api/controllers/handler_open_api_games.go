package controllers

import (
	"database/sql"
	"iGaming/internal/api/helpers"
	"iGaming/internal/services"
	"net/http"
)

type OpenApiGames interface {
	GameProcessor(w http.ResponseWriter, r *http.Request)
}

type openApiGames struct {
	balanceService services.BalanceService
	db             *sql.DB
}

func NewOpenApiGames(db *sql.DB) OpenApiGames {
	return &openApiGames{db: db}
}

func (o openApiGames) GameProcessor(w http.ResponseWriter, r *http.Request) {

	res := &helpers.JsonResponse{
		Error:   false,
		Message: "Hello from the service",
		Data:    nil,
	}
	_ = res.WriteJSON(w, http.StatusOK, nil)

}
