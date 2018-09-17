package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var (
	envInstance     Env
	envInstanceOnce sync.Once
)

func Get() Env {
	envInstanceOnce.Do(func() {
		godotenv.Load()

		envInstance = Env{}
		if err := env.Parse(&envInstance); err != nil {
			log.Fatal("Failed to load ENV")
		}
	})
	return envInstance
}

type Env struct {
	WykopAppKey string `env:"WYKOP_APP_KEY" envDefault:""`
}
