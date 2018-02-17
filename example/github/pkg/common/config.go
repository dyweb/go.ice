package common

import (
	"github.com/at15/go.ice/ice/config"
)

// TODO: logging config, need to wait for gommon/log
// TODO: use pointer allow use to detect if it is read from yaml
type ServerConfig struct {
	Verbose         bool                         `yaml:"verbose"`
	DatabaseManager config.DatabaseManagerConfig `yaml:"db-manager"`
	Http            config.HttpServerConfig      `yaml:"http"`
	Grpc            config.GrpcServerConfig      `yaml:"grpc"`
	Tracing         config.TracingConfig         `yaml:"tracing"`
}
