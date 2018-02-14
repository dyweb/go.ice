package common

import (
	"github.com/at15/go.ice/ice/config"
)

// TODO: logging
type ServerConfig struct {
	Verbose         bool                         `yaml:"verbose"`
	DatabaseManager config.DatabaseManagerConfig `yaml:"db-manager"` // TODO: use pointer allow use to detect if it is read from yaml
	Http            config.HttpServerConfig      `yaml:"http"`
	Grpc            config.GrpcServerConfig      `yaml:"grpc"`
}
