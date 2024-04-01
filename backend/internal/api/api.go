package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/hash"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/models"
	"github.com/jenkinsyoung/web_music_sanctuary/pkg/imgMethods"
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
//TODO: api/upload-image написать загрузку

func ReceiveImage(w http.ResponseWriter, r *http.Request) {
	photo := models.Photo{}
	json.NewDecoder(r.Body).Decode(&photo)

	defer r.Body.Close()
	//TODO: написать файловую структру для сохраниения картино и присовение им айдишников
	//TODO: к примеру структура storage/image/E4 - картинка
	//TODO: или storage/image/B5
	//TODO: разбиение по секторам (скорее всего по свойствам)

	err := imgMethods.SaveImageBase64(photo.Photo)
	if err != nil {
		log.Printf("Error decode base64 %s", err)
	}
}

func SetupRoutes() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/api/user", NewUser).Methods("POST")
	router.HandleFunc("/api/save-image", ReceiveImage).Methods("POST")
	return router
}
