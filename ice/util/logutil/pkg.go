package logutil

import (
	"github.com/dyweb/gommon/log"
	gommonlog "github.com/dyweb/gommon/util/logutil"
)

const Project = "github.com/dyweb/go.ice"

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
	// gain control of important libraries, NOTE: there could be duplicate and cycle when various library is involved
	// thus gommon/log would keep track of visited logger when doing recursive version of SetLevel and SetHandler
	registry.AddRegistry(gommonlog.Registry())
}
