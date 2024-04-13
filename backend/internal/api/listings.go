package api

import (
	"encoding/json"
	"fmt"
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
	if userID, ok := r.Context().Value("userId").(int64); ok {
		listing := models.Listing{}
		if err := json.NewDecoder(r.Body).Decode(&listing); err != nil {
			log.Printf("error occured decode json to adv %v", err)
		}

		defer r.Body.Close()

		idFlow := make(chan int64, len(listing.ImgList))

		//var guitarInfo models.Guitar

		var wg sync.WaitGroup

		for id, x := range listing.ImgList {
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

		guitarInfo := models.Guitar{
			Form:         "lala",
			PickupConfig: "SSH",
			Category:     "Acoustic",
		}

		guitarID, err := database.DB.CreateGuitar(guitarInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		listingID, err := database.DB.CreateListing(listing, guitarID, userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		wg.Wait()
		defer close(idFlow)

		var idPictureList []int64

		go func(flow chan int64) {
			for {
				value, ok := <-flow
				if !ok {
					return
				}
				database.DB.ImageListingCompound(listingID, value)
			}
		}(idFlow)

		//for range listing.ImgList {
		//	idPictureList = append(idPictureList, <-idFlow)
		//}

		fmt.Println(idPictureList)
		w.WriteHeader(http.StatusOK)
	} else {
		log.Printf("wrong type of payload")
		w.WriteHeader(http.StatusBadRequest)
	}
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

func GetAllListings(w http.ResponseWriter, r *http.Request) {

	fmt.Println("настя сделала запросик")

	listings, err := database.DB.GetListings()
	if err != nil {
		log.Printf("error getting listings from db %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	resp, err := json.Marshal(AllListings{Listings: listings})
	if err != nil {
		log.Printf("error marshalling user info %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(resp)
	w.WriteHeader(http.StatusOK)
}
