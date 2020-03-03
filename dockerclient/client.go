package dockerclient

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/httpclient"
	"github.com/dyweb/gommon/util/httputil"
)

type Client struct {
	version string
	h       *httpclient.Client
}

func New(host string) (*Client, error) {
	if host != "" && !strings.Contains(host, ".sock") {
		// standard docker command accept host without the protocol prefix and use tls flag to indicate https
		if !strings.HasPrefix(host, "http://") {
			host = "http://" + host
		}
	}
	h, err := httpclient.New(
		host,
		httpclient.UseJSON(),
		httpclient.WithErrorHandlerFunc(DecodeDockerError),
	)
	if err != nil {
		return nil, err
	}
	return &Client{
		version: DefaultVersion,
		h:       h,
	}, nil
}

func (dc *Client) Ping() (types.Ping, error) {
	var ping types.Ping
	res, err := dc.h.GetRaw(httpclient.Bkg(), "/_ping")
	if err != nil {
		return ping, err
	}
	defer httpclient.DrainAndClose(res)
	ping.APIVersion = res.Header.Get("API-Version")
	if res.Header.Get("Docker-Experimental") == "true" {
		ping.Experimental = true
	}
	ping.OSType = res.Header.Get("OSType")
	return ping, nil
}

func (dc *Client) Version() (types.Version, error) {
	var v types.Version
	return v, dc.h.Get(httpclient.Bkg(), "/version", &v)
}

func DecodeDockerError(status int, body []byte, res *http.Response) (decodedError error) {
	e := ErrDocker{
		Status: status,
		Method: httputil.Method(res.Request.Method),
		Url:    res.Request.URL.String(),
		Path:   res.Request.URL.Path,
		Body:   string(body),
	}
	// try to decode docker's error message, which is just a single string, well designed ...
	var derr types.ErrorResponse
	if err := json.Unmarshal(body, &derr); err != nil {
		errors.Ignore(err)
	}
	e.Message = derr.Message
	return &e
}
