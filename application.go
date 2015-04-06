package melon

import (
	"github.com/goburrow/melon/core"
)

// Application is the default application which supports server command.
type Application struct {
	// Name of the application
	name          string
	configuration interface{}
}

// Application implements core.Application interface.
var _ core.Application = (*Application)(nil)

func (app *Application) Name() string {
	if app.name == "" {
		app.name = "melon-app"
	}
	return app.name
}

func (app *Application) SetName(name string) {
	app.name = name
}

// Initializes the application bootstrap.
func (app *Application) Initialize(bootstrap *core.Bootstrap) {
	bootstrap.AddCommand(&CheckCommand{})
	bootstrap.AddCommand(&ServerCommand{})
}

// When the application runs, this is called after the Bundles are run.
// Override it to add handlers, tasks, etc. for your application.
func (app *Application) Run(interface{}, *core.Environment) error {
	return nil
}
