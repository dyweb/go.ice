package logutil

import (
	icelog "github.com/dyweb/go.ice/ice/util/logutil"
	"github.com/dyweb/gommon/log"
)

const Project = "github.com/your/project"

var registry = log.NewLibraryRegistry(Project)

func Registry() *log.Registry {
	return &registry
}

func NewPackageLoggerAndRegistry() (*log.Logger, *log.Registry) {
	logger, child := log.NewPackageLoggerAndRegistryWithSkip(Project, 1)
	registry.AddRegistry(child)
	return logger, child
}

func init() {
	// gain control of important libraries
	registry.AddRegistry(icelog.Registry())
}
