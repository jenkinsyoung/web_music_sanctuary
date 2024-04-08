package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/models"
	"log"
)

// const Connection, err := ConnectToDB()

func (c *DBConnection) CreateUser(user *models.User) (int, error) {
	if c.GetUserInfo(user.Email).Id != 0 {
		return 0, fmt.Errorf("user alredy exists")
	}
	query := fmt.Sprintf(`INSERT INTO "user" (name, surname, patronymic, email, password, phone) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`)
	var userID int
	err := c.db.QueryRow(query, user.Name, user.Surname, user.Patronymic, user.Email, user.Password, user.Phone).Scan(&userID)
	return userID, err
}

func (c *DBConnection) GetUserInfo(email string) models.User {
	var user models.User
	err := c.db.QueryRow(`SELECT * FROM "user" WHERE email=$1`, email).Scan(&user.Id, &user.Name, &user.Surname,
		&user.Patronymic, &user.Email, &user.Password, &user.Phone)
	if err != nil {
		log.Printf("Error getting user data using email %s", err)
	}
	return user
}

func (c *DBConnection) ImgInsert(img []byte) int64 {
	var id int64
	err := c.db.QueryRow(`INSERT INTO "picture" (image) VALUES ($1) RETURNING id`, img).Scan(&id)
	if err != nil {
		log.Printf("Error inserting img to db %s", err)
	}
	return id
}

//func (c *DBConnection) NewListing(ad *models.Listing) (int64, error) {
//	var listingID int64
//	query := fmt.Sprintf(`INSERT INTO "listing" (user_id, guitar_id, name, cost, description) VALUES ($1, $2, $3, $4, $5) RETURNING id`)
//	err := c.db.QueryRow(query, ad.UserId, ad.Description, ad.Name, ad.Cost, ad.TypeId).Scan(&listingID)
//	if err != nil {
//		return 0, err
//	}
//	return adID, nil
//}

func (c *DBConnection) GetListingByID(id int64) (models.Listing, error) {
	var listing models.Listing
	if err := c.db.QueryRow(`SELECT * FROM "listing" WHERE id=$1`, id).Scan(&listing.Id, &listing.UserId,
		&listing.GuitarId, &listing.GuitarName, &listing.Cost, listing.Description); err != nil {
		if err == sql.ErrNoRows {
			return listing, errors.New("advertisement with this id does not exist")
		}
		return listing, errors.New("error getting advertisement using id")
	}

	return listing, nil
}

func (c *DBConnection) GetListings() ([]models.Listing, error) {
	var listings []models.Listing

	rows, err := c.db.Query(`SELECT * FROM "listing"`)
	if err != nil {
		return listings, err
	}
	defer rows.Close()

	for rows.Next() {
		var listing models.Listing
		if err := rows.Scan(&listing.Id, &listing.UserId,
			&listing.GuitarId, &listing.GuitarName, &listing.Cost, listing.Description); err != nil {
			return listings, err
		}
		listings = append(listings, listing)
	}

	return listings, nil
}

//func (c *DBConnection) GetGuitarTypeID(name string)(int64, error){
//	var typeID int64
//	err := c.db.QueryRow(`SELECT id FROM "type" WHERE name=$1`, name).Scan(&typeID)
//	if err != nil{
//		log.Printf("error could not find type of guitar %s", err)
//		return 0, err
//	}
//	return typeID, nil
//}
