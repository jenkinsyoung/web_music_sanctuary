package database

import (
	"fmt"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/models"
	"log"
)

func (c *DBConnection) CreateUser(user *models.User) (int, error) {
	if c.GetUserInfoByEmail(user.Email).Id != 0 {
		return 0, fmt.Errorf("user alredy exists")
	}
	query := fmt.Sprintf(`INSERT INTO "user" (name, surname, patronymic, email, password, phone) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`)
	var userID int
	err := c.db.QueryRow(query, user.Name, user.Surname, user.Patronymic, user.Email, user.Password, user.Phone).Scan(&userID)
	return userID, err
}

func (c *DBConnection) GetUserInfoByEmail(email string) models.User {
	var user models.User
	err := c.db.QueryRow(`SELECT * FROM "user" WHERE email=$1`, email).Scan(&user.Id, &user.Name, &user.Surname,
		&user.Patronymic, &user.Email, &user.Password, &user.Phone)
	if err != nil {
		log.Printf("Error getting user data using email %s", err)
	}
	return user
}

func (c *DBConnection) GetUserInfoByID(userID int64) models.User {
	var user models.User
	err := c.db.QueryRow(`SELECT * FROM "user" WHERE id=$1`, userID).Scan(&user.Id, &user.Name, &user.Surname,
		&user.Patronymic, &user.Email, &user.Password, &user.Phone)
	if err != nil {
		log.Printf("Error getting user data using id %s", err)
	}
	return user
}

func (c *DBConnection) GetUserListings(userID int64) ([]models.Listing, error) {
	var res []models.Listing

	rows, err := c.db.Query(`SELECT * FROM "listing" WHERE user_id=$1`, userID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var listing models.Listing
		if err := rows.Scan(&listing.Id, &listing.UserId,
			&listing.GuitarId, &listing.GuitarName, &listing.Cost, listing.Description); err != nil {
			return res, err
		}
		res = append(res, listing)
	}

	return res, nil
}
