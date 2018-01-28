package server

import (
	"github.com/at15/go.ice/_example/github/pkg/common"
)

type App struct {
	config       Config
	configFile   string
	configLoaded bool
	verbose      bool
}

func (app *App) Config() Config {
	return app.config
}

func (app *App) Version() string {
	return common.Version()
}

// TODO: we can put logic of root command here ... actually .. it's just config ...
