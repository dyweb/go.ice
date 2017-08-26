# Go kit - A standard library for microservices

- https://github.com/go-kit/kit
- https://github.com/xephonhq/xephon-k/blob/6e40ceea31e8594824ccc9cae27fd8e05600236c/pkg/server/http_server.go while xephon-k was still using go-kit before https://github.com/xephonhq/xephon-k/pull/55
  - the problem is for each endpoint, I need to create a new go-kit server (under its `transport` package)
- the good thing is, **you got a client library when write a server endpoint** at same time
  - server  json -> decoder -> req -> svc -> resp -> encoder -> json
  - client req -> encoder -> json (c->s) -> decoder (server) -> req (server) -> svc (server) -> resp (server) -> encoder (server) -> json (s->c) -> decoder -> resp
  - example https://github.com/go-kit/kit/blob/master/examples/profilesvc/endpoints.go#L58
  - still the code looks pretty .... wrapper after wrapper ...

## Context

- it does not pass http.Request, but pass its context
- decoder drain the request body and decode into struct

## Service

- the service is an interface, and you can have both client and server implementation
- the service concept in go kit is like a endpoint, i.e. https://gokit.io/examples/stringsvc.html