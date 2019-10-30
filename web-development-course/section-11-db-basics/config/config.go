package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config modela o objeto de configuração da aplicação
type Config struct {
	DBName     string `default:"golang"`
	DBUser     string `default:"golang"`
	DBPassword string `default:"golang"`
	DBHost     string `default:"localhost"`
	DBPort     string `default:"5432"`
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
