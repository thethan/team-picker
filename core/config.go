package core

import (
	"github.com/caarlos0/env"
)

type Config struct {
	Logging *Logging
	JWTKeyString string `env:"JWT_SIGNING_KEY" envDefault:"foo.bar"`
	DataLocation string `env:"DATA_LOCATION" envDefault:"../data/"`
}


func SetConfig() (*Config, error) {
	config := Config{}

	config.Logging = &Logging{}

	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	return &config, nil
}

type Logging struct {
	LogLevel uint `env:"LOGGING_LEVEL" envDefault:"5"` // 5 is Debug and 3 is Error
}