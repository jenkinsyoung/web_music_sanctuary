package api

import (
	"encoding/json"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/database"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/models"
	"log"
	"net/http"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	if userID, ok := r.Context().Value("userId").(int64); ok {
		user := database.DB.GetUserInfoByID(userID)
		resp, err := json.Marshal(user)
		if err != nil {
			log.Printf("error marshalling user info %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(resp)
		w.WriteHeader(http.StatusOK)
	} else {
		log.Printf("wrong type of payload")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	if userID, ok := r.Context().Value("userId").(int64); ok {
		user := models.User{}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Printf("Error parsing json: %s", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		user.Id = userID
		err := database.DB.UpdateUserInfo(user)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp, err := json.Marshal(user)
		if err != nil {
			log.Printf("error marshalling user info %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(resp)
		w.WriteHeader(http.StatusOK)
	} else {
		log.Printf("wrong type of payload")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func GetListingsForUser(w http.ResponseWriter, r *http.Request) {
	if userID, ok := r.Context().Value("userId").(int64); ok {
		listings, err := database.DB.GetUserListings(userID)
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
	} else {
		log.Printf("wrong type of payload")
		w.WriteHeader(http.StatusBadRequest)
	}
}

//func UpdateListing(w http.ResponseWriter, r *http.Request) {
//	if userID, ok := r.Context().Value("userId").(int64); ok {
//
//		if err != nil {
//			log.Printf("error getting listings from db %s", err)
//			w.WriteHeader(http.StatusInternalServerError)
//		}
//		resp, err := json.Marshal(AllListings{Listings: listings})
//		if err != nil {
//			log.Printf("error marshalling user info %s", err)
//			w.WriteHeader(http.StatusInternalServerError)
//		}
//
//		w.Write(userID)
//		w.WriteHeader(http.StatusOK)
//	} else {
//		log.Printf("wrong type of payload")
//		w.WriteHeader(http.StatusBadRequest)
//	}
//}
