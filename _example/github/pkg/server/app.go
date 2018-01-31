package server

import (
	dlog "github.com/dyweb/gommon/log"

	"github.com/at15/go.ice/_example/github/pkg/common"
)

// TODO: might move this to ice package ... or make a interface etc. ....
type App struct {
	name         string
	config       Config
	configFile   string
	configLoaded bool
	verbose      bool
	log          *dlog.Logger
}

func Name(name string) func(*App) error {
	return func(app *App) error {
		app.name = name
		return nil
	}
}

func NewApp() *App {
	app := &App{}
	app.log = dlog.NewStructLogger(log, app)
	return app
}

func (app *App) Config() Config {
	return app.config
}

func (app *App) Version() string {
	return common.Version()
}

// TODO: we can put logic of root command here ... actually .. it's just config ...
