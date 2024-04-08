package api

import (
	"encoding/json"
	"fmt"
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
	"sync"
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

func CreateAdvertisement(w http.ResponseWriter, r *http.Request) {
	advertisement := models.Listing{}
	if err := json.NewDecoder(r.Body).Decode(&advertisement); err != nil {
		log.Printf("error occured decode json to adv %v", err)
	}

	defer r.Body.Close()

	idFlow := make(chan int64, len(advertisement.ImgList))

	//var guitarInfo models.Guitar

	var wg sync.WaitGroup

	for id, x := range advertisement.ImgList {
		// асинхронная загрузка фото в бд и запись их Id в канал
		wg.Add(1)

		go func(img string) {
			defer wg.Done()
			imgMethods.ImgDecode(img, idFlow)
		}(x.Image)

		// Отправка фото для ML и парсинг в стурктуру с гитарой
		if id == 0 {
			wg.Add(1)
			go func() {
				//TODO: Написать ссылку для api к ML
				time.Sleep(time.Second)
				//url := "SOME URL TO ML API"
				//
				//postBody, err := json.Marshal(advertisement.ImgList[0])
				//if err != nil {
				//	log.Printf("error marshal img to ml %s", err)
				//}
				//responseBody := bytes.NewBuffer(postBody)
				//resp, err := http.Post(url, "application/json", responseBody)
				//if err != nil {
				//	log.Printf("error occured send resp to ml %v", err)
				//}
				//defer resp.Body.Close()
				//json.NewDecoder(resp.Body).Decode(&guitarInfo)
				defer wg.Done()
			}()
		}
	}

	wg.Wait()
	defer close(idFlow)

	var idPictureList []int64

	// TODO: тут можешь сделать автоматическое добавление данных (асинх в бдху, тк айди прилетает на канал)
	for range advertisement.ImgList {
		idPictureList = append(idPictureList, <-idFlow)
	}

	fmt.Println(idPictureList)
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

//func GetAllAdvertisements(w http.ResponseWriter, r *http.Request) {
//	//advertisements, err := database.DB.GetListings()
//	//if err != nil {
//	//	log.Printf("could not select advertisements from db %s", err)
//	}
//
//	//resp, err := json.Marshal(AllAdvertisements{Advertisements: advertisements})
//	//
//	//_, err = w.Write(resp)
//
//	w.WriteHeader(http.StatusOK)
//
//}

func SetupRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/sign-up", Register).Methods("POST")
	router.HandleFunc("/api/sign-in", Authorization).Methods("POST")
	router.HandleFunc("/api/create-ad", CreateAdvertisement).Methods("POST")

	router.HandleFunc("/api/user", User).Methods("GET")
	router.HandleFunc("/api/get-ad/{id}", GetAdvertisement).Methods("GET")
	//router.HandleFunc("/api/get-ads", GetAllAdvertisements).Methods("GET")

	return router
}
