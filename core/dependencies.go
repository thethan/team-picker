package core

import log "github.com/sirupsen/logrus"

type Dependencies struct {
	Logger *log.Logger
	Config *Config
}

type JWTString string
