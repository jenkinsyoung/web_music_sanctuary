package database

import (
	"database/sql"
	"fmt"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/config"
	_ "github.com/lib/pq"
	"log"
)

type DBConnection struct {
	db *sql.DB
}

var DB *DBConnection

func ConnectToDB(dbStruct config.Database) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbStruct.Host, dbStruct.Port, dbStruct.UserName, dbStruct.Password, dbStruct.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("error connecting to database %s", err)
		return err

	}
	err = db.Ping()
	if err != nil {
		log.Printf("error pinging database %s", err)
		return err
	}
	DB = &DBConnection{db: db}
	return nil
}
