package server

import (
	"github.com/at15/go.ice/ice/config"
)

// TODO: logging
type Config struct {
	Verbose         bool                         `yaml:"verbose"`
	DatabaseManager config.DatabaseManagerConfig `yaml:"db-manager"` // TODO: use pointer allow use to detect if it is read from yaml
	Http            config.HttpServerConfig      `yaml:"http"`
}
