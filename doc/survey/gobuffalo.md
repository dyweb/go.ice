# Buffalo - Rapid Web Development in Go

- https://gobuffalo.io/
- https://github.com/gobuffalo/buffalo

## Take away

- context interface and default impl (though don't know how it deal with default http.Context)
- auth https://github.com/markbates/goth
- rake like one time task (migration etc.)
- detail html error page when dev
- background worker (no cron?)

## Request handling

### Routing

- wrap around gorilla mux
- resource CRUD

````go
type Resource interface {
  List(Context) error
  Show(Context) error
  New(Context) error
  Create(Context) error
  Edit(Context) error
  Update(Context) error
  Destroy(Context) error
}
````

### Middleware

- group middleware
- skip middleware

### Context

> The buffalo.Context interface supports context.Context so it can be passed around and used as a "standard" Go Context.

````go
type Context interface {
  context.Context
  Response() http.ResponseWriter
  Request() *http.Request
  Session() *Session
  Params() ParamValues
  Param(string) string
  Set(string, interface{})
  LogField(string, interface{})
  LogFields(map[string]interface{})
  Logger() Logger
  Bind(interface{}) error
  Render(int, render.Renderer) error
  Error(int, error) error
  Websocket() (*websocket.Conn, error)
  Redirect(int, string, ...interface{}) error
  Data() map[string]interface{}
}
````

- https://github.com/gobuffalo/buffalo/blob/master/default_context.go is the default implementation
- [ ] TODO: in its test, context.Background() is used, why it does not make use of context in http.Request

````go
type DefaultContext struct {
	context.Context
	response    http.ResponseWriter
	request     *http.Request
	params      url.Values
	logger      Logger
	session     *Session
	contentType string
	data        map[string]interface{}
	flash       *Flash
}
````

### Error Handling

- https://gobuffalo.io/docs/errors
- detailed html error page in dev

### Session

- cookie based requires gob, guess using gorilla/securecookie as well

### Cookie

## Task

- https://github.com/markbates/grift like rake in rails
- buffalo task list

## Auth

- https://github.com/markbates/goth many many many providers ...

## Database

- migration using DSL called fizz https://github.com/markbates/pop/tree/master/fizz

## Background Job Workers

````go
type Worker interface {
  // Start the worker with the given context
  Start(context.Context) error
  // Stop the worker
  Stop() error
  // Perform a job as soon as possibly
  Perform(Job) error
  // PerformAt performs a job at a particular time
  PerformAt(Job, time.Time) error
  // PerformIn performs a job after waiting for a specified amount of time
  PerformIn(Job, time.Duration) error
  // Register a Handler
  Register(string, Handler) error
}
````

## REPL

- based on https://github.com/motemen/gore
- pre-load actions & models

## Plugins

- for the `buffalo` command
- can be written in other language