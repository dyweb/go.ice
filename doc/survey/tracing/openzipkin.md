# OpenZipkin

https://zipkin.io/ (so ... zipkin is renamed to openzipkin?...)
https://github.com/openzipkin

TODO:

- [ ] database schema, MySQL, Cassandra, Els etc.
- [ ] its UI 
  - https://github.com/openzipkin/zipkin/issues/1577
  - archived angular2 https://github.com/openzipkin-attic/zipkin-ui
- [ ] a spark job is needed for aggregates spans for use in the UI
   - https://github.com/openzipkin/zipkin-dependencies

Integrations:

- Go 13 stars https://github.com/openzipkin/zipkin-go
  - 289 stars https://github.com/openzipkin/zipkin-go-opentracing
- JDBC https://github.com/openzipkin/brave/tree/master/archive/brave-p6spy

## Quick start

````bash
docker run -d -p 9411:9411 openzipkin/zipkin
````

then visit `localhost:9411`

## Architecture

https://zipkin.io/pages/architecture.html

- only propagate IDs in-band
- completed spans are reported to Zipkin out-of-band
  - reporter send trace data via transport to Zipkin collectors, which persist trace data to storage
- transports
  - http
  - Kafka
  - [ ] Scribe ??
  
### Components

- collector, validate, store and index
  - [ ] what's the index?
- storage
  - Cassandra
  - ElasticSearch
  - MySQL
- search
- web ui

## Data model

- https://zipkin.io/pages/data_model.html
- https://zipkin.io/zipkin-api/#/default/post_spans

````json
[
  {
    "traceId": "string",
    "name": "string",
    "parentId": "string",
    "id": "string",
    "kind": "CLIENT",
    "timestamp": 0,
    "duration": 0,
    "debug": true,
    "shared": true,
    "localEndpoint": {
      "serviceName": "string",
      "ipv4": "string",
      "ipv6": "string",
      "port": 0
    },
    "remoteEndpoint": {
      "serviceName": "string",
      "ipv4": "string",
      "ipv6": "string",
      "port": 0
    },
    "annotations": [
      {
        "timestamp": 0,
        "value": "string"
      }
    ],
    "tags": {
      "additionalProp1": "string",
      "additionalProp2": "string",
      "additionalProp3": "string"
    }
  }
]
````

## Libraries

https://zipkin.io/pages/existing_instrumentations.html

- java https://github.com/openzipkin/brave
- go https://github.com/openzipkin/zipkin-go-opentracing
- js https://github.com/openzipkin/zipkin-js

