package src

import (
	"github.com/caarlos0/env/v6"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/tomp332/bruteForcer/src/models"
	"log"
)

var GlobalSettings *MainSettings
var ServiceHealth *models.Health

type APISettings struct {
	PaginationDefaultLimit uint `env:"PAGINATION_DEFAULT_LIMIT" envDefault:"50"`
}

type MainSettings struct {
	DBFilePath  string `env:"DB_FILE_PATH,required"`
	ServerPort  int16  `env:"SERVER_PORT" envDefault:"8080"`
	ServerHost  string `env:"SERVER_HOST" envDefault:"localhost"`
	Environment string `env:"ENVIRONMENT" envDefault:"dev"`
	APISettings
}

func LoadSettings() {
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
	GlobalSettings = &cfg
	// Initialize the service health
	ServiceHealth = &models.Health{
		ID:     uuid.New().String(),
		Status: models.PENDING,
		Port:   GlobalSettings.ServerPort,
	}
}
