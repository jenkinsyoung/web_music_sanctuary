package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/hash"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/models"
	"log"
	"net/http"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error parsing json: %s", err)
	}
	user.Password = hash.PasswordHash(user.Password)
	//TODO: Отправить JWT
	//TODO: нужно добавить бользователя в базу данных
	fmt.Println(user)
	w.WriteHeader(http.StatusOK)
}

func LoggingUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error parsing json: %s", err)
	}
	//TODO: написать авторизацию
}

//TODO: напистаь отправку объявлений по БОЛЬШОЙ категории

func SetupRoutes() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/api/user", NewUser).Methods("POST")
	return router
}
