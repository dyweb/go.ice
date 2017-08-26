# Gin 

- https://github.com/gin-gonic/gin

## Take away

- the bind on context for different content type, json, xml, protobuf etc.

## Context

````go
// Context is the most important part of gin. It allows us to pass variables between middleware,
// manage the flow, validate the JSON of a request and render a JSON response for example.
type Context struct {
	writermem responseWriter
	Request   *http.Request
	Writer    ResponseWriter

	Params   Params
	handlers HandlersChain
	index    int8

	engine *Engine

	// Keys is a key/value pair exclusively for the context of each request.
	Keys map[string]interface{}

	// Errors is a list of errors attached to all the handlers/middlewares who used this context.
	Errors errorMsgs

	// Accepted defines a list of manually accepted formats for content negotiation.
	Accepted []string
}
````

no wrapper around response, so can't get written content length, its http log does NOT log it as well

````go
// write json
c.JSON(200, ....)
// actually calls the render
// which writes to http.ResponseWriter
````