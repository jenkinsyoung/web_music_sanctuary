package main

import (
	"github.com/jenkinsyoung/web_music_sanctuary/internal/api"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/config"
	_ "github.com/jenkinsyoung/web_music_sanctuary/internal/hash"
	"github.com/jenkinsyoung/web_music_sanctuary/pkg/imgMethods"
	"log"
	"net/http"
)

func main() {
	//TODO: написать бд-ху и забить данными
	cfg := config.MustLoad()
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      api.SetupRoutes(),
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	imgMethods.CreatePhotoDir()

	//_, err := database.ConnectToDB(cfg.Database)
	//if err != nil {
	//	log.Fatalf("Error occured: %s", err)
	//}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Error occured: %s", err)
	}
}
