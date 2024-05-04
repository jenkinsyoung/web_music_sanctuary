package database

import (
	"database/sql"
	"encoding/base64"
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

func (c *DBConnection) GetListingByID(id int64) (models.ListingFullInfo, error) {
	var listing models.ListingFullInfo
	if err := c.db.QueryRow(`SELECT * FROM "listing" WHERE id=$1`, id).Scan(&listing.Id, &listing.UserId,
		&listing.GuitarId, &listing.GuitarName, &listing.Cost, &listing.Description); err != nil {
		if err == sql.ErrNoRows {
			return listing, errors.New("advertisement with this id does not exist")
		}
		return listing, errors.New("error getting advertisement using id")
	}

	guitarInfo, err := c.GetGuitarInfo(listing.GuitarId)
	if err != nil {
		return listing, err
	}

	listing.Form = guitarInfo.Form
	listing.PickupConfig = guitarInfo.PickupConfig
	listing.Category = guitarInfo.Category

	listing.ImgList, err = c.GetListingImages(listing.Id)
	if err != nil {
		return listing, err
	}

	return listing, nil
}

func (c *DBConnection) GetListings() ([]models.ListingFullInfo, error) {
	var listings []models.ListingFullInfo

	rows, err := c.db.Query(`SELECT * FROM "listing"`)
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

func (c *DBConnection) GetGuitarInfo(guitarID int64) (models.Guitar, error) {
	var guitar models.Guitar

	if err := c.db.QueryRow(`SELECT * FROM "guitar" WHERE id=$1`, guitarID).Scan(&guitar.Id, &guitar.Form,
		&guitar.PickupConfig, &guitar.Category); err != nil {
		if err == sql.ErrNoRows {
			return guitar, errors.New("guitar with this id does not exist")
		}
		return guitar, errors.New("error getting guitar using id")
	}

	return guitar, nil
}

func (c *DBConnection) GetListingImages(listingID int64) ([]models.ImgJSON, error) {
	var images []models.ImgJSON

	rows, err := c.db.Query(`SELECT picture_id FROM "listing_pictures" WHERE listing_id=$1`, listingID)
	if err == sql.ErrNoRows {
		return images, nil
	}

	if err != nil {
		return images, err
	}
	defer rows.Close()

	for rows.Next() {
		var image []byte
		var imgID int64
		if err := rows.Scan(&imgID); err != nil {
			return images, err
		}

		err := c.db.QueryRow(`SELECT image FROM "picture" WHERE id=$1`, imgID).Scan(&image)

		if err != nil {
			return images, errors.New("picture with this id does not exist")
		}

		images = append(images, models.ImgJSON{Image: base64.StdEncoding.EncodeToString(image)})

	}

	return images, nil
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

func (c *DBConnection) UpdateListing(listing models.ListingFullInfo, userID int64) error {
	err := c.db.QueryRow(`UPDATE "listing" SET
	                 name=$1, cost=$2, description=$3 WHERE id=$4 AND user_id=$5`,
		listing.GuitarName, listing.Cost, listing.Description, listing.Id, userID)
	if err != nil {
		return errors.New("could not update user with this id")
	}

	err = c.db.QueryRow(`UPDATE "guitar" SET form=$1, pickup_config=$2 WHERE id=$3`,
		listing.Form, listing.PickupConfig, listing.GuitarId)

	if err != nil {
		return errors.New("could not update guitar with this id")
	}

	return nil
}

//TODO: delete listing
