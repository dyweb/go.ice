// Package dockerclient is a slim docker client without importing the moby repo.
package dockerclient

import (
	"fmt"

	"github.com/dyweb/gommon/util/httputil"
)

const (
	DefaultVersion   = "1.37"
	DefaultLocalHost = "unix:///var/run/docker.sock"
)

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
