package api

import (
	"encoding/json"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/database"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/hash"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/jwt"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/models"
	"github.com/jenkinsyoung/web_music_sanctuary/pkg/imgMethods"
	"log"
	"net/http"
	"strconv"
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
	userInfo := database.DB.GetUserInfo(user.Email)
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

func User(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		log.Printf("error getting cookie %s", err)
	}

	token, err := jwt2.ParseWithClaims(cookie.String(), &jwt2.MapClaims{}, func(token *jwt2.Token) (interface{}, error) {
		return []byte(jwt.SigningKey), nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	resp, err := json.Marshal(token.Claims)
	if err != nil {
		log.Printf("error marshalling json: %s", err)
	}

	_, err = w.Write(resp)
	if err != nil {
		log.Printf("error sending response: %s", err)
	}
	w.WriteHeader(http.StatusOK)
}

//TODO: напистаь отправку объявлений по БОЛЬШОЙ категории
//TODO: api/upload-image написать отправку

func ReceiveImage(w http.ResponseWriter, r *http.Request) {
	photo := models.Picture{}
	json.NewDecoder(r.Body).Decode(&photo)

	defer r.Body.Close()
	//TODO: написать файловую структру для сохраниения картино и присовение им айдишников

	err := imgMethods.SaveImageBase64(photo.Image)
	if err != nil {
		log.Printf("Error decode base64 %s", err)
	}
}

func CreateAdvertisement(w http.ResponseWriter, r *http.Request) {
	advertisement := models.Listing{}
	json.NewDecoder(r.Body).Decode(&advertisement)

	defer r.Body.Close()

	//TODO: Здесь должно быть получение данных из МЛ и подгрузка МК и К и тд...

	adID, err := database.DB.NewAdvertisement(&advertisement)
	if err != nil {
		log.Printf("Error creating advertisement %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	advertisement.Id = adID

	database.DB.NewMicrocategories(&advertisement)

	resp, err := json.Marshal(NewAdvertisementResponse{advertisement.Id})
	_, err = w.Write(resp)

	w.WriteHeader(http.StatusOK)
}

func GetAdvertisement(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("not valid id")
	}

	adv, err := database.DB.GetListingByID(int64(id))

	if err != nil {
		log.Printf(err.Error())
	}

	resp, err := json.Marshal(adv)

	_, err = w.Write(resp)
	w.WriteHeader(http.StatusOK)
}

func GetAllAdvertisements(w http.ResponseWriter, r *http.Request) {
	advertisements, err := database.DB.GetListings()
	if err != nil {
		log.Printf("could not select advertisements from db %s", err)
	}

	resp, err := json.Marshal(AllAdvertisements{Advertisements: advertisements})

	_, err = w.Write(resp)

	w.WriteHeader(http.StatusOK)

}

func SetupRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/sign-up", Register).Methods("POST")
	router.HandleFunc("/api/sign-in", Authorization).Methods("POST")
	router.HandleFunc("/api/save-image", ReceiveImage).Methods("POST")
	router.HandleFunc("/api/create-ad", CreateAdvertisement).Methods("POST")

	router.HandleFunc("/api/user", User).Methods("GET")
	router.HandleFunc("/api/get-ad/{id}", GetAdvertisement).Methods("GET")
	router.HandleFunc("/api/get-ads", GetAllAdvertisements).Methods("GET")

	return router
}
