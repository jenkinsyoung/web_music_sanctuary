package api

import (
	"encoding/json"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/database"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/hash"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/jwt"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/models"
	"log"
	"net/http"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error parsing json: %s", err)
	}

	user.Password = hash.PasswordHash(user.Password)

	id, err := database.DB.CreateUser(&user)
	if err != nil {
		log.Printf("error creating user: %s", err)
		w.WriteHeader(http.StatusConflict)
		return
	}

	token, err := jwt.GenerateToken(user.Email, user.Password)
	if err != nil {
		log.Printf("error generating access token: %s", err)
	}

	resp, err := json.Marshal(TokenResponse{UserID: int64(id), AccessToken: token})
	if err != nil {
		log.Printf("error marshalling json: %s", err)
	}

	_, err = w.Write(resp)
	if err != nil {
		log.Printf("error sending response: %s", err)
	}

	w.WriteHeader(http.StatusOK)
}

func Authorization(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error parsing json: %s", err)
	}
	userInfo := database.DB.GetUserInfoByEmail(user.Email)
	if !hash.CheckPassword(user.Password, userInfo.Password) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	token, err := jwt.GenerateToken(user.Email, user.Password)
	if err != nil {
		log.Printf("error generating access token: %s", err)
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, cookie)

	resp, err := json.Marshal(TokenResponse{UserID: userInfo.Id, AccessToken: token})
	if err != nil {
		log.Printf("error marshalling json: %s", err)
	}

	_, err = w.Write(resp)
	if err != nil {
		log.Printf("error sending response: %s", err)
	}
	w.WriteHeader(http.StatusOK)
}
