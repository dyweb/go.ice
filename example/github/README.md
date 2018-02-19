# GitHub example

- login using github oauth
- list the (signed in) user's repos
- receive webhook of user's push request

````bash
# run dep ensure in project root first!
dep ensure
# start jaeger
docker-compose up
# open another terminal
go run cmd/icehubd/main.go
# visit http://localhost:7080/github/login to login
````

![tracing](icehub-tracing.png)