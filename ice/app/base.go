package app

var _ App = (*BaseApp)(nil)

type BaseApp struct {
	name         string
	description  string
	version      string
	configFile   string
	configLoaded bool
	verbose      bool
}

// use functional options https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

type BaseAppOptions func(a *BaseApp)

func Name(name string) func(app *BaseApp) {
	return func(app *BaseApp) {
		app.name = name
	}
}

func Description(desc string) func(app *BaseApp) {
	return func(app *BaseApp) {
		app.description = desc
	}
}

func Version(ver string) func(app *BaseApp) {
	return func(app *BaseApp) {
		app.version = ver
	}
}

func (b *BaseApp) Name() string {
	return b.name
}

func (b *BaseApp) Description() string {
	return b.description
}

func (b *BaseApp) Version() string {
	return b.version
}
