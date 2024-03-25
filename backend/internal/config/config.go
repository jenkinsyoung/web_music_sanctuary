package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env string `env-default:"local"`
	HTTPServer
	Database
}

type HTTPServer struct {
	Address     string        `env:"HTTP_SERVER_ADDRESS"`
	Timeout     time.Duration `env:"HTTP_SERVER_TIMEOUT"`
	IdleTimeout time.Duration `env:"HTTP_SERVER_IDLE_TIMEOUT"`
}

type Database struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     int    `env:"POSTGRES_PORT"`
	Name     string `env:"POSTGRES_DB"`
	UserName string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
}

func MustLoad() *Config {
	currDir, err := os.Getwd()
	err = godotenv.Load(fmt.Sprintf("%s/backend/.env", currDir))
	if err != nil {
		log.Printf("Warning: unable to load environment variables from .env: %v\n", err)
	}
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Failed to read environment variables: %v\n", err)
	}
	return &cfg
}
