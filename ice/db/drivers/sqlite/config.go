package sqlite


var defaultConfig = &Config{}

type Config struct {
}

// FIXME: there is not official docker image for sqlite, and we need to use docker for sqlite3 shell
func (c *Config) DockerImage() string {
	return ""
}

func (c *Config) Port() int {
	return 0
}

func (c *Config) NativeShell() string {
	return "sqlite3"
}
