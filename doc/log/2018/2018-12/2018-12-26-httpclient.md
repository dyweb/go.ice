# 2018-12-26 HTTP Client 

This doc is survey for [#37](https://github.com/dyweb/go.ice/issues/37) A high level wrapper around net/http.

Goals

- [ ] types and constants, they are not well defined in existing net/http, see https://github.com/bradfitz/exp-httpclient
- [ ] context, might accept context.Context but do convert under the hood
- [ ] retry
- [ ] built in json encode/decode

## Retry

- same the body
- backoff strategy

- https://github.com/hashicorp/go-retryablehttp
- https://github.com/sethgrid/pester/blob/master/pester.go
- https://github.com/avast/retry-go a `retry.Do` wrapper
- https://medium.com/@nitishkr88/http-retries-in-go-e622e51d249f