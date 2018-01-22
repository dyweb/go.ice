package mysql

var defaults = &Default{}

type Default struct {
}

func (c *Default) DockerImage() string {
	return "mysql"
}

func (c *Default) Port() int {
	return 3306
}

func (c *Default) NativeShell() string {
	return "mysql"
}
