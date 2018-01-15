package logutil

import (
	"github.com/dyweb/gommon/log"
	gommonlog "github.com/dyweb/gommon/util/logutil"
)

var Registry = log.NewLibraryLogger()

func init()  {
	// gain control of important libraries, NOTE: there could be duplicate and cycle when various library is involved
	// thus gommon/log would keep track of visited logger when doing recursive version of SetLevel and SetHandler
	Registry.AddChild(gommonlog.Registry)
}