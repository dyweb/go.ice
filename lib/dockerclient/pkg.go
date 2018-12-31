// Package dockerclient is a docker client that only uses types defined in docker.
// It is an example for httpclient package, it is currently in tree and will be moved to its own repo
package dockerclient

import (
	"fmt"

	"github.com/dyweb/gommon/util/httputil"
)

const DefaultVersion = "1.37"

type ErrDocker struct {
	Method httputil.Method
	Url    string
	Path   string
	Status int
	// Message is the decoded error message from docker daemon
	Message string
	Body    string
}

func (e *ErrDocker) Error() string {
	msg := e.Message
	if e.Message == "" {
		msg = e.Body
	}
	return fmt.Sprintf("docke err %s from %d %s %s", msg, e.Status, e.Method, e.Url)
}
