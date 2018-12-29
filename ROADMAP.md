# Roadmap

## 0.1.x

Use 0.1.x for v2 features, the v2 is actually v0.2.x since there weren't a usable v2

- [#34](https://github.com/dyweb/go.ice/pull/34) initial design

Implementation orders

- [ ] [#37](https://github.com/dyweb/go.ice/issues/37) httpclient
  - [ ] docker client
  - [ ] elastic search client
- [ ] cli wrapper
  - [ ] use ctx for table writer etc.
  - [ ] depends on gommon/termutil for asking password etc.
- [ ] UI
  - [ ] use element UI, port things from goyourcassandra
- [ ] generate http client and server stub based on spec
  - [ ] still need our own definition (may need more design on that)
- [ ] tracing
  - [ ] replace open tracing w/ open census
- [ ] database interface and migration etc.

Dependencies

- gommon, of course

GitHub

- pkg
  - api
  - cli
  - db
  - httpclient
  - ui
- lib
  - dockerclient
  - esclient
  - githubclient
- example
  - udash
- depend-on
  - gommon
- depend-by
  - ayi
- type
  - bug
  - backlog
  - code-style
  - refactor
  - design
  - new-feature
  - new-package
  - new-version
 
Projects

- api
- cli
- db
- httpclient
- ui