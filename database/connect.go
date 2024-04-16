package database

import (
	"database/sql"
	"fmt"
	"iGaming/config"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var counts int64

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createDSN(dbConfig config.Database) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable connect_timeout=5", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)
}

func ConnectToDB(dbConfig config.Database) *sql.DB {
	dsn := createDSN(dbConfig)
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Waiting for DB to start...", err)
			counts++
		} else {
			log.Println("DB started")
			return connection
		}

		if counts > dbConfig.RestartCount {
			log.Fatal("DB never started")
			return nil
		}
		log.Println("Backing off for 2 seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}
