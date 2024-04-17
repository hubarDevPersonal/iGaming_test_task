package api

import (
	"database/sql"
	"iGaming/config"
	"iGaming/internal/api/controllers"
	"iGaming/internal/api/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Routes struct {
	OpenApiGames controllers.OpenApiGames
}

func NewRoutes(db *sql.DB, cfg *config.Repository) *Routes {
	helpers.Secret = cfg.Secret
	return &Routes{
		OpenApiGames: controllers.NewOpenApiGames(db),
	}
}

func (r *Routes) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Sign"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/healthCheck"))

	mux.Route("/api/v1", func(c chi.Router) {
		c.Post("/games", r.OpenApiGames.GameProcessor)
	})

	return mux
}
