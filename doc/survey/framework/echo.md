# Echo - High performance, extensible, minimalist Go web framework

## Routing

own radix tree based implementation, currently router is not our main concern, using gorilla is enough

## Context

- use interface (same as buffalo)
- need middleware and type assert when use custom context

````go
// default impl, it has a ptr to echo instance (too much info ...)
context struct {
		request  *http.Request
		response *Response
		path     string
		pnames   []string
		pvalues  []string
		query    url.Values
		handler  HandlerFunc
		store    Map
		echo     *Echo
	}
```` 

````go
type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo() {
	println("foo")
}

// convert to custom context
e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{c}
		return h(cc)
	}
})

// type assert before use
e.GET("/", func(c echo.Context) error {
	cc := c.(*CustomContext)
	cc.Foo()
	return cc.String(200, "OK")
})
````

## Response

- a custom http.ResponseWriter implementation (kind of like hijacker?), the Logger middle use this to get the size of data written

````go
type Response struct {
		echo        *Echo
		beforeFuncs []func()
		Writer      http.ResponseWriter
		Status      int
		Size        int64
		Committed   bool
	}
````

## Error Handling

> Echo advocates for centralized HTTP error handling by returning error from middleware and handlers. Centralized error handler allows us to log errors to external services from a unified location and send a customized HTTP response to the client

## Middleware

### Body Limit

````go
return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()

			// Based on content length
			if req.ContentLength > config.limit {
				return echo.ErrStatusRequestEntityTooLarge
			}

			// Based on content read
			r := pool.Get().(*limitedReader)
			r.Reset(req.Body, c)
			defer pool.Put(r)
			req.Body = r

			return next(c)
		}
	}
````

### Recover

````go
return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			defer func() {
				if r := recover(); r != nil {
					var err error
					switch r := r.(type) {
					case error:
						err = r
					default:
						err = fmt.Errorf("%v", r)
					}
					stack := make([]byte, config.StackSize)
					length := runtime.Stack(stack, !config.DisableStackAll)
					if !config.DisablePrintStack {
						c.Logger().Printf("[%s] %s %s\n", color.Red("PANIC RECOVER"), err, stack[:length])
					}
					c.Error(err)
				}
			}()
			return next(c)
		}
	}
````

### Logger

make use of context's Request and Response, Response is a custom writer wrap around default writer