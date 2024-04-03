package api

type TokenResponse struct {
	UserID      int    `json:"userID"`
	AccessToken string `json:"accessToken"`
}
