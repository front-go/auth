package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres Postgres
	Service  Service
}

type Postgres struct {
	Host     string `env:"AUTH_SERVICE_POSTGRES_HOST"`
	Port     string `env:"AUTH_SERVICE_POSTGRES_PORT"`
	User     string `env:"AUTH_SERVICE_POSTGRES_USER"`
	Password string `env:"AUTH_SERVICE_POSTGRES_PASSWORD"`
	Database string `env:"AUTH_SERVICE_POSTGRES_DBNAME"`
}

type Service struct {
	Port string `env:"AUTH_SERVICE_PORT"`
}

func MustLoad() *Config {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatalf("Can not read env variables: %s", err)
	}
	return cfg
}
