package migration

type Task interface {
	Id() string
	Name() string
	Description() string
	Up() error
	Down() error
}
