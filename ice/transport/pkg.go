package transport

type Client interface {
}

// TODO: tracing?
type Server interface {
	// ListenAddr return host:port so we can log it before actually start service
	ListenAddr() string
	// Port is to used avoid different transport using same port
	Port() int
	// Run starts the server and blocks current goroutine
	Run() error
}
