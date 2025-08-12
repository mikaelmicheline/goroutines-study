package app

import (
	"fmt"
	"time"

	"github.com/urfave/cli"
)

func SetupChannelExampleCommand(app *cli.App) {
	app.Commands = append(app.Commands, cli.Command{
		Name:  "channel",
		Usage: "Uses a Channel to send messages from an additional Goroutine to the main Goroutine (main function)",
		Action: func(c *cli.Context) {
			fmt.Println("Channel example!")
			channel := make(chan string)

			go func() {
				for i := 10; i > 0; i-- {
					time.Sleep(time.Millisecond * 400)
					channel <- fmt.Sprintf("Countdown: %d", i)
				}
				close(channel)
			}()

			for value := range channel {
				fmt.Println(value)
			}
		},
	},
	)
}
