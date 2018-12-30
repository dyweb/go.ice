package httpclient_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dyweb/go.ice/httpclient"
	"github.com/dyweb/gommon/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContext_SetParam(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/param", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		dump := make(map[string][]string)
		for key, v := range q {
			dump[key] = v
		}
		writeJSON(t, w, dump)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()
	client, err := httpclient.New(srv.URL, httpclient.UseJSON())
	require.Nil(t, err)

	ctx := httpclient.Bkg().SetParam("foo", "bar")
	dump := make(map[string][]string)
	assert.Nil(t, client.GetTo(ctx, "/param", &dump))
	assert.Equal(t, "bar", dump["foo"][0])
}

func TestContext_SetErrorHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/404/nobody", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()
	client, err := httpclient.New(srv.URL, httpclient.UseJSON())
	require.Nil(t, err)

	h := httpclient.ErrorHandlerFunc(func(status int, body []byte, res *http.Response) error {
		return errors.Errorf("custom %d", status)
	})
	ctx := httpclient.Bkg().SetErrorHandler(h)
	res, err := client.Get(ctx, "/404/nobody")
	assert.NotNil(t, res, "application error has no nil response")
	assert.Equal(t, "custom 404", err.Error())
}
