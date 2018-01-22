# Survey of existing Golang web / micro service frameworks

- [Buffalo](gobuffalo.md) https://gobuffalo.io 
  - auth https://github.com/markbates/goth
  - detail html error page when dev
  - provides tasks, repl, watch and restart
- [Echo](echo.md) https://github.com/labstack/echo 
  - context interface and default context implementation
  - wrap http.ResponseWriter to keep track of bytes in out and then log
- [Beego](beego.md) https://github.com/astaxie/beego
  - wrapped http.ResponseWriter but didn't keep track of bytes out  
- [Tango](tango.md) https://github.com/lunny/tango context is too powerful
- [Gin](gin.md) https://github.com/gin-gonic/gin nothing special

## Main issues

Context

- what to put
- how to pass, it's a common practice to put `ctx context.Context` as first argument instead of put in struct, though `http.Request` includes context in a struct (due to compatibility)
- some frameworks like [chi](chi.md) replace default context with own base context in order to pass some value
  - detailed discussion is in [chi.md](chi.md)
  - this loses ability of having the context canceled when client close connection
  - see https://golang.org/pkg/net/http/#CloseNotifier and `net/http/server/go`
  - though people argue if context is carrying too much https://dave.cheney.net/2017/08/20/context-isnt-for-cancellation 
- `context.Context` is an interface, its implementation in standard library are private, `buffalo.Context` is also an interface,
the implementation is in `default_context.go` as `buffalo.DefaultContext`

Access log with out going body size

- wrap original http.ResponseWriter
- [ ] need to pay attention to if original writer is Hijacker
  - [ ] it seems Hijacker is not need for tracking response body size
  - [ ] what does Hijacker really do?
    - The default ResponseWriter for HTTP/1.x connections supports Hijacker, but HTTP/2 connections intentionally do not.
    - you can speak raw TCP based on the example
- echo's seems to be much clear than gorilla's 
- beego, gin, tango does not track response body size though some do wrap http.ResponseWriter  