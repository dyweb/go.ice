# OpenTracing

The doc on opentracing-go seems to be outdated ... use the one in jaeger as example

- https://github.com/yurishkuro/opentracing-tutorial/tree/master/go there is one example though ....
- https://github.com/jaegertracing/jaeger/blob/master/examples/hotrod/pkg/tracing/init.go#L31 `Init` creates a new instance of Jaeger tracer
- `SamplerConfig`, source code has detail about sampler types `const`, `probabilistic`, `rateLimiting`, `remote`
  - [ ] `SamplingServerURL` ... does this need to be set? seems only needed when using `remote`?

````go
// Init creates a new instance of Jaeger tracer.
func Init(serviceName string, metricsFactory metrics.Factory, logger log.Factory, hostPort string) opentracing.Tracer {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort: hostPort,
		},
	}
	tracer, _, err := cfg.New(
		serviceName,
		config.Logger(jaegerLoggerAdapter{logger.Bg()}),
		config.Observer(rpcmetrics.NewObserver(metricsFactory, rpcmetrics.DefaultNameNormalizer)),
	)
	if err != nil {
		logger.Bg().Fatal("cannot initialize Jaeger Tracer", zap.Error(err))
	}
	return tracer
}
````