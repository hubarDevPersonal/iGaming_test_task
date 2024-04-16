package api

import (
	"database/sql"
	"iGaming/internal/api/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Routes struct {
	OpenApiGames controllers.OpenApiGames
}

func (r *Routes) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewRoutes(db *sql.DB) *Routes {
	return &Routes{
		OpenApiGames: controllers.NewOpenApiGames(db),
	}
}

func (r *Routes) Routes() http.Handler {
	mux := chi.NewRouter()
	//cors
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/healthCheck"))

	mux.Get("/games-processor", r.OpenApiGames.GameProcessor)

	return mux
}
