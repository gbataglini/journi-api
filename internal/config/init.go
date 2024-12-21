package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbHost               string `default:"db"`
	DbPassword           string `default:"dev"`
	DbName               string `default:"journi"`
	DbUser               string `default:"postgres"`
	DbPort               int    `default:"5432"`
	GoogleApiKey         string
	OpenWeatherMapApiKey string
}

var c Config

var once sync.Once

func Get() Config {
	once.Do(func() {
		envconfig.MustProcess("journiapi", &c)
	})
	return c
}
