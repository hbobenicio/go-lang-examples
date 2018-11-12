package config

import (
	"fmt"
	"os"
)

// Config type
type Config struct {
	Port string
}

// New Creates a new Config from Environment
func New() (*Config, error) {
	var c Config
	var missing []string

	port, ok := os.LookupEnv("PORT")
	if ok {
		c.Port = port
	} else {
		missing = append(missing, "PORT")
	}

	if len(missing) == 0 {
		return &c, nil
	}
	return nil, fmt.Errorf("missing env vars: %v", missing)
}
