package melon

import (
	"github.com/goburrow/melon/core"
)

// Application is the default application which does nothing.
type Application struct {
	AppName        string
	InitializeFunc func(*core.Bootstrap)
	RunFunc        func(interface{}, *core.Environment) error
}

// Application implements core.Application interface.
var _ core.Application = (*Application)(nil)

func (app *Application) Name() string {
	if app.AppName == "" {
		return "melon-app"
	}
	return app.AppName
}

// Initializes the application bootstrap.
func (app *Application) Initialize(b *core.Bootstrap) {
	if app.InitializeFunc != nil {
		app.InitializeFunc(b)
	}
}

// When the application runs, this is called after the Bundles are run.
// Override it to add handlers, tasks, etc. for your application.
func (app *Application) Run(config interface{}, env *core.Environment) error {
	if app.RunFunc != nil {
		return app.RunFunc(config, env)
	}
	return nil
}
