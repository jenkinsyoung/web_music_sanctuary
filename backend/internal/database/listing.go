package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/models"
)

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

func (c *DBConnection) CreateGuitar(guitar models.Guitar) (int64, error) {
	var guitarID int64
	query := fmt.Sprintf(`INSERT INTO "guitar" (form, pickup_config, category) VALUES ($1, $2, $3) RETURNING id`)

	err := c.db.QueryRow(query, guitar.Form, guitar.PickupConfig, guitar.Category).Scan(&guitarID)
	return guitarID, err
}

func (c *DBConnection) CreateListing(listing models.Listing, guitarID, userID int64) (int64, error) {
	var listingID int64
	err := c.db.QueryRow(`INSERT INTO "listing" (user_id, guitar_id, name, cost, description) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		userID, guitarID, listing.GuitarName, listing.Cost, listing.Description).Scan(&listingID)

	return listingID, err
}

func (c *DBConnection) ImageListingCompound(listingID, imgID int64) {
	fmt.Println(imgID)
	c.db.QueryRow(`INSERT INTO "listing_pictures" (listing_id, picture_id) VALUES ($1, $2)`, listingID, imgID)
}
