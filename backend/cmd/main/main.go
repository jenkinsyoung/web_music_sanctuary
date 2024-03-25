package main

import (
	"github.com/jenkinsyoung/web_music_sanctuary/internal/config"
	"log"
	"net/http"
)

func main() {
	cfg := config.MustLoad()
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      nil,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Error occured: %s", err)
	}
}
