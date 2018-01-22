package logutil

import (
	"github.com/dyweb/gommon/log"
	icelog "github.com/at15/go.ice/ice/util/logutil"
)

var Registry = log.NewApplicationLogger()

func NewPackageLogger() *log.Logger {
	l := log.NewPackageLoggerWithSkip(1)
	Registry.AddChild(l)
	return l
}

func init() {
	// gain control of important libraries
	Registry.AddChild(icelog.Registry)
}
