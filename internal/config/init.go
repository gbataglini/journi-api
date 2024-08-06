package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbHost string
	DbPassword string
}

var c Config

var once sync.Once

func Get() Config {
	once.Do(func (){ 
		envconfig.MustProcess("journiapi", &c)
	})
	return c
}

