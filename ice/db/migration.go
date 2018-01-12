package db

type MigrationTask interface {
	Id() string
	Name() string
	Description() string
	Up() error
	Down() error
}
