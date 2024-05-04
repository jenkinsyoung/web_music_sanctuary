package main

import (
	"github.com/jenkinsyoung/web_music_sanctuary/internal/api"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/config"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/database"
	_ "github.com/jenkinsyoung/web_music_sanctuary/internal/hash"
	"github.com/jenkinsyoung/web_music_sanctuary/pkg/imgMethods"
	"github.com/rs/cors"
	"log"
	"net/http"
)

//			handlers.AllowedMethods([]string{"*"}),
//			handlers.AllowedHeaders([]string{"*"}),
//			handlers.AllowedOrigins([]string{"*"})

func main() {
	cfg := config.MustLoad()
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      cors.Default().Handler(api.SetupRoutes()),
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	imgMethods.CreatePhotoDir()
	if err := database.ConnectToDB(cfg.Database); err != nil {
		log.Fatalf("Error occured: %s", err)
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Error occured: %s", err)
	}
}
