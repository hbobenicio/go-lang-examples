package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerPort string `envconfig:"SERVER_PORT" default:"8080"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) LoadFromEnv() {
	err := envconfig.Process("APP", c)
	if err != nil {
		log.Fatalln("error: config load:", err)
	}
}
