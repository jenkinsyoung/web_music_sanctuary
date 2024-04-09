package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/database"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/models"
	"github.com/jenkinsyoung/web_music_sanctuary/pkg/imgMethods"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func CreateListing(w http.ResponseWriter, r *http.Request) {
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

func GetListing(w http.ResponseWriter, r *http.Request) {
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
//	//	log.Printf("could not select advertisements ;from db %s", err)
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

	router.HandleFunc("/api/get-listing/{id}", GetListing).Methods("GET")
	//router.HandleFunc("/api/get-ads", GetAllAdvertisements).Methods("GET")

	authGroup := router.PathPrefix("/api/user").Subrouter()
	authGroup.Use(Authentication)
	authGroup.HandleFunc("/create-listing", CreateListing).Methods("POST")
	authGroup.HandleFunc("/profile", GetProfile).Methods("GET")
	authGroup.HandleFunc("/listings", GetListingsForUser).Methods("GET")

	return router
}
