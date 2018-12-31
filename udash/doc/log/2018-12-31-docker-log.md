# 2018-12-31 Docker log

This doc discuss how to implement a docker log UI in udash [#40](https://github.com/dyweb/go.ice/issues/40)

Backend

- [x] a container that generates log (could be udash itself ...)
  - `docker run --rm nginx -p 8080:80`
- [ ] allow stream log using docker client

UI

- [ ] init vue using yarn and vue cli
- [ ] how to 'tail' log in frontend

Dependencies

- [ ] cli on docker, might work on the cli package, or just raw cobra for now
- [ ] ui, not sure about how to reuse though ....