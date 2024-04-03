package api

type TokenResponse struct {
	UserID      int64  `json:"userID"`
	AccessToken string `json:"accessToken"`
}
