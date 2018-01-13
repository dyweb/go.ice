package server

import (
	"github.com/at15/go.ice/ice/config"
)

type Config struct {
	Verbose         bool
	DatabaseManager config.DatabaseManagerConfig `yaml:"db-manager"` // TODO: use pointer allow use to detect if it is in the yaml
}
