// Package cli is a wrapper around spf13/cobra
package cli

import (
	dlog "github.com/dyweb/gommon/log"
)

var logReg = dlog.NewRegistry()
var log = logReg.Logger()
