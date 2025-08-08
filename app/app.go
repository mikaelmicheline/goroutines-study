package app

import "github.com/urfave/cli"

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Study of Goroutines"

	SetupWaitGroupExampleCommand(app)

	return app
}
