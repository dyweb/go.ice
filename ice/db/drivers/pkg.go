package drivers

// TODO: this might be called ... default properties?
type DefaultConfig interface {
	DockerImage() string
	Port() int
	NativeShell() string
}

type Driver interface {
	DefaultConfig() DefaultConfig
}
