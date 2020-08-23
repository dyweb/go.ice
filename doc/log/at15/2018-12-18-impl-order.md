# 2018-12-18 Implementation order

This doc list implementation order based on current project status

- cli (1-2d)
  - (at15) init context
- httpserver (1 week)
  - (gaocegege) survey on generating client server stub for http server (REST API) [#35](https://github.com/dyweb/go.ice/issues/35)
  - (at15) some trial implementation in playground
  - we can use raw `http.Handler` in the middle
- httpclient (1d copy and paste)
  - (at15) base client wrapper around net/http, drain body, retry etc. needed regardless of code generation
- database (2d)
  - (at15) port existing go.ice adapters, use sqlx to replace unfinished query builder logic
- ui (1d copy and paste)
  - (at15) init ui components based on udash, it might be easier to start udash first and extract reusable part out
- udash (1d copy and paste)
  - (at15) query result and show them in table, like [goyourcassandra](https://github.com/at15/goyourcassandra)
- grpc
  - on hold
- tracing
  - (at15) do the survey on opencensus
  
Put it in a graph

````text
gaocegege   ----------- client & server gen --------------------------------------

at15  ---cli---
             ----database---
             ----  ui   ----
                            --- udash ---  
                                      --- httpclient ----
                                            --- udash with client & server gen ---
````
