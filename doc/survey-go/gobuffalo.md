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

To use your own context, a middleware is needed to wrap the default context to your own context

- [ ] TODO: is a type assertion needed in handler? based on same mechanism in echo, it is

````go
type MyContext struct {
  buffalo.Context
}

func (my MyContext) Error(status int, err error) error {
  my.Logger().Fatal(err)
  return err
}

func App() *buffalo.App {
  if app != nil {
    // ...
    app.Use(func (next buffalo.Handler) buffalo.Handler {
      return func(c buffalo.Context) error {
      // change the context to MyContext
      return next(MyContext{c})
      }
    })
    // ...
  }
  return app
}
````

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
  - similar to https://github.com/magefile/mage
- buffalo task list

## Auth

- https://github.com/markbates/goth many many many providers ...

## Database

- migration using DSL called fizz https://github.com/markbates/pop/tree/master/fizz
- https://github.com/markbates/pop wraps https://github.com/jmoiron/sqlx
- https://gobuffalo.io/docs/db
  - `develop`, `test`, `production`, supports `{{ envOr ... }}` in `database.yml`
  - support using `soda` to create, drop, migrate database based on config
- won't connect when application start, but on first incoming request
- `app/actions/app.go` use the `PopTransaction` middleware, which start a new transaction for each incoming request
- `app/models.DB` just check the config, does not really open the connection, at least for postgres
  - maybe not the case for sqlite
  
````go
// actions/app.go
func App() *buffalo.App {
    app.Use(middleware.PopTransaction(models.DB))
}
````

````go
// models/models.go
package models

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/markbates/pop"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection

func init() {
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"
}
````

````go
// buffalo/middleware/pop_transaction.go
var PopTransaction = func(db *pop.Connection) buffalo.MiddlewareFunc {
	return func(h buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			// wrap all requests in a transaction and set the length
			// of time doing things in the db to the log.
			err := db.Transaction(func(tx *pop.Connection) error {
				start := tx.Elapsed
				defer func() {
					finished := tx.Elapsed
					elapsed := time.Duration(finished - start)
					c.LogField("db", elapsed)
				}()
				c.Set("tx", tx)
				if err := h(c); err != nil {
					return err
				}
				if res, ok := c.Response().(*buffalo.Response); ok {
					if res.Status < 200 || res.Status >= 400 {
						return errNonSuccess
					}
				}
				return nil
			})
			if err != nil && errors.Cause(err) != errNonSuccess {
				return err
			}
			return nil
		}
	}
}
````

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

- [ ] FIXME: remove those packages from global path, it's not a good idea to have so many libraries in go path

````text
 go get -u -v github.com/gobuffalo/buffalo/buffalo
