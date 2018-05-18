# go.ice

[![Build Status](https://travis-ci.org/dyweb/go.ice.svg?branch=master)](https://travis-ci.org/dyweb/go.ice)
[![codecov](https://codecov.io/gh/dyweb/go.ice/branch/master/graph/badge.svg)](https://codecov.io/gh/dyweb/go.ice)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fat15%2Fgo.ice.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fat15%2Fgo.ice?ref=badge_shield)

A server application toolkit with profiling in mind

Goals

- integration with tracing and monitoring systems
- profiling remote services as well as essential local operations
- multiple transport support, HTTP, gRPC
- write server implementation and client library at same time

Non Goals

- a Rails like web framework

## Roadmap

cli

- [ ] wrapper for cobra
- [ ] interactive, confirm, progress bar etc.
- [ ] repl like, i.e. `ayi -i` this may requires significant change in how flags, config etc. are handled

RPC

- [ ] generated server and client based on DSL like [reproto](https://github.com/reproto/reproto)
  - [ ] can also generate UI like swagger (btw: I used to like swagger)
  - might just write AST in Go, there are go frameworks doing similar i.e. [goa](https://github.com/goadesign/goa)
    - it generates js code as well
  - need to generate typescript code to work with Angular
- [ ] might support raw tcp server (will be used heavily in xephon)

Database

- [ ] adapter for different databases
  - [ ] a lot of database specific features can be found in mature web frameworks like Laravel
- [ ] migration
  - may not be used production, but handy for iteration on small projects
- [ ] cli for different database like [usql](https://github.com/xo/usql)
  - [ ] including Cassandra (mainly due to need in work)
- [ ] Web UI like [pgweb](https://github.com/sosedoff/pgweb), might just use [fe-template](https://github.com/at15/fe-template)
- [ ] SQL injection test

Tracing

- [ ] another interface, use [opentracing-go](https://github.com/opentracing/opentracing-go) and [opencensus-go](https://github.com/census-instrumentation/opencensus-go) as implementation, don't use them directly

Metrics

- [ ] integration with [libtsdb-go](https://github.com/libtsdb/libtsdb-go)

Profiling

- [ ] integration with pprof
- [ ] UI for existing tools like [360EntSecGroup-Skylar/goreporter](https://github.com/360EntSecGroup-Skylar/goreporter)

## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fat15%2Fgo.ice.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fat15%2Fgo.ice?ref=badge_large)