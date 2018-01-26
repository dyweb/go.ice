# OpenTracing

- https://github.com/opentracing
- http://opentracing.io/documentation/

Caveats

- https://github.com/opentracing/opentracing-go/blob/master/gocontext.go it's still using `x/net/context` instead of `context`
  - but it seems it work because the alias feature added in go1.9 https://golang.org/doc/go1.9#language

Integrations

- https://github.com/opentracing-contrib/java-jdbc

## Introduction

- no vendor lock-in
- a DAG
- it's even possible to model a child span with multiple parents (i.e. a buffer flush may descend from the multiple writes that filled said buffer?typo)

## Quick Start

- jaeger https://github.com/jaegertracing/jaeger

````bash
docker run -d -p 5775:5775/udp -p 16686:16686 jaegertracing/all-in-one:latest
````
## Concepts and Terminology

- https://github.com/opentracing/specification/blob/master/specification.md
  - https://github.com/opentracing/specification/blob/master/specification.md#set-a-baggage-item
  - in bound, passed along, overhead could be big 
- https://github.com/opentracing/specification/blob/master/semantic_conventions.md

## API

- cross process http://opentracing.io/documentation/pages/api/cross-process-tracing.html
  - inject, extract, and carriers
  - format: text map and binary
  
> What the OpenTracing implementations choose to store in these Carriers is not formally defined by the OpenTracing specification, 
but the presumption is that they find a way to encode "tracer state" about the propagated SpanContext 
(e.g., in Dapper this would include a trace_id, a span_id, and a bitmask representing the sampling status for the given trace) 
as well as any key:value Baggage items.

## Best Practices

### Common Use Cases

In-Process Request Context Propagation

- implicit propagation
- explicit propagation i.e. Go, no thread local

Logging events

> We have already used log in the client Span use case. Events can be logged without a payload, and not just where the Span is being created / finished. 
For example, the application may record a cache miss event in the middle of execution, 
as long as it can get access to the current Span from the request context

Setting Sampling Priority Before the Trace Starts

- `tags.SAMPLING_PRIORITY: 1`

### Instrumenting Large Systems

http://opentracing.io/documentation/pages/instrumentation/instrumenting-large-systems.html

- it provides a conceptual example

### Instrumenting Frameworks

http://opentracing.io/documentation/pages/instrumentation/instrumenting-frameworks.html

- https://github.com/opentracing-contrib/go-stdlib

## Supported Tracers

http://opentracing.io/documentation/pages/supported-tracers.html

- Zipkin
- Jaeger
- Appdash
- LightStep