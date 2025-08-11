package app

import "github.com/urfave/cli"

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Study of Goroutines"
	app.Usage = "A command line Go application created to study Goroutines"

	SetupWaitGroupExampleCommand(app)
	SetupChannelExampleCommand(app)
	SetupChannelWithBufferExampleCommand(app)

	return app
}
