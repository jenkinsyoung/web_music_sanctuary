package database

import (
	"database/sql"
	"fmt"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/config"
	"log"
)

type DBConnection struct {
	db *sql.DB
}

func ConnectToDB(dbStruct config.Database) (*DBConnection, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbStruct.Host, dbStruct.Port, dbStruct.UserName, dbStruct.Password, dbStruct.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("error connecting to database %s", err)
		return nil, err

	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("error closing database %s", err)
		}
	}(db)
	err = db.Ping()
	if err != nil {
		log.Printf("error pinging database %s", err)
		return nil, err
	}
	return &DBConnection{db: db}, nil
}
