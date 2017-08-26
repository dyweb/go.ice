# Design

The overall goal for go.ice is to be an API gateway for control plane like application, 
so performance, database (ORM) and template etc. are currently not its main concern.

- [Service](#service)
- [Plugin](#plugin)
- [Transport](#transport)
- [Frontend](#frontend)
- [Test](#test)
- [Deploy](#deploy)

## Service

Service is wrapper for remote services, it has different providers, some of them may be local (i.e. cookie based session service provider)

### Required

Must have (`?` means provider in low priority) in order to migrate from existing applications

- identity (authentication)
  - username + password
  - ? oauth
  - ? ldap
- authorization (access control)
  - role based access control (the semantic of role, permission is totally up to the user, we just store it)
    - file
    - ldap
    - cache
- session
  - cookie
  - ? cache
- cache
  - redis
  - ? solr
- full text search
  - solr
  - ? elastic search
  
### Desired  

Good to have (never gonna implement /w\ )

- tracing
  - opentracing
- monitoring
  - prometheus
- pubsub

### Required feature

- service registry (just a locally object, not to be confused with service registry in microservice)

### Desired features

- circuit breaker
- monitor (tracing) service upstream

## Plugin

Plugin can make use of all the services and mount themselves to API endpoints

- you write plugin in golang, and import it your `main.go`, i.e. 

<!-- TODO: a real runnable example -->

````go
package main

import  (
    _ "github.com/gaocegege/badge" // import the badge plugin
    "github.com/at15/go.ice"
)

func main() {
    ice.Run() // visit localhost:8080/api/plugins/badge
}
````

### Built in plugins

Good to have (never gonna implement /w\ )

- user avatar (based on username)
- file upload
  - access control by upload user
  
## Transport

### HTTP

- http2 is support automatically as long as you provide private key (just generate one locally and tell your browser to go ahead)

### gRPC
  
## Frontend

We expect you to use a SPA (single page application) and interact with go.ice in REST API only, there is no template support, 
but you can use builtin [html/template](https://golang.org/pkg/html/template/) if you really need it.

### Required

- serving static content

### Desired

- embedded static content into binary

## Test

- Makefile (template) for build & test using docker with different go version (1.8 & 1.9)
- Docker compose for service upstream, i.e. Redis, Solr etc.

## Deploy

- Builder image 
- Runner image