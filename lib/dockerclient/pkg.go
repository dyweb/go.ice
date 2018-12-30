// Package dockerclient is a docker client that only uses types defined in docker.
// It is an example for httpclient package, it is currently in tree and will be moved to its own repo
package dockerclient

import (
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/dyweb/go.ice/httpclient"
)

type Client struct {
	h *httpclient.Client
}

func New(host string) (*Client, error) {
	if host != "" && !strings.Contains(host, ".sock") {
		// standard docker command accept host without the protocol prefix and use tls flag to indicate https
		if !strings.HasPrefix(host, "http://") {
			host = "http://" + host
		}
	}
	h, err := httpclient.New(host, httpclient.UseJSON())
	if err != nil {
		return nil, err
	}
	return &Client{h: h}, nil
}

func (dc *Client) Ping() (types.Ping, error) {
	var ping types.Ping
	res, err := dc.h.Get(httpclient.Bkg(), "/_ping")
	if err != nil {
		return ping, err
	}
	ping.APIVersion = res.Header.Get("API-Version")
	if res.Header.Get("Docker-Experimental") == "true" {
		ping.Experimental = true
	}
	ping.OSType = res.Header.Get("OSType")
	return ping, nil
}

func (dc *Client) Version() (types.Version, error) {
	var v types.Version
	return v, dc.h.GetTo(httpclient.Bkg(), "/version", &v)
}
