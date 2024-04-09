package database

import (
	"database/sql"
	"errors"
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
