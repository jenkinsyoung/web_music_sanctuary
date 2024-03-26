package main

import (
	"github.com/jenkinsyoung/web_music_sanctuary/internal/api"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/config"
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
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Error occured: %s", err)
	}
}
