package ctx

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestStd_HttpClientConext(t *testing.T) {
	reqCtx, _ := context.WithTimeout(context.Background(), 1*time.Millisecond)
	req, err := http.NewRequest(http.MethodGet, "https://github.com", nil)
	if err != nil {
		t.Log(err)
		return
	}
	// NOTE: you must assign the returned *http.Request to your original variable, it's NOT in place update of context
	req = req.WithContext(reqCtx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log(err)
		return
	}
	io.Copy(os.Stdout, res.Body)
	res.Body.Close()
}

func TestStd_HttpConextClientCloseConnection(t *testing.T) {
	// FIXME: it seems when http.Client have context deadline exceeded in Do, the connection is not closed, writer still works

	mux := http.NewServeMux()
	// fake remote service call by sleeping
	mux.HandleFunc("/remote", func(writer http.ResponseWriter, request *http.Request) {
		t.Log("start sleeping", time.Now().UnixNano())
		time.Sleep(10 * time.Millisecond)
		t.Log("sleep finished", time.Now().UnixNano())
		n, err := writer.Write([]byte("remote is here"))
		if err != nil {
			t.Log("/remote write error ", err)
		} else {
			t.Log("/remote wrote ", n)
		}
	})
	mux.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
		remoteUrl := request.URL.Query().Get("remote")
		t.Log("remote is ", remoteUrl, time.Now().UnixNano(), "\n")
		client := http.Client{}
		outReq, err := http.NewRequest(http.MethodGet, remoteUrl, nil)
		if err != nil {
			t.Fatal(err)
		}
		outReq = outReq.WithContext(request.Context())
		res, err := client.Do(outReq)
		// NOTE: when the incoming request is canceled by client, our outgoing request is also canceled
		if err != nil {
			t.Log("server side error", err)
			return
		}
		resB, _ := ioutil.ReadAll(res.Body)
		t.Log("remote response is: ", string(resB), time.Now().UnixNano(), "\n")
		res.Body.Close()
		n, err := writer.Write([]byte("index is here"))
		// FIXME: it seems when client have context deadline exceeded, the connection is not closed, writer still works
		if err != nil {
			t.Log("/index write error ", err)
		} else {
			t.Log("/index wrote ", n)
		}
	})
	ts := httptest.NewServer(mux)
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL+"/index", nil)
	if err != nil {
		t.Fatal(err)
		return
	}
	q := req.URL.Query()
	q.Set("remote", ts.URL+"/remote")
	req.URL.RawQuery = q.Encode()
	reqCtx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	// NOTE: you must assign the returned *http.Request to your original variable, it's NOT in place update of context
	req = req.WithContext(reqCtx)
	client := &http.Client{Transport: &http.Transport{DisableKeepAlives: true}}
	res, err := client.Do(req)
	if err != nil {
		t.Log("client side error ", err)
	} else {
		t.Log("client got response ", time.Now().UnixNano(), "\n")
		io.Copy(os.Stdout, res.Body)
		res.Body.Close()
	}
	time.Sleep(10 * time.Millisecond)
}
