package api

import "github.com/jenkinsyoung/web_music_sanctuary/internal/models"

type TokenResponse struct {
	UserID      int64  `json:"user_id"`
	AccessToken string `json:"access_token"`
}

type NewAdvertisementResponse struct {
	AdvertisementID int64 `json:"advertisement_id"`
}

type AllAdvertisements struct {
	Advertisements []models.Advertisement `json:"advertisements"`
}
