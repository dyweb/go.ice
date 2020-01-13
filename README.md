# go.ice

<h1 align="center">
        <br>
        <img width="100%" src="doc/media/go.ice.png" alt="go.ice">
        <br>
        <br>
        <br>
</h1>

[![GoDoc](https://godoc.org/github.com/dyweb/go.ice?status.svg)](https://godoc.org/github.com/dyweb/go.ice)
[![Build Status](https://travis-ci.org/dyweb/go.ice.svg?branch=master)](https://travis-ci.org/dyweb/go.ice)
[![codecov](https://codecov.io/gh/dyweb/go.ice/branch/master/graph/badge.svg)](https://codecov.io/gh/dyweb/go.ice)
[![loc](https://tokei.rs/b1/github/dyweb/go.ice)](https://github.com/dyweb/go.ice)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fat15%2Fgo.ice.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fat15%2Fgo.ice?ref=badge_shield)

go.ice is a server application toolkit with profiling in mind. It is still under early development.
See [roadmap](ROADMAP.md) for current progress.

Goals

- Write server implementation and client library at same time.
- Integrate with tracing and monitoring systems
- Support multiple transports, HTTP, gRPC
- A UI toolkit for building dashboard using [Vue](https://vuejs.org/)

Non Goals

- a Rails like web framework, [buffalo](https://github.com/gobuffalo/buffalo) is a good choice if you want one

## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fat15%2Fgo.ice.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fat15%2Fgo.ice?ref=badge_large)

## About

The name `go.ice` comes from [@arrowrowe](https://github.com/arrowrowe). The banner is drawn by [@at15][at15].

The project was started by [@at15][at15] as a server framework for developing tsdb benchmark tools & tsdb.
([Xephonhq](https://github.com/xephonhq) and [BenchHub](https://github.com/benchhub) and later transferred to dyweb.

[at15]: https://github.com/at15