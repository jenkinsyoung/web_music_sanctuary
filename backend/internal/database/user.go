package database

import (
	"errors"
	"fmt"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/models"
	"log"
)

func (c *DBConnection) CreateUser(user *models.User) (int64, error) {
	_, err := c.GetUserInfoByEmail(user.Email)
	if err == nil {
		return 0, fmt.Errorf("user alredy exists")
	}

	query := fmt.Sprintf(`INSERT INTO "user" (name, surname, patronymic, email, password, phone) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`)
	var userID int64
	err = c.db.QueryRow(query, user.Name, user.Surname, user.Patronymic, user.Email, user.Password, user.Phone).Scan(&userID)
	return userID, err
}

func (c *DBConnection) GetUserInfoByEmail(email string) (models.User, error) {
	var user models.User
	err := c.db.QueryRow(`SELECT * FROM "user" WHERE email=$1`, email).Scan(&user.Id, &user.Name, &user.Surname,
		&user.Patronymic, &user.Email, &user.Password, &user.Phone)
	if err != nil {
		log.Printf("Error getting user data using email %s", err)
		return user, err
	}
	return user, nil
}

func (c *DBConnection) GetUser(email, hashedPassword string) (models.User, error) {
	var user models.User
	err := c.db.QueryRow(`SELECT * FROM "user" WHERE email=$1 AND password=$2`, email, hashedPassword).Scan(&user.Id, &user.Name, &user.Surname,
		&user.Patronymic, &user.Email, &user.Password, &user.Phone)
	if err != nil {
		log.Printf("Error getting user data using email and password %s", err)
		return user, err
	}
	return user, nil
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

func (c *DBConnection) GetUserListings(userID int64) ([]models.ListingFullInfo, error) {
	var listings []models.ListingFullInfo

	rows, err := c.db.Query(`SELECT * FROM "listing" WHERE user_id=$1`, userID)
	if err != nil {
		return listings, err
	}
	defer rows.Close()

	for rows.Next() {
		var listing models.ListingFullInfo
		if err := rows.Scan(&listing.Id, &listing.UserId,
			&listing.GuitarId, &listing.GuitarName, &listing.Cost, &listing.Description); err != nil {
			return listings, err
		}

		guitarInfo, err := c.GetGuitarInfo(listing.GuitarId)
		if err != nil {
			return nil, err
		}

		listing.Form = guitarInfo.Form
		listing.PickupConfig = guitarInfo.PickupConfig
		listing.Category = guitarInfo.Category

		listing.ImgList, err = c.GetListingImages(listing.Id)

		if err != nil {
			return listings, err
		}

		listings = append(listings, listing)
	}

	return listings, nil
}

//func (c *DBConnection) IsUsersListing(userID, listingID int64) error {
//
//}
//
//func (c *DBConnection) UpdateListing(userID int64, listing models.Listing) error {
//
//}

func (c *DBConnection) UpdateUserInfo(user models.User) error {
	err := c.db.QueryRow(`UPDATE "user" SET 
                  name=$1, surname=$2, patronymic=$3, email=$4, phone=$5 WHERE id=$6`,
		user.Name, user.Surname, user.Patronymic, user.Email, user.Phone, user.Id)
	if err != nil {
		return errors.New("could not update user with this id")
	}
	return nil
}
