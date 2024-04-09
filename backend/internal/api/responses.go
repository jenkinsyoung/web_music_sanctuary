package api

import "github.com/jenkinsyoung/web_music_sanctuary/internal/models"

type TokenResponse struct {
	UserID      int64  `json:"user_id"`
	AccessToken string `json:"access_token"`
}

type NewListingResponse struct {
	ListingID int64 `json:"listing_id"`
}

type AllListings struct {
	Listings []models.Listing `json:"listings"`
}
