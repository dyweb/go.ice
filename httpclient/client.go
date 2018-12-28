package httpclient

import (
	"net/http"
)

type Client struct {
	// h is the underlying http.Client, we use value instead of pointer because http.Client (without its methods)
	// is just a config for connection pool, redirect and timeout ...
	h http.Client
}

// Transport returns the transport of the http.Client being used for user to modify tls config etc.
// It returns (nil, false) if the transport is the default transport or the type didn't match.
// You should NOT use http.DefaultTransport in the first place, and modify it is even worse
func (c *Client) Transport() (tr *http.Transport, ok bool) {
	// avoid nil ptr
	if c == nil {
		return
	}
	// get the transport from http.Client
	tr, ok = c.h.Transport.(*http.Transport)
	if !ok {
		return
	}
	// it's default, you should not modify it
	// TODO: might add our own un exported transports
	if tr == http.DefaultTransport {
		return nil, false
	}
	return
}
