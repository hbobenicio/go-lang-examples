package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config modela o objeto de configuração da aplicação
type Config struct {
	DBName     string `envconfig:"DB_NAME" default:"golang"`
	DBUser     string `envconfig:"DB_USER" default:"golang"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"golang"`
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     string `envconfig:"DB_PORT" default:"5432"`
	DBSSLMode  string `envconfig:"DB_NAME" default:"disable"`
}

// New cria uma nova instância "ZeroValue" de config
func New() *Config {
	return &Config{}
}

// LoadFromEnv carrega as configurações pelas variáveis de ambiente
func (c *Config) LoadFromEnv() {
	if err := envconfig.Process("APP", c); err != nil {
		log.Fatalln(err)
	}
}
