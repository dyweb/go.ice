# 2018-12-15 Components

This doc is just some scratch I wrote thinking about how to develop go.ice and application using go.ice while
waiting my roommate to finish his operation work (on Friday night..., they are brave)

## New example application

A new application that contains client for docker, elasticsearch, k8s and unified dashboard.
Proposed names:

- iceka, icecream (ice in a container) by @arrowrowe
- unidashboard by @gaocegege
- iceuboard by @at15 (not very good ....)

Main purpose is to have a unified http client (high level wrapper on net/http) and a plugable project layout.

## Components

- cli, the standard layout for using cobra
  - a context implementation, application can create their own and embed this one
  - use context to set/read flag to abstract the underlying cli framework (mostly flag parsing)
  - also make actual cli logic testable
- web
  - show error in rest API
    - ajax (json)
    - html (user, developer, the error message should contains useful debug information)
- sql
  - use sqlx for now
  - migrate the manager, migration etc.
- http server util
  - router that show routes (originally simple router)
  - json API wrapper to avoid json encode decode
  - validation on request (need to have proper validation error, better check existing validation libraries)
- http client
  - add context to save body, log etc.
  - need to consider those tracing libraries
- ui
  - common templates (stuff to copy and paste, table, ajax etc.)
  - modular
- clients
  - docker client
  - k8s client
  - mesos client (only use official one's protobuf)