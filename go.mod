module github.com/dyweb/go.ice

go 1.13

require (
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dyweb/gommon v0.0.13
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/spf13/cobra v0.0.5
	google.golang.org/grpc v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.2.7
	gotest.tools v2.2.0+incompatible // indirect
)

replace github.com/docker/docker v1.13.1 => github.com/docker/engine v0.0.0-20190717161051-705d9623b7c1
