# 2018-12-29 HTTP Client context

This doc describes what is needed for context and client

Both

- header
- retry
- encoding

Context only

- query parameter
- is a stream request

Client only

- bath path `/api/v1`

Basic features

- apply header
- apply query parameter
- encode request body as json if it is not `io.Reader`

Advanced features

- decode response body when 2xx
- decode response body when 4xx
  - this is why we always needed extra wrapper ...
- save response and request body