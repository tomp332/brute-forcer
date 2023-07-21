package src

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

type MainSettings struct {
	DatabaseConnectionUrl string `env:"DB_URL,required"`
	ServerPort            int16  `env:"SERVER_PORT" envDefault:"8080"`
	ServerHost            string `env:"SERVER_HOST" envDefault:"localhost"`
	Environment           string `env:"ENVIRONMENT" envDefault:"dev"`
}

func LoadSettings() MainSettings {
	// Loading the environment variables from '.env' file.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	cfg := MainSettings{} // 👈 new instance of `Config`
	err = env.Parse(&cfg) // 👈 Parse environment variables into `Config`
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}
	log.Printf("Settings loaded successfully: %+v\n", cfg)
	return cfg
}
