# 2018-12-26 HTTP Client 

This doc is survey for [#37](https://github.com/dyweb/go.ice/issues/37) A high level wrapper around net/http.

Goals

- [ ] types and constants, they are not well defined in existing net/http, see https://github.com/bradfitz/exp-httpclient
- [ ] context, might accept context.Context but do convert under the hood
- [ ] retry
- [ ] built in json encode/decode
- [ ] make the default easy to use

## Types

https://github.com/bradfitz/exp-httpclient/blob/master/http/method.go

- define method using type is way better than just using string
- enforce the req/res based on method is good, though actually both client and server can disobey the convention ..

````go
// Method is an HTTP method. Although HTTP methods are case insensitive,
// values of this type must contain only capital letters.
type Method string

// Common HTTP methods.
//
// Unless otherwise noted, these are defined in RFC 7231 section 4.3.
const (
	Get     Method = "GET"
	Head    Method = "HEAD"
	Post    Method = "POST"
	Put     Method = "PUT"
	Patch   Method = "PATCH" // RFC 5789
	Delete  Method = "DELETE"
	Connect Method = "CONNECT"
	Options Method = "OPTIONS"
	Trace   Method = "TRACE"
)

func (m Method) RequestBodyAllowed() bool {
	panic("TODO")
}

func (m Method) RequestBodyCommon() bool {
	panic("TODO")
}

func (m Method) ResponseBodyAllowed() bool {
	panic("TODO")
}
````

https://github.com/bradfitz/exp-httpclient/blob/master/http/status.go

- make status a struct
  - [ ] I didn't know there is reason phrase for http 1.1, but based on comment, http 2 no longer have it
- [ ] didn't add is redirect
  - [x] go's http client is handling redirect automatically? by default, yes
  - https://stackoverflow.com/questions/23297520/how-can-i-make-the-go-http-client-not-follow-redirects-automatically
    - set `CheckRedirect` and return `http.ErrUseLastResponse`

````go
// IsSuccess reports whether s is in the Successful (2xx) status code class,
// as defined by RFC 7231 section 6.3.
func (s Status) IsSuccess() bool {
	return s.code >= 200 && s.code <= 299
}

// IsClientError reports whether s is in the Client Error (4xx) status code class,
// as defined by RFC 7231 section 6.5.
func (s Status) IsClientError() bool {
	return s.code >= 400 && s.code <= 499
}

// IsServerError reports whether s is in the Server Error (5xx) status code class,
// as defined by RFC 7231 section 6.6.
func (s Status) IsServerError() bool {
	return s.code >= 500 && s.code <= 599
}

// IsNotModified reports whether s is the 304 Not Modified status.
func (s Status) IsNotModified() bool { return s.code == 304 }
````

https://github.com/bradfitz/exp-httpclient/blob/master/httpclient/example_test.go

- allow decode json into struct automatically, it is relying on generic, but whe can use the `Unmarshal(b, ptrToVal)` way and use `interface{}` for now

## Go http clients

- https://awesome-go.com/#http-clients
- https://golanglibs.com/category/http-clients?sort=top has more w/ more stars

## httputil

- its [DumpRequestOut](https://golang.org/src/net/http/httputil/dump.go?s=1848:1913#L56) is saving request body
- `drainBody` use `ReadFrom`

### gorequest

- https://github.com/parnurzeal/gorequest
- builder

````go
request := gorequest.New()
resp, body, errs := request.Get("http://example.com").
  RedirectPolicy(redirectPolicyFunc).
  Set("If-None-Match", `W/"wyzzy"`).
  End()
````

### resty

- https://github.com/go-resty/resty
- builder + decode

````go
resp, err := resty.R().
      SetHeader("Content-Type", "application/json").
      SetBody(`{"username":"testuser", "password":"testpass"}`).
      SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
      Post("https://myapp.com/login")
````

### grequests

- https://github.com/levigross/grequests

````go
response := Get("http://some-wonderful-file.txt", nil)

// This call to .Bytes caches the request bytes in an internal byte buffer – which can be used again and again until it is cleared
response.Bytes() == `file-bytes`
response.String() == "file-string"
````

### Gentleman

- https://github.com/h2non/gentleman/blob/master/context/context.go has its own context that used for middleware
- https://github.com/h2non/gentleman/tree/master/middleware 
- https://github.com/h2non/gentleman/blob/master/request.go its own request struct that contains context, middles 
- https://github.com/h2non/gentleman/blob/master/response.go Response have `SaveToFile`, `JSON` etc.
- https://github.com/h2non/gentleman/tree/master/plugins most of them can just be an config option ...

## Heimdall

- https://github.com/gojektech/heimdall nice logo
- has retry and hystrix circuit breaking

## Sling

- https://github.com/dghubble/sling
- a builder style
- can **use struct for query parameters** https://github.com/dghubble/sling#querystruct

## Other language

Python

- requests http://docs.python-requests.org/en/master/
  - have a json method to return decoded json

````python
r = requests.get('https://api.github.com/events')
r.json()
r.raw

# stream
r = requests.get('https://api.github.com/events', stream=True)
````

Javascript

- https://github.com/axios/axios
  - `all` and `spread allow callback after both are complete

```js
function getUserAccount() {
  return axios.get('/user/12345');
}

function getUserPermissions() {
  return axios.get('/user/12345/permissions');
}

axios.all([getUserAccount(), getUserPermissions()])
  .then(axios.spread(function (acct, perms) {
    // Both requests are now complete
  }));
```


## Retry

- same the body
- backoff strategy

- https://github.com/hashicorp/go-retryablehttp
- https://github.com/sethgrid/pester/blob/master/pester.go
- https://github.com/avast/retry-go a `retry.Do` wrapper
- https://medium.com/@nitishkr88/http-retries-in-go-e622e51d249f