// Package http wraps net/http
package http // import "github.com/dyweb/go.ice/ice/transport/http"

import (
	"github.com/dyweb/go.ice/ice/util/logutil"
)

// TODO: client, may need to drain https://github.com/at15/go-solr/issues/2#issuecomment-361133772

var log, _ = logutil.NewPackageLoggerAndRegistry()
