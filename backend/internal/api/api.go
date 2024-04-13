package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/sign-up", Register).Methods("POST")
	router.HandleFunc("/api/sign-in", Authorization).Methods("POST")

	router.HandleFunc("/api/listing/{id}", GetListing).Methods("GET")
	router.HandleFunc("/api/listings", GetAllListings).Methods("GET")

	authGroup := router.PathPrefix("/api/user").Subrouter()
	authGroup.Use(Authentication)
	authGroup.HandleFunc("/create-listing", CreateListing).Methods("POST")

	authGroup.HandleFunc("/profile", GetProfile).Methods("GET")
	authGroup.HandleFunc("/listings", GetListingsForUser).Methods("GET")
	authGroup.HandleFunc("/logout", Logout).Methods("GET")

	authGroup.HandleFunc("/profile", UpdateProfile).Methods("PUT")

	return router
}
