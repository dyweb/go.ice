// Package grpc wraps google.golang.org/grpc
package grpc // import "github.com/dyweb/go.ice/ice/transport/grpc"

// TODO: golang/protobuf change https://groups.google.com/forum/#!topic/golang-nuts/F5xFHTfwRnY
// TODO: https://github.com/gogo/protobuf
// https://github.com/xephonhq/xephon-k/tree/master/pkg/server/grpc

import (
	"github.com/dyweb/go.ice/ice/util/logutil"
)

var log, _ = logutil.NewPackageLoggerAndRegistry()
