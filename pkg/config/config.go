package config

import (
	"github.com/caarlos0/env/v10"
)

type Config struct {
	App struct {
		Port string `env:"APP_PORT" envDefault:"8080"`
	}
	Postgres struct {
		Host     string `env:"HOST" envDefault:"localhost"`
		User     string `env:"USER " envDefault:"postgres"`
		Password string `env:"PASSWORD" envDefault:"password"`
		DB       string `env:"DB" envDefault:"lumel_db"`
		Port     string `env:"PORT" envDefault:"5432"`
	} `envPrefix:"POSTGRES_"`
}

var config *Config

func ReadConfig() (*Config, error) {

	config = &Config{}
	if err := env.Parse(config); err != nil {
		panic(err)
	}
	return config, nil
}

func GetConfig() *Config {
	if config == nil {
		config, _ = ReadConfig()
	}
	return config
}