github.com/gobuffalo/buffalo (download)
github.com/fatih/color (download)
github.com/gobuffalo/envy (download)
github.com/joho/godotenv (download)
github.com/mitchellh/go-homedir (download)
github.com/gobuffalo/makr (download)
github.com/markbates/inflect (download)
github.com/pkg/errors (download)
github.com/sirupsen/logrus (download)
Fetching https://golang.org/x/crypto/ssh/terminal?go-get=1
Parsing meta tags from https://golang.org/x/crypto/ssh/terminal?go-get=1 (status code 200)
get "golang.org/x/crypto/ssh/terminal": found meta tag get.metaImport{Prefix:"golang.org/x/crypto", VCS:"git", RepoRoot:"https://go.googlesource.com/crypto"} at https://golang.org/x/crypto/ssh/terminal?go-get=1
get "golang.org/x/crypto/ssh/terminal": verifying non-authoritative meta tag
Fetching https://golang.org/x/crypto?go-get=1
Parsing meta tags from https://golang.org/x/crypto?go-get=1 (status code 200)
golang.org/x/crypto (download)
Fetching https://golang.org/x/sys/unix?go-get=1
Parsing meta tags from https://golang.org/x/sys/unix?go-get=1 (status code 200)
get "golang.org/x/sys/unix": found meta tag get.metaImport{Prefix:"golang.org/x/sys", VCS:"git", RepoRoot:"https://go.googlesource.com/sys"} at https://golang.org/x/sys/unix?go-get=1
get "golang.org/x/sys/unix": verifying non-authoritative meta tag
Fetching https://golang.org/x/sys?go-get=1
Parsing meta tags from https://golang.org/x/sys?go-get=1 (status code 200)
golang.org/x/sys (download)
github.com/gobuffalo/packr (download)
Fetching https://golang.org/x/sync/errgroup?go-get=1
Parsing meta tags from https://golang.org/x/sync/errgroup?go-get=1 (status code 200)
get "golang.org/x/sync/errgroup": found meta tag get.metaImport{Prefix:"golang.org/x/sync", VCS:"git", RepoRoot:"https://go.googlesource.com/sync"} at https://golang.org/x/sync/errgroup?go-get=1
get "golang.org/x/sync/errgroup": verifying non-authoritative meta tag
Fetching https://golang.org/x/sync?go-get=1
Parsing meta tags from https://golang.org/x/sync?go-get=1 (status code 200)
golang.org/x/sync (download)
Fetching https://golang.org/x/net/context?go-get=1
Parsing meta tags from https://golang.org/x/net/context?go-get=1 (status code 200)
get "golang.org/x/net/context": found meta tag get.metaImport{Prefix:"golang.org/x/net", VCS:"git", RepoRoot:"https://go.googlesource.com/net"} at https://golang.org/x/net/context?go-get=1
get "golang.org/x/net/context": verifying non-authoritative meta tag
Fetching https://golang.org/x/net?go-get=1
Parsing meta tags from https://golang.org/x/net?go-get=1 (status code 200)
golang.org/x/net (download)
github.com/gobuffalo/plush (download)
github.com/gobuffalo/tags (download)
github.com/fatih/structs (download)
github.com/markbates/validate (download)
github.com/markbates/going (download)
github.com/satori/go.uuid (download)
github.com/serenize/snaker (download)
github.com/russross/blackfriday (download)
github.com/shurcooL/github_flavored_markdown (download)
github.com/microcosm-cc/bluemonday (download)
Fetching https://golang.org/x/net/html?go-get=1
Parsing meta tags from https://golang.org/x/net/html?go-get=1 (status code 200)
get "golang.org/x/net/html": found meta tag get.metaImport{Prefix:"golang.org/x/net", VCS:"git", RepoRoot:"https://go.googlesource.com/net"} at https://golang.org/x/net/html?go-get=1
get "golang.org/x/net/html": verifying non-authoritative meta tag
Fetching https://golang.org/x/net/html/atom?go-get=1
Parsing meta tags from https://golang.org/x/net/html/atom?go-get=1 (status code 200)
get "golang.org/x/net/html/atom": found meta tag get.metaImport{Prefix:"golang.org/x/net", VCS:"git", RepoRoot:"https://go.googlesource.com/net"} at https://golang.org/x/net/html/atom?go-get=1
get "golang.org/x/net/html/atom": verifying non-authoritative meta tag
github.com/shurcooL/highlight_diff (download)
github.com/sergi/go-diff (download)
github.com/sourcegraph/annotate (download)
github.com/sourcegraph/syntaxhighlight (download)
github.com/shurcooL/highlight_go (download)
github.com/shurcooL/octiconssvg (download)
github.com/shurcooL/sanitized_anchor_name (download)
github.com/spf13/cobra (download)
github.com/spf13/pflag (download)
github.com/markbates/pop (download)
github.com/cockroachdb/cockroach-go (download)
github.com/lib/pq (download)
github.com/go-sql-driver/mysql (download)
github.com/jmoiron/sqlx (download)
github.com/mattn/anko (download)
github.com/daviddengcn/go-colortext (download)
github.com/mattn/go-sqlite3 (download)
Fetching https://gopkg.in/yaml.v2?go-get=1
Parsing meta tags from https://gopkg.in/yaml.v2?go-get=1 (status code 200)
get "gopkg.in/yaml.v2": found meta tag get.metaImport{Prefix:"gopkg.in/yaml.v2", VCS:"git", RepoRoot:"https://gopkg.in/yaml.v2"} at https://gopkg.in/yaml.v2?go-get=1
gopkg.in/yaml.v2 (download)
github.com/markbates/deplist (download)
github.com/markbates/grift (download)
github.com/markbates/refresh (download)
github.com/fsnotify/fsnotify (download)
github.com/markbates/sigtx (download)
github.com/fatih/color/vendor/github.com/mattn/go-isatty
github.com/fatih/color/vendor/github.com/mattn/go-colorable
github.com/joho/godotenv
github.com/mitchellh/go-homedir
github.com/gobuffalo/buffalo/generators/assets
github.com/markbates/inflect
github.com/pkg/errors
golang.org/x/sys/unix
golang.org/x/net/context
github.com/gobuffalo/plush/token
golang.org/x/sync/errgroup
github.com/fatih/color
github.com/fatih/structs
github.com/markbates/going/wait
github.com/gobuffalo/plush/ast
github.com/gobuffalo/envy
github.com/gobuffalo/plush/lexer
github.com/gobuffalo/packr/builder
github.com/markbates/validate
github.com/gobuffalo/tags
github.com/gobuffalo/buffalo/generators
github.com/gobuffalo/packr
github.com/gobuffalo/plush/parser
github.com/satori/go.uuid
github.com/markbates/going/defaults
github.com/serenize/snaker
github.com/russross/blackfriday
golang.org/x/net/html/atom
github.com/sergi/go-diff/diffmatchpatch
golang.org/x/net/html
github.com/markbates/validate/validators
github.com/sourcegraph/annotate
github.com/shurcooL/sanitized_anchor_name
github.com/sourcegraph/syntaxhighlight
github.com/spf13/pflag
github.com/lib/pq/oid
github.com/lib/pq
github.com/shurcooL/highlight_go
golang.org/x/crypto/ssh/terminal
github.com/gobuffalo/makr
github.com/gobuffalo/buffalo/meta
github.com/gobuffalo/tags/form
github.com/shurcooL/highlight_diff
github.com/gobuffalo/buffalo/generators/assets/standard
github.com/gobuffalo/buffalo/generators/docker
github.com/sirupsen/logrus
github.com/gobuffalo/buffalo/generators/grift
github.com/gobuffalo/buffalo/generators/mail
github.com/gobuffalo/buffalo/generators/resource
github.com/gobuffalo/tags/form/bootstrap
github.com/gobuffalo/buffalo/generators/refresh
github.com/microcosm-cc/bluemonday
github.com/shurcooL/octiconssvg
github.com/go-sql-driver/mysql
github.com/jmoiron/sqlx/reflectx
github.com/markbates/going/randx
github.com/jmoiron/sqlx
github.com/spf13/cobra
github.com/markbates/pop/columns
github.com/gobuffalo/buffalo/generators/assets/webpack
github.com/gobuffalo/buffalo/generators/action
github.com/mattn/anko/ast
github.com/daviddengcn/go-colortext
github.com/mattn/go-sqlite3
gopkg.in/yaml.v2
github.com/mattn/anko/parser
github.com/cockroachdb/cockroach-go/crdb
github.com/gobuffalo/buffalo/buffalo/cmd/destroy
github.com/gobuffalo/buffalo/buffalo/cmd/generate
github.com/gobuffalo/buffalo/plugins
github.com/markbates/deplist
github.com/markbates/grift/grift
github.com/fsnotify/fsnotify
github.com/markbates/sigtx
github.com/shurcooL/github_flavored_markdown
github.com/markbates/grift/cmd
github.com/gobuffalo/plush
github.com/mattn/anko/vm
github.com/gobuffalo/buffalo/buffalo/cmd/build
github.com/markbates/refresh/refresh
github.com/mattn/anko/builtins/errors
github.com/mattn/anko/builtins/encoding/json
github.com/mattn/anko/builtins/flag
github.com/mattn/anko/builtins/io/ioutil
github.com/mattn/anko/builtins/github.com/daviddengcn/go-colortext
github.com/mattn/anko/builtins/fmt
github.com/mattn/anko/builtins/io
github.com/mattn/anko/builtins/math
github.com/mattn/anko/builtins/math/big
github.com/mattn/anko/builtins/math/rand
github.com/mattn/anko/builtins/net
github.com/mattn/anko/builtins/net/http
github.com/mattn/anko/builtins/net/url
github.com/mattn/anko/builtins/os
github.com/mattn/anko/builtins/os/exec
github.com/mattn/anko/builtins/os/signal
github.com/mattn/anko/builtins/path
github.com/mattn/anko/builtins/path/filepath
github.com/mattn/anko/builtins/regexp
github.com/mattn/anko/builtins/runtime
github.com/mattn/anko/builtins/sort
github.com/mattn/anko/builtins/strconv
github.com/mattn/anko/builtins/strings
github.com/mattn/anko/builtins/time
github.com/mattn/anko/builtins
github.com/markbates/pop/fizz
github.com/markbates/pop/fizz/translators
github.com/markbates/pop
github.com/markbates/pop/soda/cmd/generate
github.com/markbates/pop/soda/cmd/schema
github.com/gobuffalo/buffalo/generators/soda
github.com/markbates/pop/soda/cmd
github.com/gobuffalo/buffalo/generators/newapp
github.com/gobuffalo/buffalo/buffalo/cmd
github.com/gobuffalo/buffalo/buffalo
````