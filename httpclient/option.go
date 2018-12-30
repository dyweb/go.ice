package httpclient

import (
	"net/http"

	"github.com/dyweb/gommon/errors"
)

type Option func(c *Client) error

func UseJSON() Option {
	return func(c *Client) error {
		c.json = true
		return nil
	}
}

func WithTransport(tr *http.Transport) Option {
	return func(c *Client) error {
		if c.h == nil {
			return errors.New("native http client is empty, can't set transport")
		}
		c.h.Transport = tr
		return nil
	}
}

func applyOptions(c *Client, opts ...Option) error {
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return errors.Wrap(err, "error apply option")
		}
	}
	return nil
}
