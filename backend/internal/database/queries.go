package database

import (
	"database/sql"
	"fmt"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/models"
	"log"
)

//const Connection, err := ConnectToDB()

func (c *DBConnection) GetUserID(email, hashedPassword string) (int, error) {
	var userID int
	//query := fmt.Sprintf("")
	err := c.db.QueryRow(`SELECT id FROM "user" WHERE email=$1 AND password=$2`, email, hashedPassword).Scan(&userID)
	if err == sql.ErrNoRows {
		log.Printf("error no such user in database")
	}
	fmt.Println("from get user", hashedPassword)
	return userID, err
}

func (c *DBConnection) CreateUser(user *models.User) (int, error) {
	query := fmt.Sprintf(`INSERT INTO "user" (email, password, phone) VALUES ($1, $2, $3) RETURNING id`)
	var userID int
	err := c.db.QueryRow(query, user.Email, user.Password, user.Phone).Scan(&userID)
	return userID, err

}
