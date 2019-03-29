package koa

import "os"

var (
	app *Application
)

func App() *Application {
	if app == nil {
		configFile := ""

		for i, arg := range os.Args {
			if arg == "--config" {
				if len(arg) > i+1 {
					configFile = os.Args[i+1]
					break
				}
			}
		}

		app = &Application{}
		app.initConfig(configFile)
	}

	return app
}
