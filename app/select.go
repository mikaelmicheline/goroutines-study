package app

import (
	"fmt"
	"time"

	"github.com/urfave/cli"
)

func SetupSelectExampleCommand(app *cli.App) {
	app.Commands = append(app.Commands, cli.Command{
		Name:  "select",
		Usage: "Uses a Select command to receive from the first channel to send a value",
		Action: func(c *cli.Context) {
			fmt.Println("Select example!")
			channel1, channel2 := make(chan string), make(chan string)

			countdown := func(label string, value int, sleepTime time.Duration, channel chan string) {
				for i := value; i > 0; i-- {
					time.Sleep(sleepTime)
					channel <- fmt.Sprintf("%s%d", label, i)
				}
				close(channel)
			}

			go countdown("Countdown A:", 10, time.Millisecond*400, channel1)
			go countdown("Countdown B:", 10, time.Millisecond*700, channel2)

		ForLoop:
			for {
				select {
				case value, open := <-channel1:
					if open {
						fmt.Println(value)
						continue
					}
					channel1 = nil
					if channel2 == nil {
						break ForLoop
					}
				case value, open := <-channel2:
					if open {
						fmt.Println(value)
						continue
					}
					channel2 = nil
					if channel1 == nil {
						break ForLoop
					}
				}
			}
		},
	},
	)
}
