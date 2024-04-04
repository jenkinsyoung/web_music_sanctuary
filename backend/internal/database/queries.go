package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/models"
	"log"
)

// const Connection, err := ConnectToDB()

// TODO: полная хуйня? или нет для получения толького Id?
func (c *DBConnection) GetUserID(email, hashedPassword string) (int, error) {
	var userID int
	err := c.db.QueryRow(`SELECT id FROM "user" WHERE email=$1 AND password=$2`, email, hashedPassword).Scan(&userID)
	if err == sql.ErrNoRows {
		log.Printf("error no such user in database")
	}
	return userID, err
}

func (c *DBConnection) CreateUser(user *models.User) (int, error) {
	if c.GetUserInfo(user.Email).Id != 0 {
		return 0, fmt.Errorf("user alredy exists")
	}
	query := fmt.Sprintf(`INSERT INTO "user" (email, password, phone) VALUES ($1, $2, $3) RETURNING id`)
	var userID int
	err := c.db.QueryRow(query, user.Email, user.Password, user.Phone).Scan(&userID)
	return userID, err

}

func (c *DBConnection) GetUserInfo(email string) models.User {
	var user models.User
	err := c.db.QueryRow(`SELECT * FROM "user" WHERE email=$1`, email).Scan(&user.Id, &user.Email, &user.Password, &user.Phone)
	if err != nil {
		log.Printf("Error get user data using email %s", err)
	}
	return user
}

func (c *DBConnection) NewAdvertisement(ad *models.Advertisement) (int64, error) {
	var adID int64
	query := fmt.Sprintf(`INSERT INTO "advertisement" (user_id, description, name, cost, type_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`)
	err := c.db.QueryRow(query, ad.UserId, ad.Description, ad.Name, ad.Cost, ad.TypeId).Scan(&adID)
	if err != nil {
		return 0, err
	}
	return adID, nil
}

func (c *DBConnection) NewMicrocategories(ad *models.Advertisement) {
	for _, mc := range ad.Microcategories {
		c.db.QueryRow(`INSERT INTO advertisement_category (advertisement_id, microcategory_id) VALUES ($1, $2)`, ad.Id, mc)
	}
}

func (c *DBConnection) GetAdvertisementByID(id int64) (models.Advertisement, error) {
	var adv models.Advertisement
	if err := c.db.QueryRow(`SELECT * FROM "advertisement" WHERE id=$1`, id).Scan(&adv.Id, &adv.UserId,
		adv.Description, adv.Name, adv.Cost, adv.TypeId); err != nil {
		if err == sql.ErrNoRows {
			return adv, errors.New("advertisement with this id does not exist")
		}
		return adv, errors.New("error getting advertisement using id")
	}

	adv.Microcategories = c.GetMicrocategoriesForAdvertisement(adv.Id)

	return adv, nil
}

func (c *DBConnection) GetAdvertisements() ([]models.Advertisement, error) {
	var advertisements []models.Advertisement

	rows, err := c.db.Query(`SELECT * FROM "advertisement"`)
	if err != nil {
		return advertisements, err
	}
	defer rows.Close()

	for rows.Next() {
		var adv models.Advertisement
		if err := rows.Scan(&adv.Id, &adv.UserId, &adv.Description, &adv.Name, &adv.Cost, &adv.TypeId); err != nil {
			return advertisements, err
		}
		adv.Microcategories = c.GetMicrocategoriesForAdvertisement(adv.Id)
		advertisements = append(advertisements, adv)
	}

	return advertisements, nil
}

func (c *DBConnection) GetMicrocategoriesForAdvertisement(adID int64) []int64 {
	var res []int64

	rows, err := c.db.Query(`SELECT microcategory_id FROM "advertisement_category" WHERE advertisement_id=$1`, adID)
	if err != nil {
		return res
	}

	defer rows.Close()

	for rows.Next() {
		var mcId int64
		if err := rows.Scan(&mcId); err != nil {
			return res
		}
		res = append(res, mcId)
	}

	return res
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
