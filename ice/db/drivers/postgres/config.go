package postgres

var defaults = &Default{}

type Default struct {
}

func (c *Default) DockerImage() string {
	return "postgres"
}

func (c *Default) Port() int {
	return 5432
}

func (c *Default) NativeShell() string {
	return "psql"
}
