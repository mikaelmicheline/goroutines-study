package app

import (
	"fmt"
	"time"

	"github.com/urfave/cli"
)

func SetupChannelWithBufferExampleCommand(app *cli.App) {
	app.Commands = append(app.Commands, cli.Command{
		Name:  "channel-with-buffer",
		Usage: "Uses a Channel with Buffer to pause the sender when the Buffer's capacity is reached and to release the sender when the receiver frees space",
		Action: func(c *cli.Context) {
			fmt.Println("Channel with Buffer example!")
			channel := make(chan string, 4)

			go func() {
				for i := 10; i > 0; i-- {
					channel <- fmt.Sprintf("Countdown: %d", i)
					fmt.Println("Filling the Buffer")
				}
				close(channel)
			}()

			time.Sleep(time.Second * 2)
			for value := range channel {
				time.Sleep(time.Millisecond * 400)
				fmt.Println(value)
			}
		},
	},
	)
}
