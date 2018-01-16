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

## Why gave up go-kit

https://github.com/at15/go.ice/issues/3

Good

- it's good to have the interface and let both server and client implement it
- multiple transport

Bad

- for each endpoint, need to explicit pass encoder, decoder
  - this can simply solved by adding another factory function to avoid factory function for each handler
- don't like pass a function which is a result of function that accepts a function and return a function
  
````go
infoHandler := httptransport.NewServer(
  infoSvcHTTPFactory.MakeEndpoint(infoSvc),
  infoSvcHTTPFactory.MakeDecode(),
  infoSvcHTTPFactory.MakeEncode(),
  options...,
)

func (InfoServiceHTTPFactory) MakeEndpoint(service Service) endpoint.Endpoint {
	infoSvc, ok := service.(InfoService)
	if !ok {
		log.Panic("must pass info service to info service factory")
	}
	// FIXME: the naming here is misleading, the info actually return all the info, more than just version
	// and how to hand things like info/version in go-kit
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		v := infoSvc.Version()

		return infoResponse{Version: v}, nil
	}
}

func (InfoServiceHTTPFactory) MakeDecode() httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		return infoRequest{}, nil
	}
}

func (InfoServiceHTTPFactory) MakeEncode() httptransport.EncodeResponseFunc {
	return httptransport.EncodeJSONResponse
}
````