package apiserver

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	BindAddr string `envconfig:"BIND_ADDR" required:"true"`
	LogLevel string `envconfig:"LOG_LEVEL" required:"true"`
}

func NewConfig() (*Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return &Config{}, err
	}

	return &config, nil
}

func NewConfigMust() *Config {
	config, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	return config
}
