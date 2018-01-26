# Survey

For survey on individual Go frameworks, see [survey-go](../survey-go)

TODO

- API documentation, swagger etc.
  - https://github.com/jaegertracing/jaeger-idl/blob/master/swagger/zipkin2-api.yaml
  - [ ] https://github.com/savaki/swag generate swagger doc from API definition
- membership etc.
  - https://github.com/hashicorp/memberlist used by https://github.com/hashicorp/serf
  - https://github.com/weaveworks/mesh I think is used by prometheus alertmanager
    - [Replace weaveworks/mesh](https://github.com/prometheus/alertmanager/issues/1200) 
  - https://github.com/hashicorp/go-discover use cloud service provider's API to discover nodes
- cmd
  - prometheus is using kingpin https://github.com/prometheus/prometheus/issues/2455 instead of cobra
    - I don't like viper because it is not type safe, cobra is better than https://github.com/urfave/cli, maybe try kingpin some time.