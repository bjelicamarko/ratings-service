package config

import (
	"log"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type Config struct {
	Host     string `env:"HOST,required" envDefault:"localhost"`
	User     string `env:"USER,required"`
	Password string `env:"PASSWORD,required"`
	DbName   string `env:"DBNAME,required"`
	Port     string `env:"PORT,required"`
}

func ReturnConfig() Config {
	var cfg = Config{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	return cfg
}
