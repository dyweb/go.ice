package config

// based on jaeger-client-go/config, some comment are copy and pasted here

type TracingConfig struct {
	Adapter  string                 `yaml:"adapter"`
	Sampler  TracingSamplerConfig   `yaml:"sampler"`
	Reporter TracingReporterConfig  `yaml:"reporter"`
	XXX      map[string]interface{} `yaml:",inline"`
}

type TracingSamplerConfig struct {
	// - for "const" sampler, 0 or 1 for always false/true respectively
	// - for "probabilistic" sampler, a probability between 0 and 1
	// - for "rateLimiting" sampler, the number of spans per second
	// - for "remote" sampler, param is the same as for "probabilistic"
	Type  string  `yaml:"type"`
	Param float64 `yaml:"param"`
	// TODO: we ignore sampling server, which can control sampling rate
	XXX map[string]interface{} `yaml:",inline"`
}

type TracingReporterConfig struct {
	LogSpans           bool                   `yaml:"logSpans"`
	LocalAgentHostPort string                 `yaml:"localAgentHostPort"`
	XXX                map[string]interface{} `yaml:",inline"`
}
