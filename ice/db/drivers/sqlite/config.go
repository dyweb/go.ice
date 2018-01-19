package sqlite


var defaults = &Default{}

type Default struct {
}

// FIXME: there is not official docker image for sqlite, and we need to use docker for sqlite3 shell
func (c *Default) DockerImage() string {
	return ""
}

func (c *Default) Port() int {
	return 0
}

func (c *Default) NativeShell() string {
	return "sqlite3"
}
