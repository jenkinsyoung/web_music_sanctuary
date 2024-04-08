package api

type TokenResponse struct {
	UserID      int64  `json:"user_id"`
	AccessToken string `json:"access_token"`
}

type NewAdvertisementResponse struct {
	AdvertisementID int64 `json:"advertisement_id"`
}

// TODO: переписать под новые структуры
//type AllAdvertisements struct {
//	Advertisements []models.Advertisement `json:"advertisements"`
//}
