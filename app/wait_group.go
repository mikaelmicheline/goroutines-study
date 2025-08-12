package app

import (
	"fmt"
	"sync"
	"time"

	"github.com/urfave/cli"
)

func SetupWaitGroupExampleCommand(app *cli.App) {
	app.Commands = append(app.Commands, cli.Command{
		Name:  "wait-group",
		Usage: "Uses a WaitGroup to wait for Goroutines to finish",
		Action: func(c *cli.Context) {
			fmt.Println("WaitGroup example!")
			var waitGroup sync.WaitGroup

			countdown := func(label string, value int, sleepTime time.Duration) {
				for i := value; i > 0; i-- {
					fmt.Println(label, i)
					time.Sleep(sleepTime)
				}
				waitGroup.Done()
			}

			waitGroup.Add(2)
			go countdown("Countdown A:", 10, time.Millisecond*400)
			go countdown("Countdown B:", 10, time.Millisecond*700)
			waitGroup.Wait()
		},
	},
	)
}
