package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT" envDefault:"8000"`
	Password string `env:"PASSWORD,unset"`
}

// Validate -
func (c Config) Validate() error {
	panic("implement me")
}

func ParseFromEnv() (Config, error) {
	// https://github.com/caarlos0/env
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

}
