package postgres

var defaultConfig = &Config{}

type Config struct {
}

func (c *Config) DockerImage() string {
	return "postgres"
}

func (c *Config) Port() int {
	return 5432
}

func (c *Config) NativeShell() string {
	return "psql"
}
