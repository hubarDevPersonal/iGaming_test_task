package main

import (
	"fmt"
	"iGaming/config"
	"iGaming/database"
	"log"
	"net/http"
)

func main() {
	cfg := config.New()

	db := database.ConnectToDB(cfg.DB)

	if db == nil {
		log.Fatal("Cannot connect to DB")
	}

	app := NewApp(db, cfg)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.RESTPort),
		Handler: app.Routes.Routes(),
	}

	//start server
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
