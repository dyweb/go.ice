package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/util/httputil"
)

type Client struct {
	// configured by user
	base string
	// json means both request and response are talking in json
	json    bool
	headers map[string]string

	// h is the underlying http.Client
	h *http.Client
}

// TODO: allow user config transport using options and config
func New(base string, opts ...Option) (*Client, error) {
	base = strings.TrimSpace(base)
	if base == "" {
		return nil, errors.New("base path is empty")
	}
	var c *Client
	switch {
	case strings.HasPrefix(base, "http"):
		u, err := url.Parse(base)
		if err != nil {
			return nil, errors.Wrap(err, "error parse http(s) url")
		}
		c = &Client{
			base: u.String(),
			h:    httputil.NewUnPooledClient(),
		}
	case strings.HasPrefix(base, "unix") || strings.HasPrefix(base, "/"):
		// TODO: might use url.Parse to handle unix:// better?
		sock := strings.TrimPrefix(base, "unix://")
		c = &Client{
			base: UnixBasePath,
			h:    httputil.NewClient(httputil.NewUnPooledUnixTransport(sock)),
		}
	default:
		return nil, errors.Errorf("unknown protocol, didn't find http or unix in %s", base)
	}
	return c, applyOptions(c, opts...)
}

func (c *Client) Get(ctx *Context, path string) (*http.Response, error) {
	return c.Do(ctx, httputil.Get, path, nil)
}

func (c *Client) GetTo(ctx *Context, path string, val interface{}) error {
	if !c.json {
		return errors.New("only json encoding is support for GeTo")
	}
	res, err := c.Do(ctx, httputil.Get, path, nil)
	if err != nil {
		return err
	}
	// TODO: the decode to logic should be generated for all methods
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "error drain response body")
	}
	if closeErr := res.Body.Close(); closeErr != nil {
		// TODO: warn it? this should rarely happen
	}
	if err := json.Unmarshal(b, val); err != nil {
		return errors.Wrap(err, "error decode response body as json to struct")
	}
	return nil
}

// TODO: error handling based on status code is required, it should be configured at client and
func (c *Client) Do(ctx *Context, method httputil.Method, path string, body interface{}) (*http.Response, error) {
	if c == nil || c.h == nil {
		return nil, errors.New("client is not initialized")
	}

	var (
		encodedBody io.Reader
		err         error
	)
	if body != nil {
		if encodedBody, err = encodeBody(body, c.json); err != nil {
			return nil, err
		}
	}
	u := JoinPath(c.base, path)
	req, err := http.NewRequest(string(method), u, encodedBody)
	if err != nil {
		return nil, errors.Wrap(err, "error create http request")
	}
	// TODO: I think range through nil map is a noop?
	if len(c.headers) > 0 {
		for k, v := range c.headers {
			req.Header.Set(k, v)
		}
	}
	if len(ctx.headers) > 0 {
		for k, v := range ctx.headers {
			req.Header.Set(k, v)
		}
	}
	if len(ctx.params) > 0 {
		q := req.URL.Query()
		for k, v := range ctx.params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	req = req.WithContext(ctx)
	res, err := c.h.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func encodeBody(body interface{}, encodeToJson bool) (io.Reader, error) {
	if _, ok := body.(io.Reader); ok {
		return body.(io.Reader), nil
	}
	switch body.(type) {
	case io.Reader:
		return body.(io.Reader), nil
	case []byte:
		return bytes.NewReader(body.([]byte)), nil
	case string:
		return strings.NewReader(body.(string)), nil
	}
	if !encodeToJson {
		return nil, errors.New("request body must be io.Reader/bytes/string or set the client to auto encode request to json")
	}
	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, errors.Wrap(err, "error encode request body to json")
	}
	return &buf, nil
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
