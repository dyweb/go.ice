package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dyweb/gommon/errors"
)

func DrainResponseBody(res *http.Response) ([]byte, error) {
	if res == nil {
		return nil, errors.New("http.Response is nil, can not drain and restore body")
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error read all response body")
	}
	if closeErr := res.Body.Close(); closeErr != nil {
		return b, errors.Wrap(closeErr, "error close drained body")
	}
	// restore body, so it can still be used
	res.Body = ioutil.NopCloser(bytes.NewReader(b))
	return b, nil
}

// JoinPath does not sanitize path like path.Join, which would change https:// to https:/, it only remove duplicated
// slashes to avoid // in url i.e. http://myapi.com/api/v1//comments/1
func JoinPath(base string, sub string) string {
	return strings.TrimRight(base, "/") + "/" + strings.TrimLeft(sub, "/")
}
